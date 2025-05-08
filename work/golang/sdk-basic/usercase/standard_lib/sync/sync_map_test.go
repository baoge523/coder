package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	cache := sync.Map{}
	cache.Store("k1","k1")
	cache.Store("k2",123)
	cache.Store("k3",false)

	if value, ok := cache.Load("k2"); ok {
		fmt.Println(value)
	}

	previous, loaded := cache.Swap("k4", "k4")

	fmt.Printf("k4 not exists: old = %v, isExists = %t \n",previous,loaded)

	previous, loaded = cache.Swap("k1", "k4")

	fmt.Printf("k4 exists: old = %v, isExists = %t \n",previous,loaded)

	if value, ok := cache.Load("k1"); ok {
		fmt.Println(value)
	}

	deleted := cache.CompareAndDelete("k3", false)
	fmt.Println(deleted)
	if value, ok := cache.Load("k3"); ok {
		fmt.Println(value)
	}

}

func TestRange(t *testing.T) {
	cache := sync.Map{}
	cache.Store("k1","k1")
	cache.Store("k2",123)
	cache.Store("k3",false)

	cache.Range(func(key, value any) bool {
		fmt.Println(key)
		if key == "k2" {
			return false
		}
		return true
	})
}
