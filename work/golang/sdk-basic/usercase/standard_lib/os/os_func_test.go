package os

import (
	"fmt"
	"os"
	"testing"
)

func TestFuncOSExpand(t *testing.T) {
	text := os.Expand("firstname = ${FirstName} , lastname = $LastName", func(placeholder string) string {
		switch placeholder {
		case "FirstName":
			return "dave"
		case "LastName":
			return "andy"
		default:
			return "--"
		}
	})
	fmt.Println(text)

	os.Setenv("FirstName", "hello")
	os.Setenv("LastName", "world")
	text2 := os.ExpandEnv("firstname = ${FirstName} , lastname = $LastName")
	fmt.Println(text2)
}

func TestFuncOSEnv(t *testing.T) {

	FirstName := os.Getenv("FirstName")
	fmt.Printf("firstname:%s \n", FirstName)
	os.Setenv("FirstName", "hello")
	FirstName = os.Getenv("FirstName")
	fmt.Printf("firstname:%s \n", FirstName)

	getegid := os.Getegid()
	getgid := os.Getgid()
	fmt.Printf("getegid = %d \n", getegid)
	fmt.Printf("getgid = %d \n", getgid)

	getpagesize := os.Getpagesize()
	fmt.Printf("getpagesize = %d \n", getpagesize) // 16384 = 16k

}
