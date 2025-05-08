package runtime

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMethod(t *testing.T) {

	goroutine := runtime.NumGoroutine()
	fmt.Printf("current goroutine num = %d \n",goroutine)
	version := runtime.Version()
	fmt.Printf("version = %s \n",version)
	cpu := runtime.NumCPU()
	fmt.Printf("current cpu num = %d \n",cpu)
}
