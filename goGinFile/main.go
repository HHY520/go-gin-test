package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/uplood", func(c *gin.Context) {
		// 从请求中读取文件
		f, err := c.FormFile("f1")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": false,
			})
		} else {
			// 将读取到的文件保存在服务端本地
			dst := fmt.Sprintf("./%s", f.Filename)
			//dst := path.Join("./", f.Filename)
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": true,
			})
		}
	})
	r.Run(":9099")
}
