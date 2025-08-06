package _interface

import (
	"fmt"
	"testing"
)

/**

验证golang中的组合关系(继承) interface 和 struct

 interface interface   "继承方法"
 interface struct    --- 不能操作
 struct interface  "继承了接口的方法，但是如果在初始化struct时没有传入interface的具体实现，那么就会报panic"
 struct struct   "继承了其方法"
*/

// 用于检查Swallow 是否实现了Life的所有接口，如果没有，就会编译的时候报错
var _ Life = (*Swallow)(nil)

type Runner interface {
	Run()
}

type Flyer interface {
	Fly()
}

type Bird interface {
	Runner
	Flyer
}

type BirdBase struct {
}

func (bb *BirdBase) Run() {
	fmt.Println("base run")
}

func (bb *BirdBase) Fly() {
	fmt.Println("base fly")
}

// ------------
type Eater interface {
	Eat()
}

type EatBase struct {
}

func (e *EatBase) Eat() {
	fmt.Println("base eat")
}

// ----------
type Life interface {
	Run()
	Fly()
	Eat()
}

// Swallow 燕子
type Swallow struct {
	Bird
	EatBase
}

func TestIface(t *testing.T) {

}

func TestStruct(t *testing.T) {
	s := Swallow{Bird: &BirdBase{}} // 依赖于 Bird interface,在初始化的时候，不是必填，但是使用Bird中的接口时，会报panic错误：invalid memory address or nil pointer dereference
	s.Run()
	s.Fly()
	s.Eat()
	fmt.Println("over")
}
