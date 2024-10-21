package slice

import (
	"fmt"
	"testing"
)

// 可变参数

func TestVariableParam(t *testing.T) {
	useRange()    // 调用ok
	useRange(nil) // 报空指针异常，会将nil当成对象调用
}

func useRange(arr ...A) {

	for _, s := range arr {
		fmt.Printf("----- %s", s())
	}

}

type A func() string
