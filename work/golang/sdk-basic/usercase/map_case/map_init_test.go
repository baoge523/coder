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
	a.M["aa"] = "aabb"
	fmt.Println(a)
	a.M["bb"] = "bb"
	fmt.Println(a)
	delete(a.M, "bb")
	fmt.Println(a)

}
