package fmt_all

import (
	"fmt"
	"testing"
)

/**
%p	address of 0th element in base 16 notation, with leading 0x
*/

func TestSlice(t *testing.T) {

	arr := make([]string, 0)
	arr = append(arr, "a")
	arr = append(arr, "b")
	arr = append(arr, "c")
	fmt.Printf("%v \n", arr)  // output: [a b c]
	fmt.Printf("%+v \n", arr) // output: [a b c]
	fmt.Printf("%#v \n", arr) // []string{"a", "b", "c"}
	fmt.Printf("%p \n", arr)  // 0x1400007a8c0 打印切片中第一个元素的地址

}
