package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{7, 8, 9, 3, 4, 5}

	lenA := len(a) - 1
	lenB := len(b) - 1

	total := lenA + lenB

	i := 0
	j := 0
	target := -1
	for i <= total {
		var valA, valB int
		if i <= lenA {
			valA = a[i]
		} else {
			valA = b[i-lenA]
		}

		if j <= lenB {
			valB = b[j]
		} else {
			valB = a[j-lenB]
		}
		i++
		j++
		if valA == valB {
			target = valA
			break
		}

	}
	fmt.Println(target)
}
