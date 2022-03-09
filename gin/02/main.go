package main

import (
	"net/http"

	"gthub.com/gin-gonic/gin"
)

func Json(c *gin.Context) {
	// 方法一：使用map
	// data := map[string]interface{}{
	// 	"name":    "它是你",
	// 	"message": "好久不见！",
	// 	"age":     18,
	// }
	// 使用gin.H
	data := gin.H{
		"name":    "它是你",
		"message": "好久不见！",
		"age":     18,
	}

	// 方法二：使用struct
	var msg struct {
		Name    string `json:"user"`
		Message string
		Age     int
	}
	msg.Name = "小王子"
	msg.Message = "Hello world!"
	msg.Age = 18
	c.JSON(http.StatusOK, data)
	c.JSON(http.StatusOK, msg)
}
func main() {
	r := gin.Default()

	r.GET("/json", Json)
	r.Run(":9000")
}
