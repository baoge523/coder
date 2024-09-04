package _interface

import (
	"fmt"
	"testing"
)

type AInter interface {
	A()
	B()
}
type Impl struct {

}

func (i Impl) A() {
	fmt.Println("a")
}

func (i *Impl) B() {
    fmt.Println("b")
}

func print(a AInter) {
	a.A()
}


type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	s.data = str
}

// 使用值接收器的方法既可以通过值调用，也可以通过指针调用。
// 带指针接收器的方法只能通过指针或 addressable values 调用。
func TestInterfacePointer(t *testing.T) {
	sVals := map[int]S{1: {"A"}}

	// 你通过值只能调用 Read
	sVals[1].Read()

	// 这不能编译通过：
	// sVals[1].Write("test")

	sPtrs := map[int]*S{1: {"A"}}

	// 通过指针既可以调用 Read，也可以调用 Write 方法
	sPtrs[1].Read()
	sPtrs[1].Write("test")
}

// 接口检查
func TestInterfaceCheck(t *testing.T) {
	var _ AInter = (*Impl)(nil)  // 检查Impl是否实现了AInter接口的所有方法

	// 指针方法只能指针调用，value方法可以指针调用和对象调用
	// var a AInter = Impl{}   // 会报错，因为B()是基于指针的，所以针对于对象是没有全部实现方法的，所以不能赋值给AInter接口
}