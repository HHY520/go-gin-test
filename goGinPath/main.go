package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/:username/:age", func(c *gin.Context) {
		username := c.Param("username")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"usaename": username,
			"age":      age,
		})
	})
	r.Run(":9099")
}
