package map_case

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapPoint(t *testing.T) {

	mapPoint := make(map[string]string)

	mapPoint["aa"] = "aa"
	mapPoint["bb"] = "bb"

	aaa(&mapPoint)

}

func aaa(m *map[string]string) {
	marshal, err := json.Marshal(m) // 这里使用 *m 也是可以的
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println("ok")
		fmt.Println(string(marshal))
	}
}
