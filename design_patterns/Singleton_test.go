package design_patterns

import (
	"fmt"
	"testing"
	"time"
)

func TestObjInstance(t *testing.T) {

	// 这里返回的单例对象使用&取地址各不相同，但是都指向同一个最终内存地址
	for i := 0; i < 3; i++ {
		go func() {
			lazyInstance := ObjLazyInstance()

			fmt.Println("懒汉式Object : ", lazyInstance.Id, lazyInstance.Name)
		}()
	}

	for i := 0; i < 3; i++ {
		go func() {
			starveInstance := ObjStarveInstance()

			fmt.Println("饿汉式Object : ", starveInstance.Id, starveInstance.Name)
		}()
	}

	time.Sleep(5 * time.Second)
}
