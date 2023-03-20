package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由与路由组

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"method": "Post",
			})
		default:
			c.JSON(http.StatusOK, gin.H{
				"method": "NO know",
			})
		}
	})
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		})
		shopGroup.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		})
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"method": "404",
		})
	})
	r.Run(":9099")
}
