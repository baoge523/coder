package json_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 注意调用json.Marshal 或者UnMarshal时，自定义的结构体对象需要是public的访问权限，不然json无法访问

type User struct {
	Name string // public 访问
	age int // package 访问
}


func TestJsonAccess(t *testing.T) {

	user := &User{
		Name: "Andy",
		age: 18,
	}

	marshal, _ := json.Marshal(user)

	fmt.Println(string(marshal))

	// 从结果可以看出，age是没有被json序列化的，因为age只在当前的json_test可用
	assert.Equal(t, "{\"Name\":\"Andy\"}",string(marshal))
}
