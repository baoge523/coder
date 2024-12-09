package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // 得到一个engine
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/createUser", CreateUser)

	r.Run("127.0.0.1:8080") // 监听并在 0.0.0.0:8080 上启动服务
}

// json:"GroupId" binding:"required"`
type User struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age"  binding:"required"`
}

func CreateUser(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("%s", err.Error()),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":  "create ok",
		"userInfo": fmt.Sprintf("name = %s,age = %d", u.Name, u.Age),
	})
}
