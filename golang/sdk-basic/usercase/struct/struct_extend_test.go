package _struct

import (
	"encoding/json"
	"fmt"
	"testing"
)

type BasicParam struct {
	ID  string
	UIN string
}

type RequestParam struct {
	BasicParam BasicParam `json:"basic_param"`
	Name       string     `json:"Name"`
}

func TestExtend(t *testing.T) {

	req := RequestParam{
		BasicParam: BasicParam{
			ID:  "111",
			UIN: "222",
		},
		Name: "aaa",
	}
	marshal, _ := json.Marshal(&req)

	fmt.Printf("req %s \n", string(marshal))
	fmt.Printf("req %+v \n", req)
}
