package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Helllo golang!",
	})
}

func main() {
	r := gin.Default() //返回一个路由引擎

	//指定用户使用GET请求访问/hello时,执行sayHello函数
	// r.GET("/hello", sayHello)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})

	r.POST("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "DELETE",
		})
	})

	// 启动服务,如果不设置则默认在0.0.0.0:8080启动服务
	r.Run(":9090")
}
