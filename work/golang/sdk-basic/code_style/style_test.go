package code_style

import (
	"encoding/json"
	"fmt"
	"testing"
)

// json 不会操作private权限的属性
func TestSecurity0(t *testing.T) {
	confStr := `{"username":"zhangsan","password":"aaaaa","supervisor_agent":"127.0.0.1","port":8088,"ttl":5000,"min_conn":5,"max_conn":100}`
	conf := &DBConfig0{}
	// 原因是 username 和 password 的访问是私有的，json操作不同包的struct对象时，需要public的访问权限
	err := json.Unmarshal([]byte(confStr), conf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conf.getPassword())
}

// 能控制print的打印，但是不能控制json的打印
func TestSecurity1(t *testing.T) {
	config := &DBConfig1{
		Username: "zhangsan",
		Password: "aaaaa",
		Host:     "127.0.0.1",
		Port:     8088,
		Ttl:      5000,
		MinConn:  5,
		MaxConn:  100,
	}
	fmt.Println(config)

	fmt.Print("aaa")

	// json.Marshal需要public访问权限
	marshal, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
	}
	// {"username":"zhangsan","password":"aaaaa","supervisor_agent":"127.0.0.1","port":8088,"ttl":5000,"min_conn":5,"max_conn":100}
	fmt.Println(string(marshal))
}

func TestSecurity(t *testing.T) {

	config := &DBConfig{
		Username: "zhangsan",
		Password: "abcdegf",
		Host:     "127.0.0.1",
		Port:     8088,
		Ttl:      5000,
		MinConn:  5,
		MaxConn:  100,
	}
	fmt.Println(config)
	fmt.Print(config)
	fmt.Println()

	// json.Marshal需要public访问权限
	marshal, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
	}
	// {"username":"zhangsan","password":"aaaaa","supervisor_agent":"127.0.0.1","port":8088,"ttl":5000,"min_conn":5,"max_conn":100}
	fmt.Println(string(marshal))
}
