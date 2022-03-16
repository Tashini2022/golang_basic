package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Middleware

// 统计函数耗时
func statecost() gin.HandlerFunc {
	//方便加入其余功能模块
	return func(c *gin.Context) {
		state := time.Now()
		fmt.Println("func1 in!")
		//c.Set在请求上下文中设置值，使后续的处理函数能够取到该值
		c.Set("name", "它是你")
		//c.Next()调用该请求的其余处理函数
		c.Next()
		//c.Abort()//阻止调用剩余处理函数
		cost := time.Since(state)
		log.Println(cost)
		fmt.Println("func1 out!")
	}
}

func func2() gin.HandlerFunc {
	//方便加入其余功能模块
	return func(c *gin.Context) {
		fmt.Println("func2 in!")
		c.Next()
		//c.Abort()//阻止调用剩余处理函数
		fmt.Println("func2 out!")
	}
}

func respond(c *gin.Context) {
	fmt.Println("func3 now!")
	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
	})
	//c.Get在请求上下文中获取值，使后续的处理函数能够取到该值
	name := c.MustGet("name").(string)

	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}
func main() {
	r := gin.Default()
	// 为全局路由注册中间件
	r.Use(statecost(), func2())

	r.GET("/user", respond)
	r.Run()

}

/*为路由组注册中间件
//写法一
usergroup:=r.Group("/xxx")
usergroup.Use(middleware())
{
	usergroup.GET("/123",func(c *gin.Context){...})
	...
}
//写法二
usergroup:=r.Group("/xxx"，middleware())
{
	usergroup.GET("/123",func(c *gin.Context){...})
	...
}
*/
