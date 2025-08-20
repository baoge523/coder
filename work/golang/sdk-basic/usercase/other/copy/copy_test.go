package copy

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/mohae/deepcopy"
	"testing"
)

type A struct {
	Name string
	BB   *B
}

type B struct {
	Name string
}

func TestCopy(t *testing.T) {

	b := B{
		Name: "bb",
	}

	a := A{Name: "aa", BB: &b}

	// 数组的拷贝
	// copy()

	// 深拷贝
	aa := deepcopy.Copy(a)

	if a2, ok := aa.(A); !ok {
		fmt.Println("error")
		return
	} else {
		a2.Name = "a2"
		a2.BB.Name = "b2"
		fmt.Printf("a2 = %+v\n", a2)
	}
	fmt.Printf("a = %+v\n", a)

	var ac A
	if err := copier.CopyWithOption(&ac, a, copier.Option{DeepCopy: false}); err != nil {
		fmt.Println("copier error")
		return
	}
	ac.BB.Name = "bac"
	fmt.Printf("ac = %v\n", ac)
	fmt.Printf("a = %v\n", a)

}
