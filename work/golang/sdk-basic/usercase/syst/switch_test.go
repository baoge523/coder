package syst

import (
	"fmt"
	"testing"
)

// 验证switch case:的执行流程
func TestSwitch(t *testing.T) {

	num := []string{"0","1","2","-1"}

	for _, s := range num {
		switch s {
		case "0":
		case "1":
		default:
			fmt.Printf("default %s ", s)
		}
		fmt.Println("do anther things")
	}



}