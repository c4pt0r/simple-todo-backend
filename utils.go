package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetTodoItemFromGinContext(c *gin.Context) (*TodoItem, error) {
	body, err := c.GetRawData()
	if err != nil {
		return nil, err
	}
	// get item content from body
	var item TodoItem
	json.Unmarshal(body, &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
