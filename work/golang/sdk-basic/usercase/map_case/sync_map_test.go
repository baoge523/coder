package map_case

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	m := sync.Map{}
	m.Store("k1", "v1")
	if value, ok := m.Load("k11"); !ok {
		fmt.Println("not ok")
	} else {
		fmt.Println(value)
	}
}

func TestEmptyMap(t *testing.T) {
	var m map[string]string
	if s, ok := m["aaa"]; !ok {
		fmt.Println("not ok")
	} else {
		fmt.Println(s)
	}

}

func TestMap(t *testing.T) {
	m := make(map[string]string, 10)
	m["aa"] = "bb"

}
