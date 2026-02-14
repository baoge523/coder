package _struct

/**
  实践研究一下 struct 和 point 本质上有什么区别

  已经知道 struct对象方法，可以被对象和指针使用，但是指针方法只能被指针使用

*/

import (
	"fmt"
	"testing"
)

type User struct {
	Name    string
	Age     int
	Hobbies []string
}

func (u User) ChangeName(name string) {
	u.Name = name
}

func (u *User) ResetName(name string) {
	u.Name = name
}

func TestName(t *testing.T) {

	u := User{
		Name: "test",
	}
	fmt.Println(u)
	u.ChangeName("test1")
	fmt.Println(u)
	u.ResetName("test2")
	fmt.Println(u)

}
