package fmt_all

import (
	"errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {

	const name, age = "Andy", 18

	err := errors.New("new error")

	err = fmt.Errorf("error: name = %s, age =%d, err = %w", name, age, err)

	fmt.Printf("%v \n",err)


}
