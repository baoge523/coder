package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/gin/entity"
)

// json:"GroupId" binding:"required"`

func CreateUser(c *gin.Context) {
	var u entity.User
	// 获取从 c.set中的数据
	name := c.MustGet("name").(string)
	u.Name = name
	c.JSON(200, gin.H{
		"message":  "create ok",
		"userInfo": fmt.Sprintf("name = %s,age = %d", u.Name, u.Age),
	})
}

func CreateUser2(c *gin.Context) {
	var u entity.User
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
