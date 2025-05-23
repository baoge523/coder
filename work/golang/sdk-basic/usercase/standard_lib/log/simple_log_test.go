package log

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestSimpleLog(t *testing.T) {

	logger := log.Default()
	prefix := logger.Prefix()
	fmt.Println(prefix)
	logger.Println("aaaa")

	var buffer bytes.Buffer // 创建一个默认的bytes.Buffer对象，如果要改变对象里面的信息，需要传递地址
	l := log.New(&buffer, "custom-prefix: ", log.Lshortfile)
	l.Printf("this is my log info %s", "normal")
	fmt.Println(&buffer)

	// 根据recover catch zhe
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover panic %v", err)
		}
	}()

	l.Panicf("this is panic info %s", "panic") // 在次之后的代码都不会执行了，但会执行defer操作
	// l.Fatalf("this is fatal info %s", "yes") // os.Exit(1) 会退出程序
	fmt.Println("success")
}

func TestLog(t *testing.T) {
	// 将log对象的的输出设置为标准输出，然后log操作的日志，都会输出到控制台
	log.SetOutput(os.Stdout)
	log.Printf("log set output to os.stdout %s", "hello world")
}
