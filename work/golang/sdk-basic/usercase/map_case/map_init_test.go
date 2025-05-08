package map_case

import (
	"fmt"
	"testing"
)

type A struct {
	M    map[string]string
	Name string
}

func TestMapInit(t *testing.T) {

	a := A{
		M: make(map[string]string),
	}
	a.M["aa"] = "aa"
	fmt.Println(a)
}
