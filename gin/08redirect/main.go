package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 重定向到必应
	r.GET("/exe", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://cn.bing.com")
	})
	// 重定向到/x
	r.GET("/index", func(c *gin.Context) {
		// 修改请求的url
		c.Request.URL.Path = "/x"
		// 继续处理
		r.HandleContext(c)
	})
	r.GET("/x", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "转移成功",
		})
	})

	//r.Any("/test", func(c *gin.Context) {.
	//	可同时包含get post delete 等方法.})
	// 用户访问未定义地址
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "请输入正确地址"})
	})
	r.Run(":9090")
}

// 路由组
/*userGroup := r.Group("/user")
{
	userGroup.GET("/index", func(c *gin.Context) {...})
	userGroup.GET("/login", func(c *gin.Context) {...})
	userGroup.POST("/login", func(c *gin.Context) {...})

}

//支持嵌套
shopGroup := r.Group("/shop")
{
	shopGroup.GET("/index", func(c *gin.Context) {...})
	shopGroup.GET("/cart", func(c *gin.Context) {...})
	shopGroup.POST("/checkout", func(c *gin.Context) {...})
	// 嵌套路由组
	xx := shopGroup.Group("xx")
	xx.GET("/oo", func(c *gin.Context) {...})
}
*/
