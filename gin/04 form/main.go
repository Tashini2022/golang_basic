package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// get data thst submitted by form chart
func main() {
	r := gin.Default()
	// input template files
	r.LoadHTMLGlob("template/*")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		// username := c.PostForm("usrname")
		password := c.DefaultPostForm("pwd", "****")
		username, ok := c.GetPostForm("usrname")
		if !ok {
			username = "somebody"
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})
	r.Run(":9090")
}
