package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"user" json:"user"`
	Password string `form:"password" json:"password"`
}

// 参数绑定
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		username := c.Query("user")
		password := c.Query("password")
		u := UserInfo{
			Username: username,
			Password: password,
		}
		fmt.Printf("%#v\n", u)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "好久不见！",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}

	})
	r.POST("/user", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"Status": "ok",
			})
		}
	})
	r.Run(":9090")
}
