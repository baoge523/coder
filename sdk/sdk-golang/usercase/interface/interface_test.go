package _interface

import (
	"fmt"
	"testing"
)


type Animal interface {
    Eat()
}

type AdapterAnimal struct {
	name string
}

// 指针的方法: 只支持指针访问
func (a *AdapterAnimal) Eat()  {
	fmt.Println("AdapterAnimal Eat")
}

// 对象的方法：支持指针和对象访问
func (a AdapterAnimal) Have() {
	fmt.Println("AdapterAnimal Have")
}

type Cat struct {
	*AdapterAnimal // 继承  基于指针的方式继承 继承方法和属性，但是属性的访问使用指针来的，所以需要先初始化指针  同时不能使用对象定义的方法
}
// Eat 方法重写
func (c *Cat) Eat() {
	c.AdapterAnimal.Eat()
	fmt.Println("Cat Eat")
}

func (c Cat) Have() {
    c.AdapterAnimal.Have()  // 通过指针的方式访问了对象的方法，会报错；但是通过对象的方式并初始化了AdapterAnimal指针就可以访问
    fmt.Println("Cat Have")
}


type Dog struct {
    AdapterAnimal   // 继承其所有方法和属性 -- 包括指针、对象的方法
}



// 结论: 指针只能访问指针定义的方法，
//      对象可以访问指针和对象定义的方法,但对象中如果使用了指针，需要先初始化指针
func TestInterfaceOverride(t *testing.T) {

	//cat := new(Cat)
	//cat.Eat()


	//cat2 := Cat{ &AdapterAnimal{name: "cat2"}}
	//cat2.Eat()
	//cat2.Have()

	dog := new(Dog)
	dog.Eat()  // dog.AdapterAnimal.Eat()  等价
	dog.Have() // dog.AdapterAnimal.Have()  等价

	dog2 := Dog{}
	dog2.Eat()
	dog2.Have()

	dog3 := Dog{AdapterAnimal: AdapterAnimal{name: "dog3"}}
	dog3.Eat()
	dog3.Have()

}
