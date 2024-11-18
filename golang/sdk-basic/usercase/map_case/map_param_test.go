package map_case

import (
	"fmt"
	"testing"
)


func TestMapParam(t *testing.T) {

	var m = map[string]string {
		"a":"1",
	}
	mapParamPoint(&m)
	fmt.Println(m)
	mapParamNotPoint(m)
	fmt.Println(m)
	// result
	// map[a:1 name:zhang san]
	// map[a:1 age:18 name:zhang san]

}

// 基于map指针，可以通过方法改变其map中的值；注意通过*获取指针的值操作
func mapParamPoint(m *map[string]string) {
	(*m)["name"] = "zhang san"
}

func mapParamNotPoint(m map[string]string) {
	m["age"] = "18"
}



func TestMapDelete(t *testing.T) {
	var m = make(map[string]string)

	m["name"] = "lisi"
	m["age"] = "18"

	delete(m,"name")
	fmt.Println(m)

}