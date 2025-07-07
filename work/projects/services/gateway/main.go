package main

import (
	"github.com/gin-gonic/gin"
	"projects/services/gateway/bootstrap"
)

// ab -n 2 -c 2 http://127.0.0.1:8081/v1/test
func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	bootstrap.InitRoute(router)
	router.Run(":8081") // 监听并在 0.0.0.0:8081 上启动服务
}
