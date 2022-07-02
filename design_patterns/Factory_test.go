// Package design_patterns go 工厂方法模式
package design_patterns

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {

	var product Product

	var factoryCreator Factory = &FactoryCreator{}

	// 创建手机
	product = factoryCreator.GetObject(FMobile)
	fmt.Printf("产品类型：%s,  价格：%f", product.Name(), product.Price())

	fmt.Println()

	// 创建车辆
	product = factoryCreator.GetObject(FCar)
	fmt.Printf("产品类型：%s, 价格：%f", product.Name(), product.Price())

	fmt.Println()
}
