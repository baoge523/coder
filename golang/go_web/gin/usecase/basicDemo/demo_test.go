package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type User struct {
	Name string
}

func buildServer() *gin.Engine {
	engine := gin.Default()
	engine.GET("/ping", func(context *gin.Context) {
		context.String(200, "ping")

	})
	return engine
}

func addPost(e *gin.Engine) *gin.Engine {
	e.POST("/add/user", func(context *gin.Context) {
		var user User
		context.BindJSON(&user)
		context.JSON(200, user)
	})
	return e
}

// TestGetRequest 验证gin如何注册get请求，并请求得到对应的响应
func TestGetRequest(t *testing.T) {
	server := buildServer() // 构建gin-web

	w := httptest.NewRecorder() // 得到一个初始化的响应器 ResponseRecorder

	// 通过http.newRequest 构造请求
	req, _ := http.NewRequest("GET", "/ping", nil)

	// gin的engine执行web请求，w用于处理接收的响应、req表示请求
	server.ServeHTTP(w, req)

	// "github.com/stretchr/testify/assert"
	// 断言操作
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "ping", w.Body.String())
}

func TestPostRequest(t *testing.T) {
	server := buildServer()
	server = addPost(server)

	res := httptest.NewRecorder() // 获取一个httptest的responseRecorder:用于接收请求响应

	user := &User{
		Name: "andy",
	}
	marshal, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(marshal))
	request, _ := http.NewRequest("POST", "/add/user", bytes.NewBuffer(marshal))

	server.ServeHTTP(res, request)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, string(marshal),res.Body.String())
}
