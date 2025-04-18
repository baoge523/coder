package fmt_all

import (
	"fmt"
	"testing"
)

func TestFmtType(t *testing.T) {

	num := 10
	fmt.Printf("%T %T \n", num, &num) // output: int *int

	c := make(chan string)
	fmt.Printf("%T %T \n", c, &c) // output: chan string *chan string

	m := make(map[string]string)
	fmt.Printf("%T %T \n", m, &m) // output: map[string]string *map[string]string

}
