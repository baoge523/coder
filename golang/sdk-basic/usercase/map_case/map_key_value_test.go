package map_case

import (
	"fmt"
	"testing"
)

type Value struct {
	name string
}

// 指针定义了，如果没有显示赋值，那么默认为nil，使用会报错
func TestMapKeyValue(t *testing.T) {

	values := make(map[string]*Value, 0)

	// 这里因为map的value是一个指针，没有赋值时，直接使用会报空指针异常
	// invalid memory address or nil pointer dereference
	name := values["aaa"].name
	fmt.Printf("%s \n", name)
}

// 非指针的value，会初始化默认值
func TestMapKeyNotPointValue(t *testing.T) {
	values := make(map[string]Value, 0)
	name := values["aaa"].name
	t.Log(fmt.Sprintf("name = %s \n", name))
}

func TestMapGetKey(t *testing.T) {
	values := make(map[string]Value, 0)
	if _, ok := values["aaa"]; ok {
		fmt.Printf("key = aaa exists")
	} else {
		fmt.Printf("key = aaa  not exists")
	}

}
