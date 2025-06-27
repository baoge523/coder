package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelClose(t *testing.T) {
	b := check()
	fmt.Println(b)
	time.Sleep(4 * time.Second)
	fmt.Println("success")
}

func check() bool {

	c := make(chan bool, 1)

	go time.AfterFunc(2*time.Second, func() {
		c <- true
		fmt.Println("2s do ok")
	})

	go time.AfterFunc(3*time.Second, func() {
		c <- false
		fmt.Println("3s do ok")
	})

	return <-c

}
