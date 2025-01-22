package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go_web/gin/entity"
	"go_web/gin/usecase/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

func buildServer() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.BasicRequest())
	engine.GET("/ping", func(context *gin.Context) {
		context.String(200, "ping")
	})
	engine.POST("/createUser", CreateUser)

	group := engine.Group("/group1")
	{
		group.POST("/create", CreateUser2)
	}

	return engine
}

// addPost 添加接口
func addPost(e *gin.Engine) *gin.Engine {
	e.POST("/add/user", func(context *gin.Context) {
		var user entity.User
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

	user := &entity.User{
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
	assert.Equal(t, string(marshal), res.Body.String())
}

// test: if struct define such as: `json:"age"  binding:"required"`, when use gin,if age not exists,it must throw error
func TestCreateUser(t *testing.T) {
	server := buildServer()
	server = addPost(server)

	res := httptest.NewRecorder() // 获取一个httptest的responseRecorder:用于接收请求响应
	uStr := "{\"name\":\"Andy\"}"
	request, _ := http.NewRequest("POST", "/createUser", bytes.NewBuffer([]byte(uStr)))

	server.ServeHTTP(res, request)
	fmt.Println(res)
}

func TestJson(t *testing.T) {
	var u entity.User
	uStr := "{\"name\":\"Andy\"}"

	if err := json.Unmarshal([]byte(uStr), &u); err != nil {
		fmt.Println("error")
		return
	}

	fmt.Printf("%v", u)

}
