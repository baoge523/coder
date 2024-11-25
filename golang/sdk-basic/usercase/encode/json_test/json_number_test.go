package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

/**
  json.Unmarshal 会对长的number类型的数字进行科学计数法处理，导致业务异常
  解决方式
     使用json.encoder的方式
*/

const (
	info1 = "{\"appid\":1251763868,\"city\":\"shhqcft\",\"instanceid\":\"1c88ee9e-a575-11ef-88e0-246e96755480\",\"insttype\":\"master\",\"projectid\":0,\"uInstanceId\":\"cdb-gfbwgjrc\"}"
	info2 = "{\"appid\":\"1251763868\",\"city\":\"shhqcft\",\"instanceid\":\"1c88ee9e-a575-11ef-88e0-246e96755480\",\"insttype\":\"master\",\"projectid\":0,\"uInstanceId\":\"cdb-gfbwgjrc\"}"
)

func TestUnmarshal(t *testing.T) {

	var infoMap map[string]interface{}
	if err := json.Unmarshal([]byte(info1), &infoMap); err != nil {
		fmt.Printf("%v \n", err)
		return
	}
	// map[appid:1.251763868e+09 city:shhqcft instanceid:1c88ee9e-a575-11ef-88e0-246e96755480 insttype:master projectid:0 uInstanceId:cdb-gfbwgjrc]
	fmt.Printf("%v \n", infoMap)

}

// 通过json的Decoder的方式告诉其不做科学计数法转换
func TestDecoder(t *testing.T) {
	var infoMap map[string]interface{}

	decoder := json.NewDecoder(bytes.NewBufferString(info1))
	// 表示大number，不使用科学计数法
	decoder.UseNumber()

	if err := decoder.Decode(&infoMap); err != nil {
		fmt.Printf("%v \n", err)
		return
	}
	// map[appid:1251763868 city:shhqcft instanceid:1c88ee9e-a575-11ef-88e0-246e96755480 insttype:master projectid:0 uInstanceId:cdb-gfbwgjrc]
	fmt.Printf("%v \n", infoMap)
}
