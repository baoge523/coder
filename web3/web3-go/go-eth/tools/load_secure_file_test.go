package tools

import (
	"fmt"
	"testing"
)

func TestLoadSecureFile(t *testing.T) {
	got, err := LoadSecureFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(got)
}
