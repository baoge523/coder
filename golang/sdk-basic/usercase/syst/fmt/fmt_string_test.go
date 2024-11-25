package fmt_all

import (
	"fmt"
	"testing"
)

/**
String and slice of bytes

%s	the uninterpreted bytes of the string or slice
%q	a double-quoted string safely escaped with Go syntax
%x	base 16, lower-case, two characters per byte
%X	base 16, upper-case, two characters per byte
*/

func TestString(t *testing.T) {

	name := "Andy"
	nameByte := []byte(name)  // 将sting转换成byte[]
	fmt.Printf("%s \n", name) // Andy
	fmt.Printf("%q \n", name) // "Andy"
	fmt.Printf("%x \n", name) // 416e6479
	fmt.Printf("%X \n", name) // 416E6479

	// 打印byte数组
	fmt.Printf("%s \n", nameByte) // Andy
	fmt.Printf("%q \n", nameByte) // "Andy"
	fmt.Printf("%x \n", nameByte) // 416e6479
	fmt.Printf("%X \n", nameByte) // 416E6479

}
