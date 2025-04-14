package template_handle

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"text/template"
)

// 自定义的函数和系统提供的函数
// step1: 定义 FuncMap 对象
// step2: 注册到template中

func TestFunc1(t *testing.T) {

	// 定义 FuncMap对象
	funcMap := template.FuncMap{
		"add": Add,
	}

	param := map[string]int{
		"num1": 10,
		"num2": 20,
	}

	constStr := `totalNum: {{ add .num1 .num2}}`
	// Funcs(funcMap) 注册到template中
	parse, err := template.New("test").Funcs(funcMap).Parse(constStr)
	if err != nil {
		fmt.Printf("test template error %v", err)
	}
	err = parse.Execute(os.Stdout, param)
	if err != nil {
		fmt.Printf("test template error %v", err)
	}
	fmt.Println()
	fmt.Println("success")
}

func TestFunc2(t *testing.T) {

	funcMap := template.FuncMap{
		"add": Add,
	}
	// use $声明变量后，可以在后面使用
	constStr := `totalNum:{{$num1 := index . 0}}{{$num2 := index . 1}}    {{- add $num1 $num2}}`

	parse, err := template.New("test").Funcs(funcMap).Parse(constStr)
	if err != nil {
		fmt.Printf("test template error %v", err)
	}
	err = parse.Execute(os.Stdout, []int{11, 22})
	if err != nil {
		fmt.Printf("test template error %v", err)
	}
	fmt.Println()
	fmt.Println("success")
}

func TestFunc3(t *testing.T) {

	// 使用系统提供的函数
	funcMap := template.FuncMap{
		"toUp": strings.ToUpper,
	}

	// use $声明变量后，可以在后面使用
	constStr := `first name: {{toUp .}}`

	parse, err := template.New("test").Funcs(funcMap).Parse(constStr)
	if err != nil {
		fmt.Printf("test template error %v", err)
	}
	err = parse.Execute(os.Stdout, "andy")
	if err != nil {
		fmt.Printf("test template error %v", err)
	}
	fmt.Println()
	fmt.Println("success")
}

func Add(a, b int) int {
	return a + b
}
