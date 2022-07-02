// Package design_patterns go 工厂方法模式
package design_patterns

/* ========== 产品族定义 ==========*/

// Product 定义Product抽象接口
type Product interface {
	Name() string
	Price() float32
}

// Mobile 定义手机产品
type Mobile struct {
}

func (m *Mobile) Name() string {
	return "mobile"
}

func (m *Mobile) Price() float32 {
	return 3000.00
}

// Car 定义车产品
type Car struct {
}

func (c *Car) Name() string {
	return "car"
}

func (c *Car) Price() float32 {
	return 10 * 10000
}

/* ========== 工厂类定义 ==========*/

// Factory 工厂接口
type Factory interface {
	GetObject(name string) Product
}

// FactoryCreator 工厂实际创建者
type FactoryCreator struct {
}

func (creator *FactoryCreator) GetObject(name string) Product {
	var product Product
	switch name {
	case "mobile":
		product = &Mobile{}
	case "car":
		product = &Car{}
	default:
		panic("创建对象需要为:{'mobile', 'car'}, 但是当前是:" + name)

	}
	return product
}

const (
	FMobile = "mobile"
	FCar    = "car"
)
