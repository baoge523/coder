package math

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func TestRand(t *testing.T) {

	intn := rand.Intn(10) // [0,n)
	fmt.Println(intn)

	perm := rand.Perm(10) // [0,n) array  rand per one
	fmt.Printf("%v \n", perm)

	words := strings.Fields("ink runs from the corners of my mouth")
	fmt.Printf("shuffle: before %v \n", words)
	rand.Shuffle(len(words), func(i, j int) { // shuffle means rand order for each exec
		words[i], words[j] = words[j], words[i] // change i to j
	})
	// shuffle: after [from corners the of runs my mouth ink]
	// shuffle: after [corners ink my mouth of runs the from]
	fmt.Printf("shuffle: after %v \n", words)

}
