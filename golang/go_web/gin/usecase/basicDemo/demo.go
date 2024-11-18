package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()  // 得到一个engine
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("127.0.0.1:8080") // 监听并在 0.0.0.0:8080 上启动服务
}
