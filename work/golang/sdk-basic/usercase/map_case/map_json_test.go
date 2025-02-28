package map_case

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapPoint(t *testing.T) {
	map1 := make(map[string]string)
	map1["aa"] = "aa"
	map1["bb"] = "bb"
	mapPoint(&map1)
	mapNotPoint(map1)
}

// json.Marshal 支持map指针,和map
// json.Unmarshal 需要传递map指针，如果直接是map赋值失败
func mapPoint(m *map[string]string) {
	marshal, err := json.Marshal(m) // 这里使用 *m 也是可以的
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println("ok1")
		fmt.Println(string(marshal))
	}
	var copyMap = make(map[string]string)
	err = json.Unmarshal(marshal, &copyMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("copy map 1")
	fmt.Println(copyMap)
}

func mapNotPoint(m map[string]string) {
	marshal, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println("ok2")
		fmt.Println(string(marshal))
	}
	var copyMap = make(map[string]string)
	err = json.Unmarshal(marshal, &copyMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("copy map 2")
	fmt.Println(copyMap)
}

// 测试验证将map转成json对象，并以json的方式返回
func TestJSONMap(t *testing.T) {

	a1 := []string{"aaa"}
	fmt.Printf("%+v \n", a1)

	//m1 := make(map[string]string)
	//m1["aa"] = "aa"
	//m1["bb"] = "bb"
	//m2 := make(map[string]interface{})
	//bytes, _ := json.Marshal(m1)
	//fmt.Printf("%s \n", string(bytes))
	//m2["object"] = m1
	//m2["test"] = "test"
	//
	//marshal, _ := json.Marshal(m2)
	//fmt.Printf("%s \n", string(marshal))

}
