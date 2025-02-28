package func_and_method

import (
	"fmt"
	"testing"
)

/**
在golang中，函数的参数传递是通过值传递的，值传递就是通过拷贝数据来达到传参的效果
*/

func TestParam(t *testing.T) {
	u := User{
		Name: "zhangsan",
		Age:  11,
	}
	/**
	对于参数u,在调用changeParam时，传入的u会被拷贝一份
	*/
	fmt.Printf("%p \n", &u)
	changeParam(u)
	fmt.Printf("%s \n", &u)

	changeParamPoint(&u)

	fmt.Printf("%s \n", &u)

}

type User struct {
	Name string
	Age  int
}

/*
*
1、当String()方法属于对象的时候，需要 fmt.Printf("%s \n", u)

2、当String()方法属于指针的时候，需要 fmt.Printf("%s \n", &u)
*/
func (u *User) String() string {
	return fmt.Sprintf("Name = %s, Age = %d", u.Name, u.Age)
}

func changeParam(u User) {
	fmt.Printf("changeParam u address = %p \n", &u)
	u.Name = "change"
	u.Age = 16
}

func changeParamPoint(u *User) {
	// 打印指针指向的地址
	fmt.Printf("changeParamPoint u address = %p \n", u)
	// 打印指针本身的地址
	fmt.Printf("changeParamPoint u address = %p \n", &u)
	u.Name = "changePoint"
	u.Age = 18
}

// golang 中，在同一个package下不允许出现相同的函数名，即便函数的参数个数或者类型不一样，都是不行的
//func changeParam(u User,name string) {
//
//}

/*
*
测试验证数组的传参
*/
func TestParmaArray(t *testing.T) {

	fmt.Println("----- 这是数组的测试验证--------")
	arr := [3]int{1, 2, 3}
	fmt.Printf("%p \n", &arr)
	checkArraySimple(arr)
	fmt.Printf("%p \n", &arr)

	fmt.Println("----- 这是切片的测试验证--------")
	arr2 := []int{1, 2, 3}
	fmt.Printf("%p \n", arr2)
	checkArraySlice(arr2)
	fmt.Printf("%p \n", arr2)

	fmt.Println("----- 这是数组指针的测试验证--------")
	fmt.Printf("%p \n", &arr)
	checkArraySimplePoint(&arr)
	fmt.Printf("%p \n", &arr)

}

// 这种没有指定数组长度的是切片,切片本身就是指针
func checkArraySlice(arr []int) {
	fmt.Printf("%p \n", arr)
}

// 这种指定了长度的是数组
func checkArraySimple(arr [3]int) {
	fmt.Printf("%p \n", &arr)
}

// 这种指定了长度的是数组，这里传递的是数组指针，避免了数据的拷贝，指拷贝了指针
func checkArraySimplePoint(arr *[3]int) {
	fmt.Printf("%p \n", arr)
}
