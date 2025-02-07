package middleware

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"time"
)

// 自定义的middleware
func BasicRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		reqJsonByte, err := ioutil.ReadAll(c.Request.Body) // 注意这里的body只允许读取一次，第二次读取时，会读取不到
		if err != nil {
			log.Printf("error %v \n", err)
			return
		}
		log.Printf("request boy %s", string(reqJsonByte))
		// Set example variable
		c.Set("name", "aaa") // 需要通过c.MustGet("name")的方式获取

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
