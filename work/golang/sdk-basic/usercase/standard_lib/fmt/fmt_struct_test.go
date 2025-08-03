package fmt_all

import (
	"fmt"
	"testing"
)

// 结论通过%v %+v 都无法打印出指针类型的值，指针类型只能打印其地址
func TestFmtStruct(t *testing.T) {
	user := User{
		Name: "Andy",
		Age:  18,
		Look: &Look{
			Face: "beautiful",
		},
	}

	fmt.Printf("%v \n", user)  // output: {Andy 18}
	fmt.Printf("%+v \n", user) // output: {Name:Andy Age:18}
	fmt.Printf("%#v \n", user) // output: fmt_all.User{Name:"Andy", Age:18}
	fmt.Println("-----------")
	fmt.Printf("%v \n", &user)  // output: &{Andy 18}
	fmt.Printf("%+v \n", &user) // output: &{Name:Andy Age:18}
	fmt.Printf("%#v \n", &user) // output: &fmt_all.User{Name:"Andy", Age:18}
}
