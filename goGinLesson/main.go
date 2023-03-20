package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"pass"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//u := UserInfo{
		//	Username: username,
		//	Password: password,
		//}
		//fmt.Println(u)
		//c.JSON(http.StatusOK, gin.H{
		//	"massage": "ok",
		//})
		var u UserInfo
		err := c.ShouldBind(&u) // 需要打 tag 标签 "from"
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"username": u.Username,
				"password": u.Password,
			})
		}
	})

	r.POST("/form", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u) // 需要打 tag 标签 "from"
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"username": u.Username,
				"password": u.Password,
			})
		}
	})

	r.POST("/json", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u) // 需要打 tag 标签 "from"
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"username": u.Username,
				"password": u.Password,
			})
		}
	})

	r.Run(":9099")
}
