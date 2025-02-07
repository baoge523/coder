package error_handle

import (
	"fmt"
	"testing"
)

func TestErrorMessage(t *testing.T) {
	err := fmt.Errorf("this is err: %s", "yes")
	fmt.Printf("current: %s \n",err.Error())
}


