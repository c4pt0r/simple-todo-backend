package main

import (
	"github.com/gin-gonic/gin"
)

type TodoItem struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`

	UserID string
}

type APIResp struct {
	RetCode int         `json:"c"`
	Message string      `json:"m"`
	Data    interface{} `json:"d"`
}

var (
	mockTodoItems = []TodoItem{
		{
			ID:        "1",
			Title:     "hello 1",
			Completed: false,
		},
		{
			ID:        "2",
			Title:     "hello 2",
			Completed: false,
		},
		{
			ID:        "3",
			Title:     "hello 3",
			Completed: false,
		},
	}
)

func OKResp(data interface{}) APIResp {
	return APIResp{
		Data: data,
	}
}

func ErrResp(code int, err error) APIResp {
	return APIResp{
		RetCode: code,
		Message: err.Error(),
	}
}

func main() {
	initDB()
	router := gin.Default()
	router.Static("/static", "./assets/static")
	router.StaticFile("/", "./assets/index.html")
	// Listen and serve on 0.0.0.0:8080

	router.GET("/api/todos", func(c *gin.Context) {
		c.JSON(200, OKResp(mockTodoItems))
	})

	router.PATCH("/api/todo", func(c *gin.Context) {
		item, err := GetTodoItemFromGinContext(c)
		if err != nil {
			c.JSON(500, ErrResp(500, err))
			return
		}
		// if it's not a completed event, that should be update content
		if item.Completed == false {
			// update todo item
		}
		c.JSON(200, OKResp(item))
	})

	router.DELETE("/api/todo", func(c *gin.Context) {
		_, err := GetTodoItemFromGinContext(c)
		if err != nil {
			c.JSON(500, ErrResp(500, err))
			return
		}
		// delete item
	})

	router.POST("/api/todo", func(c *gin.Context) {
		item, err := GetTodoItemFromGinContext(c)
		if err != nil {
			c.JSON(500, ErrResp(500, err))
			return
		}
		err = CreateItemForUser(0, item)
		if err != nil {
			c.JSON(500, ErrResp(500, err))
			return
		}
		c.JSON(200, OKResp(item))
	})

	router.Run(":8080")
}
