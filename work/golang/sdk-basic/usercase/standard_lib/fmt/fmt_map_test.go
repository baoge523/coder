package fmt_all

import (
	"fmt"
	"testing"
)

// Maps are printed in a consistent order,
// sorted by the values of the keys.  以key的值排序
func TestFmtMap(t *testing.T) {

	m := make(map[string]string)
	m["f"] = "f"
	m["a"] = "a"
	m["z"] = "z"
	m["d"] = "d"
	fmt.Printf("%v \n", m)  // output: map[a:a d:d f:f z:z]
	fmt.Printf("%+v \n", m) // output: map[a:a d:d f:f z:z]
	fmt.Printf("%#v \n", m) // map[string]string{"a":"a", "d":"d", "f":"f", "z":"z"}
}
