package main

import (
	"fmt"
	"go_project/design_pattern_go/design_patterns"
	"time"
)

func main() {
	//singleton()
	//factory()
}

// 单例
func singleton() {

	// 这里返回的单例对象使用&取地址各不相同，但是都指向同一个最终内存地址
	for i := 0; i < 3; i++ {
		go func() {
			lazyInstance := design_patterns.ObjLazyInstance()

			fmt.Println("懒汉式Object : ", lazyInstance.Id, lazyInstance.Name)
		}()
	}

	for i := 0; i < 3; i++ {
		go func() {
			starveInstance := design_patterns.ObjStarveInstance()

			fmt.Println("饿汉式Object : ", starveInstance.Id, starveInstance.Name)
		}()
	}

	time.Sleep(5 * time.Second)
}

// 工厂方法
func factory() {
	var product design_patterns.Product

	var factoryCreator design_patterns.Factory = &design_patterns.FactoryCreator{}

	// 创建手机
	product = factoryCreator.GetObject(design_patterns.FMobile)
	fmt.Printf("产品类型：%s,  价格：%f", product.Name(), product.Price())

	fmt.Println()

	// 创建车辆
	product = factoryCreator.GetObject(design_patterns.FCar)
	fmt.Printf("产品类型：%s, 价格：%f", product.Name(), product.Price())

}
