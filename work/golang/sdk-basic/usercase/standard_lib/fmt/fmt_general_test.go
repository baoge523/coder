package fmt_all

import (
	"fmt"
	"testing"
)

/**
%v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value
%T	a Go-syntax representation of the type of the value
%%	a literal percent sign; consumes no value
*/

type Look struct {
	Face string
}

type User struct {
	Name string
	Age  int
	Look *Look
}

// 验证%v 和 %+v
func TestGeneral(t *testing.T) {

	user := &User{
		Name: "Andy",
		Age:  18,
	}

	fmt.Printf("%v \n", *user)  // {Andy 18}
	fmt.Printf("%v \n", user)   // &{Andy 18}
	fmt.Printf("%+v \n", *user) // {Name:Andy Age:18}
	fmt.Printf("%+v \n", user)  // &{Name:Andy Age:18}
}

// 验证 %#v   package.struct{field:value,...}
func TestGeneral2(t *testing.T) {
	user := &User{
		Name: "Andy",
		Age:  18,
	}

	fmt.Printf("%#v \n", *user) // fmt_all.User{Name:"Andy", Age:18}
	fmt.Printf("%#v \n", user)  // &fmt_all.User{Name:"Andy", Age:18}

}

// 验证 %T   只打印类型
func TestGeneral3(t *testing.T) {
	user := &User{
		Name: "Andy",
		Age:  18,
	}

	fmt.Printf("%T \n", *user) // fmt_all.User   结构体类型
	fmt.Printf("%T \n", user)  // *fmt_all.User   结构体指针类型

}

// 验证 %%   不消费任何占位符
func TestGeneral4(t *testing.T) {
	fmt.Printf("%% fmt general \n") // % fmt general
}

func TestGeneralString(t *testing.T) {

	appid := "1251763868"
	appid1 := 1251763868

	fmt.Printf("%v \n", appid)
	fmt.Printf("%v \n", appid1)

}
