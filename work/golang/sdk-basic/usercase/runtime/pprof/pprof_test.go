package pprof

import (
	"fmt"
	"os"
	"runtime/pprof"
	"testing"
)

/**
内部提供了以下的pprof信息
allocs     // allocs memory
block
goroutine
heap      // head memory
mutex
threadcreate
*/
func TestMethod(t *testing.T) {
	profiles := pprof.Profiles()
	for _, prof := range profiles {
		fmt.Printf("%+v \n",prof.Name())
	}
}
// 指定创建pprof文件的方式
// - go test 的方式
// - main方法通过pprof
// - net/http/pprof
func TestPprof(t *testing.T) {

	// go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

	// cpu
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = pprof.StartCPUProfile(cpuFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// memory
	memFile, err := os.Create("mem.prof")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = pprof.Lookup("allocs").WriteTo(memFile, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// other



	// net/http/pprof
}
