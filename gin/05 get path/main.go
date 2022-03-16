package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// to get URL address path parameter
func main() {
	r := gin.Default()
	r.GET("/user/:name/:age", func(c *gin.Context) {
		// to get path parameter:Param
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"Name": name,
			"Age":  age,
		})
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		// to get path parameter:Param
		Year := c.Param("year")
		Month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"Year":  Year,
			"Month": Month,
		})
	})
	r.Run(":9090")
}
