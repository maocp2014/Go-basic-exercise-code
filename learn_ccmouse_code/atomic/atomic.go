package main

import (
	"fmt"
	"sync"
	"time"
)

// 不加锁，数据访问冲突
/*
type atomicInt int

func (a *atomicInt) increment() {
	*a++
}

func (a *atomicInt) get() int {
	return int(*a)
}
func main() {
	var a atomicInt
	a.increment()

	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
*/

// 加锁
type atomicInt struct {
	value int
	lock  sync.Mutex
}

// 一般加锁操作
// func (a *atomicInt) increment() {
// 	a.lock.Lock()
// 	defer a.lock.Unlock()
// 	a.value++
// }

// 实现代码块原子操作（加锁）
func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	// lock控制在函数体中，实现了代码块加锁释放锁
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
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
