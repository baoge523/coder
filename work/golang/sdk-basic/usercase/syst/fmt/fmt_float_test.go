package fmt_all

import (
	"fmt"
	"testing"
)

/**
%b	decimalless scientific notation with exponent a power of two,
	in the manner of strconv.FormatFloat with the 'b' format,
	e.g. -123456p-78
%e	scientific notation, e.g. -1.234456e+78
%E	scientific notation, e.g. -1.234456E+78
%f	decimal point but no exponent, e.g. 123.456
%F	synonym for %f
%g	%e for large exponents, %f otherwise. Precision is discussed below.
%G	%E for large exponents, %F otherwise
%x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
%X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

*/

func TestFloat(t *testing.T) {
	num1 := float32(3.14)
	num2 := float64(3.14)
	fmt.Printf("%g \n", num1)
	fmt.Printf("%g \n", num2)
}

func TestFloat1(t *testing.T) {
	num1 := float32(3.14)
	fmt.Printf("%f \n", num1) // 3.140000
	fmt.Printf("%3.2f \n", num1) // 3.14
	fmt.Printf("%3.3f \n", num1) // 3.140
	fmt.Printf("%4.3f \n", num1) // 3.140

}
