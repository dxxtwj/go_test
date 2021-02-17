package main

import (
	"fmt"
	"sync"
	"time"
)
type atomicInt struct {
	value int
	lock sync.Mutex
}
func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	// 用匿名函数生成一块区域，defer会在匿名函数执行完执行
	func() {
		a.lock.Lock()// 锁住
		defer a.lock.Unlock()// 解锁
		a.value++
	}()
}
func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}
func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
