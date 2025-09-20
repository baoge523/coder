package stack_exe

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	a := "(){}[)]"

	var stack []string
	m1 := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	m2 := map[string]struct{}{
		"(": struct{}{},
		"{": struct{}{},
		"[": struct{}{},
	}
	result := false
	for _, v := range a {
		if _, ok := m2[string(v)]; ok {
			stack = append(stack, string(v))
		} else {
			if val, ok := m1[string(v)]; ok {
				if len(stack) == 0 || val != string(stack[len(stack)-1]) {
					break
				}
				stack = stack[:len(stack)-1]

			} else {
				break
			}
		}
	}
	if len(stack) == 0 {
		result = true
	}

	fmt.Println(result)

}
