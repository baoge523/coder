package main

import (
	"fmt"
	"time"
)

// Subject 主题接口
type Subject interface {
	Request() string
}

// RealSubject 真实主题
type RealSubject struct {
	name string
}

func NewRealSubject(name string) *RealSubject {
	return &RealSubject{name: name}
}

func (r *RealSubject) Request() string {
	// 模拟耗时操作
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("RealSubject: 处理请求 - %s", r.name)
}

// Proxy 代理
type Proxy struct {
	realSubject *RealSubject
	cache       map[string]string
}

func NewProxy() *Proxy {
	return &Proxy{
		cache: make(map[string]string),
	}
}

func (p *Proxy) Request() string {
	// 延迟初始化
	if p.realSubject == nil {
		fmt.Println("Proxy: 创建 RealSubject 实例")
		p.realSubject = NewRealSubject("数据")
	}
	
	// 访问控制
	fmt.Println("Proxy: 检查访问权限")
	
	// 缓存
	cacheKey := "request"
	if result, exists := p.cache[cacheKey]; exists {
		fmt.Println("Proxy: 从缓存返回结果")
		return result
	}
	
	// 调用真实对象
	fmt.Println("Proxy: 调用 RealSubject")
	result := p.realSubject.Request()
	
	// 缓存结果
	p.cache[cacheKey] = result
	
	// 日志记录
	fmt.Println("Proxy: 记录访问日志")
	
	return result
}

// 实际应用示例：图片代理

// Image 图片接口
type Image interface {
	Display()
}

// RealImage 真实图片
type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	img := &RealImage{filename: filename}
	img.loadFromDisk()
	return img
}

func (r *RealImage) loadFromDisk() {
	fmt.Printf("加载图片: %s\n", r.filename)
	time.Sleep(200 * time.Millisecond) // 模拟加载时间
}

func (r *RealImage) Display() {
	fmt.Printf("显示图片: %s\n", r.filename)
}

// ProxyImage 图片代理
type ProxyImage struct {
	filename  string
	realImage *RealImage
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{filename: filename}
}

func (p *ProxyImage) Display() {
	if p.realImage == nil {
		p.realImage = NewRealImage(p.filename)
	}
	p.realImage.Display()
}

func main() {
	fmt.Println("=== 代理模式 ===\n")
	
	// 基础代理示例
	fmt.Println("基础代理:")
	proxy := NewProxy()
	
	// 第一次请求
	fmt.Println("\n第一次请求:")
	result1 := proxy.Request()
	fmt.Printf("结果: %s\n", result1)
	
	// 第二次请求（使用缓存）
	fmt.Println("\n第二次请求:")
	result2 := proxy.Request()
	fmt.Printf("结果: %s\n", result2)
	
	// 图片代理示例
	fmt.Println("\n\n图片代理:")
	image1 := NewProxyImage("photo1.jpg")
	image2 := NewProxyImage("photo2.jpg")
	
	// 图片只在第一次显示时加载
	fmt.Println("\n第一次显示 photo1.jpg:")
	image1.Display()
	
	fmt.Println("\n第二次显示 photo1.jpg:")
	image1.Display()
	
	fmt.Println("\n第一次显示 photo2.jpg:")
	image2.Display()
}
