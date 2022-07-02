// Package design_patterns go 单例模式
package design_patterns

import (
	"sync"
)

// object 单例对象
type object struct {
	Id   int32
	Name string
}

var (
	once      = sync.Once{}
	objLazy   *object
	objStarve *object = &object{1, "饿汉式"}
)

// ObjLazyInstance 获取懒汉式单例
func ObjLazyInstance() *object {

	// 可以查看Once.Do 内部使用CAS轻量级锁保证该方法只会执行一次, 这里在判断一次相当于加了两次校验锁
	if objLazy == nil {

		once.Do(func() {
			objLazy = &object{2, "懒汉式"}
		})
	}
	return objLazy
}

// ObjStarveInstance 获取饿汉式单例
func ObjStarveInstance() *object {
	return objStarve
}
