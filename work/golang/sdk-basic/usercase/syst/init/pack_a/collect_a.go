package pack_a

import "fmt"

var bbb string

func init()  {
	fmt.Println("pack_a init")
	fmt.Println(bbb)
	bbb = "lisi"
}

func init()  {
	fmt.Println("pack_a init22")
	fmt.Println(bbb)
}

var name = func() string {

	fmt.Println("do func name")

	return "zhangsan"
}