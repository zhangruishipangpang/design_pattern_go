// Package design_patterns go 责任链模式
package design_patterns

import (
	"fmt"
	"reflect"
)

// Handler 定义处理器接口
type Handler interface {
	DoHandler(interface{})
	Support(interface{}) bool
}

// HandlerChain 具体的执行责任链的处理器
type HandlerChain struct {
	Handlers []*Handler
}

func (chain *HandlerChain) init() {
	chain.Handlers = make([]*Handler, 10)
}

func (chain *HandlerChain) DoHandler(req interface{}) {
	if len(chain.Handlers) == 0 {
		fmt.Println("本次职责链未配置职责链条，执行结束")
		return
	}
	for _, handler := range chain.Handlers {
		if (*handler).Support(req) {
			(*handler).DoHandler(req)
		}
	}
	fmt.Println("职责链执行完成，职责链数量：", len(chain.Handlers))
}

func (chain *HandlerChain) Support(interface{}) bool {
	// 永远返回true，其实这个返回无所谓，因为它是执行调用链的具体struct
	return true
}

// Invoke1 具体的执行者
type Invoke1 struct {
}

func (Invoke1) DoHandler(req interface{}) {
	fmt.Println("Invoke1 执行... ", req)
}

func (Invoke1) Support(req interface{}) bool {
	if req != nil && reflect.TypeOf(req).Name() == "int" {
		return true
	}
	return false
}

// Invoke2 具体的执行者
type Invoke2 struct {
}

func (Invoke2) DoHandler(req interface{}) {
	fmt.Println("Invoke2 执行... ", req)
}

func (Invoke2) Support(req interface{}) bool {
	if req != nil && reflect.TypeOf(req).Name() == "string" {
		return true
	}
	return false
}
