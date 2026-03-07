package main

import "fmt"

// Computer 产品
type Computer struct {
	CPU     string
	Memory  string
	Disk    string
	GPU     string
	Monitor string
}

func (c *Computer) Show() {
	fmt.Printf("电脑配置:\n")
	fmt.Printf("  CPU: %s\n", c.CPU)
	fmt.Printf("  内存: %s\n", c.Memory)
	fmt.Printf("  硬盘: %s\n", c.Disk)
	fmt.Printf("  显卡: %s\n", c.GPU)
	fmt.Printf("  显示器: %s\n", c.Monitor)
}

// ComputerBuilder 建造者接口
type ComputerBuilder interface {
	SetCPU(cpu string) ComputerBuilder
	SetMemory(memory string) ComputerBuilder
	SetDisk(disk string) ComputerBuilder
	SetGPU(gpu string) ComputerBuilder
	SetMonitor(monitor string) ComputerBuilder
	Build() *Computer
}

// ConcreteComputerBuilder 具体建造者
type ConcreteComputerBuilder struct {
	computer *Computer
}

func NewComputerBuilder() *ConcreteComputerBuilder {
	return &ConcreteComputerBuilder{
		computer: &Computer{},
	}
}

func (b *ConcreteComputerBuilder) SetCPU(cpu string) ComputerBuilder {
	b.computer.CPU = cpu
	return b
}

func (b *ConcreteComputerBuilder) SetMemory(memory string) ComputerBuilder {
	b.computer.Memory = memory
	return b
}

func (b *ConcreteComputerBuilder) SetDisk(disk string) ComputerBuilder {
	b.computer.Disk = disk
	return b
}

func (b *ConcreteComputerBuilder) SetGPU(gpu string) ComputerBuilder {
	b.computer.GPU = gpu
	return b
}

func (b *ConcreteComputerBuilder) SetMonitor(monitor string) ComputerBuilder {
	b.computer.Monitor = monitor
	return b
}

func (b *ConcreteComputerBuilder) Build() *Computer {
	return b.computer
}

// Director 指挥者
type Director struct {
	builder ComputerBuilder
}

func NewDirector(builder ComputerBuilder) *Director {
	return &Director{builder: builder}
}

// BuildGamingComputer 构建游戏电脑
func (d *Director) BuildGamingComputer() *Computer {
	return d.builder.
		SetCPU("Intel i9-13900K").
		SetMemory("32GB DDR5").
		SetDisk("2TB NVMe SSD").
		SetGPU("NVIDIA RTX 4090").
		SetMonitor("4K 144Hz").
		Build()
}

// BuildOfficeComputer 构建办公电脑
func (d *Director) BuildOfficeComputer() *Computer {
	return d.builder.
		SetCPU("Intel i5-12400").
		SetMemory("16GB DDR4").
		SetDisk("512GB SSD").
		SetGPU("集成显卡").
		SetMonitor("1080P 60Hz").
		Build()
}

func main() {
	fmt.Println("=== 建造者模式 ===\n")
	
	// 使用指挥者构建
	builder := NewComputerBuilder()
	director := NewDirector(builder)
	
	fmt.Println("游戏电脑:")
	gamingPC := director.BuildGamingComputer()
	gamingPC.Show()
	
	fmt.Println("\n办公电脑:")
	builder = NewComputerBuilder() // 重新创建建造者
	director = NewDirector(builder)
	officePC := director.BuildOfficeComputer()
	officePC.Show()
	
	// 直接使用建造者（链式调用）
	fmt.Println("\n自定义电脑:")
	customPC := NewComputerBuilder().
		SetCPU("AMD Ryzen 7 5800X").
		SetMemory("32GB DDR4").
		SetDisk("1TB SSD").
		SetGPU("AMD RX 6800 XT").
		SetMonitor("2K 165Hz").
		Build()
	customPC.Show()
}
