package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func msg(c *gin.Context) {
	// // three methods:1. get query string parameter"key=value"
	// password := c.Query("user password")
	// // three methods:2.
	// hi := c.DefaultQuery("query", "Please input right message!")
	// // three methods:3.
	name, ok := c.GetQuery("name")
	if !ok {
		name = "can not get name!"
	}

	c.JSON(http.StatusOK, gin.H{
		"user": name,
		// "message":  hi,
		// "password": password,
	})
}

// Query string
func main() {
	r := gin.Default()
	r.GET("/web", msg)
	r.Run(":9000")
}
