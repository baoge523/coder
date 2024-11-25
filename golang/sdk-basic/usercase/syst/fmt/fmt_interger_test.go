package fmt_all

import (
	"fmt"
	"testing"
)

/**
%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%O	base 8 with 0o prefix
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"
*/

func TestInteger(t *testing.T) {
	number := 100

	// 打印类型
	fmt.Printf("%T \n", number)

	fmt.Printf("%b \n", number)  // 1100100  二进制
	fmt.Printf("%c \n", number)  // d  unicode code
	fmt.Printf("%d \n", number)  // 100
	fmt.Printf("%o \n", number)  // 144
	fmt.Printf("%O \n", number)  // 0o144
	fmt.Printf("%q \n", number)  // 'd'
	fmt.Printf("%x \n", number)  // 64  16进制带a-f
	fmt.Printf("%X \n", number)  // 64  16进制带A-F
	fmt.Printf("%U \n", number)  // U+0064

}
