// Package observer go 观察者
package observer

/*
	定义一个观察者对象
*/

type Observer interface {
	Listening(interface{})
	Support(interface{}) bool
}
