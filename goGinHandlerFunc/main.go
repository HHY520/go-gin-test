package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Handlerfunc
func indexHandler(c *gin.Context) {
	fmt.Println("index in...")
	name, ok := c.Get("name") // 可以通过c.Get在请求上下文中获取值
	if !ok {
		name = "默认值"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 定义一个中间件
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	start := time.Now()
	// 计时
	// 调用该请求的剩余处理程序
	c.Next()
	// 不调用该请求的剩余处理程序
	// c.Abort()
	cost := time.Since(start)
	fmt.Println(cost)
	fmt.Println("m1 out...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	start := time.Now()
	//go func xx(c.Copy()) // 只能使用c的拷贝
	c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
	// 计时
	// 调用该请求的剩余处理程序
	c.Next()
	// 阻止用该请求的剩余处理程序
	// c.Abort()
	cost := time.Since(start)
	fmt.Println(cost)
	fmt.Println("m2 out...")
}

func main() {
	r := gin.Default() // 默认使用了Logger和Recovery中间件
	//r := gin.New() // 不注册任何中间件
	r.Use(m1, m2) // 全局注册中间件
	// GET is a shortcut for router.Handle("GET", path, handlers).
	r.GET("/index", indexHandler)

	r.Run(":9099")
}
