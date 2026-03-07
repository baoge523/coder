package main

import "fmt"

// Product 产品接口
type Product interface {
	Use() string
}

// ConcreteProductA 具体产品A
type ConcreteProductA struct {
	name string
}

func (p *ConcreteProductA) Use() string {
	return fmt.Sprintf("使用产品A: %s", p.name)
}

// ConcreteProductB 具体产品B
type ConcreteProductB struct {
	name string
}

func (p *ConcreteProductB) Use() string {
	return fmt.Sprintf("使用产品B: %s", p.name)
}

// ProductType 产品类型
type ProductType string

const (
	TypeA ProductType = "A"
	TypeB ProductType = "B"
)

// Factory 工厂
type Factory struct{}

// CreateProduct 创建产品（简单工厂）
func (f *Factory) CreateProduct(productType ProductType) Product {
	switch productType {
	case TypeA:
		return &ConcreteProductA{name: "产品A"}
	case TypeB:
		return &ConcreteProductB{name: "产品B"}
	default:
		return nil
	}
}

// Creator 工厂方法接口
type Creator interface {
	CreateProduct() Product
}

// ConcreteCreatorA 具体工厂A
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ConcreteProductA{name: "工厂方法产品A"}
}

// ConcreteCreatorB 具体工厂B
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ConcreteProductB{name: "工厂方法产品B"}
}

func main() {
	fmt.Println("=== 简单工厂模式 ===")
	factory := &Factory{}
	
	productA := factory.CreateProduct(TypeA)
	fmt.Println(productA.Use())
	
	productB := factory.CreateProduct(TypeB)
	fmt.Println(productB.Use())
	
	fmt.Println("\n=== 工厂方法模式 ===")
	var creator Creator
	
	creator = &ConcreteCreatorA{}
	product1 := creator.CreateProduct()
	fmt.Println(product1.Use())
	
	creator = &ConcreteCreatorB{}
	product2 := creator.CreateProduct()
	fmt.Println(product2.Use())
}
