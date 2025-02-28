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

	// 这样能编译过, 原因是，s是一个可寻址的地址，可以通过其指针调用Write方法
	s := sVals[1]
	s.Write("aaa") // 因为是编译器将 s 优化成 (&s).Write("aaa")

	// 这里编译不过？ 为什么 原因是sVals[1] 是一个值，不是地址，所以会报错
	// sVals[1].Write("aaa")  // 编译器无法将 sVals[1] 优化成 &sVals[1]

	sPtrs := map[int]*S{1: {"A"}}

	// 通过指针既可以调用 Read，也可以调用 Write 方法
	sPtrs[1].Read()
	sPtrs[1].Write("test")
}

func TestStructAndPoint(t *testing.T) {
	// 这是结构体对象
	s := S{data: "s"}
	s.Read()
	// 先找对象方法，再找指针方法，然后将s 优化成 &s
	s.Write("test") // 编译器优化 (&s).Write("test")

	// 这是结构体指针
	sp := &S{data: "s point"}
	// 先找指针方法，再找对象方法，将 sp 优化成 *sp
	sp.Read() // 编译器优化  (*sp).Read()
	sp.Write("point")
}

// 接口检查
func TestInterfaceCheck(t *testing.T) {
	var _ AInter = (*Impl)(nil) // 检查Impl是否实现了AInter接口的所有方法

	// 指针方法只能指针调用，value方法可以指针调用和对象调用
	// var a AInter = Impl{}   // 会报错，因为B()是基于指针的，所以针对于对象是没有全部实现方法的，所以不能赋值给AInter接口
	// var b AInter = &Impl{}  // ok
}

type Duck interface {
	Quack()
}

type Cat1 struct{}

func (c Cat1) Quack() {
	fmt.Println("meow")
}

type Cat2 struct{}

func (c *Cat2) Quack() {
	fmt.Println("meow")
}

// 当我们使用指针实现接口时，只有指针类型的变量才会实现该接口；
// 当我们使用结构体实现接口时，指针类型和结构体类型都会实现该接口
func TestInvoke(t *testing.T) {

	fmt.Println("----- 对象方法赋值给接口-----")
	var c1 Duck = Cat1{}
	c1.Quack()

	var c2 Duck = &Cat1{}
	c2.Quack()

	fmt.Println("----- 指针方法赋值给接口-----")
	// 编译报错
	//var c3 Duck = Cat2{} // 指针方法 将对象赋值给接口，会报接口没有实现
	//c3.Quack()

	var c4 Duck = &Cat2{}
	c4.Quack()

}
