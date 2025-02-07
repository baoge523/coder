package fmt_all

import (
	"fmt"
	"testing"
)

/**
%p	base 16 notation, with leading 0x
The %b, %d, %o, %x and %X verbs also work with pointers,
formatting the value exactly as if it were an integer.
*/

func TestPoint(t *testing.T) {
	userPoint := &User{
		Name: "Andy",
		Age:  18,
	}

	fmt.Printf("%p \n", userPoint) // 0x14000130090 指针地址
}

func TestPointOther(t *testing.T) {
	num := 123
	numPoint := &num

	fmt.Printf("%d \n", num) // %d 打印int
	fmt.Printf("%d \n", numPoint)  // 使用%d打印int指针
}
