package map_case

import (
	"fmt"
	"testing"
)

type Value struct {
	name string
}

func TestMapKeyValue(t *testing.T) {

	values := make(map[string]*Value, 0)

	name := values["aaa"].name
	fmt.Printf("%s \n", name)
}
