package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	f := gin.Default()

	f.GET("/web", func(c *gin.Context) {
		//name := c.Query("name") // 获取 ? 后面携带的参数
		//name := c.DefaultQuery("name", "空") // 取不到设置默认值
		name, ok := c.GetQuery("name")
		if !ok {
			name = "空"
		}
		c.JSON(http.StatusOK, gin.H{"name": name})
	})

	f.Run(":9099")
}
