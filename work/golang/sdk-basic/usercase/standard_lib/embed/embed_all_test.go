package embed

import (
	"embed"
	"fmt"
	"os"
	"test/usercase/standard_lib"
	"testing"
	"text/template"
)
import _ "embed"

//go:embed string_test.txt
var content string

//go:embed string_test.txt
var arrayByteContent []byte

//go:embed template.json
var fs embed.FS

func TestEmbedString(t *testing.T) {
	fmt.Printf("embed info string content: %s \n", content)
}

func TestEmbedByte(t *testing.T) {
	fmt.Printf("embed info []byte content: %s \n", string(arrayByteContent))
}

func TestEmbedFS(t *testing.T) {

	cont, err := fs.ReadFile("string_test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("embed info fs string_test content: %s \n", string(cont))

	cont, err = fs.ReadFile("test2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("embed info fs test2 content: %s \n", string(cont))
}

/*
embed FS 和 text/template 搭配使用
注意：.New("template.json")指定的名称需要和加载的文件名一致，建议不指定
*/
func TestEmbedFS_temp(t *testing.T) {

	u := standard_lib.User{
		Name: "andy",
		Age:  18,
	}

	// 这个会报错 "test" is an incomplete or empty template, 原因是ParseFS返回的template是new出来的，根本原因是：名称不一致导致的
	// template.New("test").ParseFS(fs, "template.json")

	// .New("template.json") 可以省略
	temp, err := template.New("template.json").ParseFS(fs, "template.json")
	if err != nil {
		fmt.Printf("parse: %v", err)
		return
	}

	err = temp.Execute(os.Stdout, &u)
	if err != nil {
		fmt.Printf("exec: %v", err)
		return
	}

}
