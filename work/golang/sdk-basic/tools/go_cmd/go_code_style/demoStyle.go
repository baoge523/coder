package go_code_style

import "fmt"

func styleFmt() {
	// gofmt -r 'test_a -> testA' -w before
	// test_a := "testA"
	// gofmt -r 'test_a -> testA' -w after
	// testA := "testA"
	testA := "testA"
	fmt.Println(testA)

	fmt.Printf("aa %d,bb %s", 55)

}
