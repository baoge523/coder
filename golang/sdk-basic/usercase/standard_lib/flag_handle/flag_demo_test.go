package flag_handle

import (
	"flag"
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {

	ageP := flag.Int("age", 18, "usage age: please input your age")

	fmt.Printf("age = %d \n", *ageP)

}
