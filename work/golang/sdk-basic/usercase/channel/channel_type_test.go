package channel

import (
	"fmt"
	"testing"
)

func TestChannelType(t *testing.T) {

	c1 := make(chan int)
	var c2 chan int
	needChannelType(c1)
	needChannelType(c2)

}

func needChannelType(c chan int) {
	if c == nil {
		fmt.Printf("chan is nil v %v \n", c) // <nil>
		fmt.Printf("chan is nil p %p \n", c) // 0x0
		return
	}
	fmt.Printf("chan is not nil p %p \n", c) // 0x14000020300
	fmt.Printf("chan is not nil p %v \n", c) // 0x14000020300
}
