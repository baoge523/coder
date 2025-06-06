package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"testing"

	_ "net/http/pprof"
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
	defer pprof.StopCPUProfile()

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
	// memory2
	memFile2, err := os.Create("mem2.prof")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = pprof.WriteHeapProfile(memFile2)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer memFile2.Close()


}

// net/http/pprof
func TestHttp(t *testing.T) {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

}
