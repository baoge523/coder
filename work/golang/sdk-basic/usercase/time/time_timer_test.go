package time

import (
	"fmt"
	"testing"
	"time"
)

func TestAfterFunc(t *testing.T) {

	timer := time.AfterFunc(time.Second, func() {
		fmt.Println("func done")
	})
	timer.Stop() // cancel the func to call, but cannot use timer.C,because it is nil

	fmt.Println(timer.C) // <nil>

	time.Sleep(5* time.Second)
	fmt.Println("main done")

}
