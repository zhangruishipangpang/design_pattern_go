// Package design_patterns go 策略模式
package design_patterns

import "fmt"

// Strategy 定义策略接口
type Strategy interface {
	Do()
}

// Context 策略执行器
type Context struct {
	Strategy Strategy
}

func (c Context) Invoke() {
	c.Strategy.Do()
}

// ========== 具体实现 ==========

type Strategy1 struct {
}

func (s Strategy1) Do() {
	fmt.Println(" Strategy1 Do")
}

type Strategy2 struct {
}

func (s Strategy2) Do() {
	fmt.Println(" Strategy2 Do")
}
