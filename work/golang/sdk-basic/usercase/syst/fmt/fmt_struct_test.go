package fmt_all

import (
	"fmt"
	"testing"
)

func TestFmtStruct(t *testing.T) {
	user := User{
		Name: "Andy",
		Age:  18,
	}

	fmt.Printf("%v \n", user)  // output: {Andy 18}
	fmt.Printf("%+v \n", user) // output: {Name:Andy Age:18}
	fmt.Printf("%#v \n", user) // output: fmt_all.User{Name:"Andy", Age:18}
	fmt.Println("-----------")
	fmt.Printf("%v \n", &user)  // output: &{Andy 18}
	fmt.Printf("%+v \n", &user) // output: &{Name:Andy Age:18}
	fmt.Printf("%#v \n", &user) // output: &fmt_all.User{Name:"Andy", Age:18}
}
