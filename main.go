package main

import "fmt"
import "go_project/design_patterns/observer"

func main() {

	//var h1 patterns.Handler = patterns.Invoke1{}
	//var h2 patterns.Handler = patterns.Invoke2{}
	//
	//chain := patterns.HandlerChain{Handlers: []*patterns.Handler{&h1, &h2}}
	//
	//chain.DoHandler("1")

	//since()

	observerTTT()
}

func observerTTT() {

	instance := observer.Instance()

	instance.Attach(&O1{})
	instance.Attach(&O2{})
	instance.Attach(&O2{})
	instance.Attach(&O2{})
	instance.Attach(&O2{})
	instance.Attach(&O2{})
	instance.Attach(&O2{})
	instance.Attach(&O2{})
	instance.Attach(&O2{})
	instance.Attach(&O2{})
	instance.Attach(&O2{})

	instance.Notify("通知来了")
}

type O1 struct {
}

func (o1 *O1) Listening(msg interface{}) {
	fmt.Println(" 执行O1 observer ", msg)
}

func (o1 *O1) Support(msg interface{}) bool {
	return true
}

type O2 struct {
}

func (o2 *O2) Listening(msg interface{}) {
	fmt.Println(" 执行O2 observer ", msg)
}

func (o2 *O2) Support(msg interface{}) bool {
	return true
}

func since() {

	i := make([]s1, 5)

	for in := range i {
		i[in] = s1{in}
	}

	fmt.Println(i)

	i = append(i[:2], i[2+1:]...) // 删除中间1个元素

	fmt.Println(i)
}

type s1 struct {
	id int
}
