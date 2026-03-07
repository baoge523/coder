package main

import (
	"fmt"
	"sync"
)

// Config 配置管理器（懒汉式 - 使用 sync.Once）
type Config struct {
	DatabaseURL string
	APIKey      string
}

var (
	instance *Config
	once     sync.Once
)

// GetInstance 获取单例实例（线程安全）
func GetInstance() *Config {
	once.Do(func() {
		fmt.Println("创建单例实例...")
		instance = &Config{
			DatabaseURL: "localhost:5432",
			APIKey:      "secret-key-123",
		}
	})
	return instance
}

// EagerConfig 饿汉式单例
type EagerConfig struct {
	AppName string
	Version string
}

var eagerInstance = &EagerConfig{
	AppName: "MyApp",
	Version: "1.0.0",
}

// GetEagerInstance 获取饿汉式单例实例
func GetEagerInstance() *EagerConfig {
	return eagerInstance
}

func main() {
	// 测试懒汉式单例
	fmt.Println("=== 懒汉式单例测试 ===")
	config1 := GetInstance()
	config2 := GetInstance()
	
	fmt.Printf("config1 地址: %p\n", config1)
	fmt.Printf("config2 地址: %p\n", config2)
	fmt.Printf("是否为同一实例: %v\n", config1 == config2)
	fmt.Printf("配置内容: %+v\n", config1)
	
	// 测试饿汉式单例
	fmt.Println("\n=== 饿汉式单例测试 ===")
	eager1 := GetEagerInstance()
	eager2 := GetEagerInstance()
	
	fmt.Printf("eager1 地址: %p\n", eager1)
	fmt.Printf("eager2 地址: %p\n", eager2)
	fmt.Printf("是否为同一实例: %v\n", eager1 == eager2)
	fmt.Printf("配置内容: %+v\n", eager1)
}
