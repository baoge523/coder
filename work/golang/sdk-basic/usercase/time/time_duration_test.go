package time

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {

	//ti := time.Now()
	//time.Sleep(time.Second)
	//since := time.Since(ti)  // now - ti --> now().Sub(ti)
	//fmt.Println(since)

	ti := time.Now()
	ti = ti.Add(time.Duration(5) * time.Second) // add 需要返回时间才对
	time.Sleep(time.Second)
	until := time.Until(ti) // ti - now
	fmt.Printf("Duration until future time: %.0f seconds\n", math.Ceil(until.Seconds()))


	futureTime := time.Now().Add(5 * time.Second)
	durationUntil := time.Until(futureTime)
	fmt.Printf("Duration until future time: %.0f seconds\n", math.Ceil(durationUntil.Seconds()))

}



func testParse() {
	hours, _ := time.ParseDuration("10h")
	complex, _ := time.ParseDuration("1h10m10s")
	micro, _ := time.ParseDuration("1µs")
	// The package also accepts the incorrect but common prefix u for micro.
	micro2, _ := time.ParseDuration("1us")

	fmt.Println(hours)
	fmt.Println(complex)
	fmt.Printf("There are %.0f seconds in %v.\n", complex.Seconds(), complex)
	fmt.Printf("There are %d nanoseconds in %v.\n", micro.Nanoseconds(), micro)
	fmt.Printf("There are %6.2e seconds in %v.\n", micro2.Seconds(), micro2)
}
