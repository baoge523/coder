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
