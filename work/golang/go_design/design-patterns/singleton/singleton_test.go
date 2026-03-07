package main

import (
	"sync"
	"testing"
)

func TestSingletonConcurrency(t *testing.T) {
	const goroutines = 100
	var wg sync.WaitGroup
	instances := make([]*Config, goroutines)
	
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(index int) {
			defer wg.Done()
			instances[index] = GetInstance()
		}(i)
	}
	wg.Wait()
	
	// 验证所有实例都是同一个
	firstInstance := instances[0]
	for i := 1; i < goroutines; i++ {
		if instances[i] != firstInstance {
			t.Errorf("实例 %d 与第一个实例不同", i)
		}
	}
}

func TestEagerSingleton(t *testing.T) {
	instance1 := GetEagerInstance()
	instance2 := GetEagerInstance()
	
	if instance1 != instance2 {
		t.Error("饿汉式单例返回了不同的实例")
	}
}
