// Package observer go 观察者
package observer

import "sync"

/*
	定义一个观察者的通知结构体【也就是被观察对象】
*/

type subject struct {
	lock      sync.Locker
	observers []Observer
}

var SS subject = subject{
	lock:      &sync.RWMutex{},
	observers: make([]Observer, 1<<2),
}

func Instance() subject {
	return SS
}

func init() {
	//SS := subject{}
	//SS.lock = &sync.RWMutex{}
	//SS.observers = make([]Observer, 1<<2)
}

func (s *subject) Attach(observer Observer) bool {

	// 添加一个观察者对象时需要添加锁，避免并发添加问题
	s.lock.Lock()
	defer s.lock.Unlock()

	s.observers = append(s.observers, observer)

	//index := len(s.observers) + 1
	//s.observers[index] = observer

	return true
}

func (s *subject) Detach(observer Observer) bool {

	// 取消一个观察者对象时需要添加锁，避免并发添加问题
	s.lock.Lock()
	defer s.lock.Unlock()

	var deleteIndex = -1
	for index, ob := range s.observers {
		if ob == observer {
			s.observers[index] = nil
			deleteIndex = index
			break
		}
	}
	if deleteIndex == -1 {
		return false
	}
	// 掩埋删除的节点
	s.observers = append(s.observers[:deleteIndex], s.observers[deleteIndex+1:]...)
	return true
}

func (s *subject) Notify(msg interface{}) {
	currentObserver := s.observers
	// 开始执行通知
	for _, ob := range currentObserver {
		if ob != nil && ob.Support(msg) {
			ob.Listening(msg)
		}
	}
}
