package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//r.LoadHTMLFiles("templates/posts/index.tmpl", "templates/users/index.tmpl") // 模版解析
	r.LoadHTMLGlob("templates/**/*")             // 模版解析
	r.GET("/posts/index", func(c *gin.Context) { // 模版渲染
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "posts",
		})
	})
	r.GET("users/index", func(c *gin.Context) { // 模版渲染
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "users",
		})
	})

	r.Run(":9099")
}
