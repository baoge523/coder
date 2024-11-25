package fmt_all

import (
	"fmt"
	"testing"
)

/**
%t	the word true or false
 */

func TestBool(t *testing.T) {
	flag_1 := true
	flag_2 := false

	fmt.Printf("flag_1 = %t and flag_2 = %t \n", flag_1, flag_2)

}
