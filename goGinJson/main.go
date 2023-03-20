package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		// 方法一
		//data := map[string]interface{}{
		//	"name":    "小王子",
		//	"age":     18,
		//	"message": "hello",
		//}
		//data := gin.H{"name": "小王子", "age": 20, "message": "hello"}

		// 方法二
		var msg struct {
			Name    string `json:"name"`
			Age     int
			Message int
		}
		msg.Name = "哈哈"
		msg.Age = 18
		msg.Message = 200
		c.JSON(http.StatusOK, msg)
	})

	r.Run(":9099")

}
