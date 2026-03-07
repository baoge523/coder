package main

import "fmt"

// Beverage 饮料接口
type Beverage interface {
	GetDescription() string
	Cost() float64
}

// Espresso 浓缩咖啡
type Espresso struct{}

func (e *Espresso) GetDescription() string {
	return "浓缩咖啡"
}

func (e *Espresso) Cost() float64 {
	return 25.0
}

// HouseBlend 混合咖啡
type HouseBlend struct{}

func (h *HouseBlend) GetDescription() string {
	return "混合咖啡"
}

func (h *HouseBlend) Cost() float64 {
	return 20.0
}

// CondimentDecorator 调料装饰器基类
type CondimentDecorator struct {
	beverage Beverage
}

// Mocha 摩卡装饰器
type Mocha struct {
	CondimentDecorator
}

func NewMocha(beverage Beverage) *Mocha {
	return &Mocha{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
	}
}

func (m *Mocha) GetDescription() string {
	return m.beverage.GetDescription() + ", 摩卡"
}

func (m *Mocha) Cost() float64 {
	return m.beverage.Cost() + 5.0
}

// Milk 牛奶装饰器
type Milk struct {
	CondimentDecorator
}

func NewMilk(beverage Beverage) *Milk {
	return &Milk{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
	}
}

func (m *Milk) GetDescription() string {
	return m.beverage.GetDescription() + ", 牛奶"
}

func (m *Milk) Cost() float64 {
	return m.beverage.Cost() + 3.0
}

// Whip 奶泡装饰器
type Whip struct {
	CondimentDecorator
}

func NewWhip(beverage Beverage) *Whip {
	return &Whip{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
	}
}

func (w *Whip) GetDescription() string {
	return w.beverage.GetDescription() + ", 奶泡"
}

func (w *Whip) Cost() float64 {
	return w.beverage.Cost() + 4.0
}

func printOrder(beverage Beverage) {
	fmt.Printf("%s: ¥%.2f\n", beverage.GetDescription(), beverage.Cost())
}

func main() {
	fmt.Println("=== 装饰器模式 - 咖啡订单系统 ===\n")

	// 订单1: 浓缩咖啡
	beverage1 := &Espresso{}
	printOrder(beverage1)

	// 订单2: 混合咖啡 + 摩卡
	beverage2 := &HouseBlend{}
	beverage := NewMocha(beverage2)
	printOrder(beverage)

	// 订单3: 浓缩咖啡 + 双倍摩卡 + 奶泡
	beverage3 := &Espresso{}
	beverage31 := NewMocha(beverage3)
	beverage32 := NewMocha(beverage31)
	beverage33 := NewWhip(beverage32)
	printOrder(beverage33)

	// 订单4: 混合咖啡 + 牛奶 + 摩卡 + 奶泡
	beverage4 := &HouseBlend{}
	beverage41 := NewMilk(beverage4)
	beverage42 := NewMocha(beverage41)
	beverage43 := NewWhip(beverage42)
	printOrder(beverage43)
}
