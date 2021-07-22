package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func initDB() {
	var err error
	dsn := "root@tcp(127.0.0.1:4000)/simple_todo"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

var (
	SOURCE_GITHUB string = "github"
)

type TodoItemModel struct {
	ID        int64  `gorm:"column:id"`
	Title     string `gorm:"column:title"`
	Completed bool   `gorm:"column:completed"`
	UserID    int64  `gorm:"column:user_id"`
}

func (TodoItemModel) TableName() string {
	return "todo_items"
}

type UserModel struct {
	ID       int64  `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Source   string `gorm:"column:source"`
	Token    string `gorm:"column:token"`
}

func (UserModel) TableName() string {
	return "users"
}

func ItemToItemModel(userID int64, item *TodoItem) (*TodoItemModel, error) {
	ret := &TodoItemModel{
		Title:     item.Title,
		Completed: item.Completed,
		UserID:    userID,
	}
	if item.ID != "" {
		id, err := strconv.ParseInt(item.ID, 10, 64)
		if err != nil {
			return nil, errors.New("invalid todo item")
		}
		ret.ID = id
	}
	return ret, nil
}

func CreateItemForUser(userID int64, item *TodoItem) error {
	m, err := ItemToItemModel(userID, item)
	if err != nil {
		return err
	}
	db.Create(m)
	if db.Error != nil {
		return db.Error
	}
	item.ID = fmt.Sprintf("%d", m.ID)
	return nil
}
