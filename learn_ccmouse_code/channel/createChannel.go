package main

import (
	"fmt"
	"time"
)

// 创建channel
/*
func createChan() {
	// c1是1个channel，里面内容类型是int
	var c1 chan int // c1 = nil，channel为nil时一般没法使用（select中有特殊作用）
	fmt.Println(c1)
	c2 := make(chan int)  // 创建1个channel, c2 = 0xc00003e060(内存地址)
	fmt.Println(c2)
	c2 <- 1
	c2 <- 2
	n := <-c2
	fmt.Println(n)  // fatal error: all goroutines are asleep - deadlock!
	// channel是goroutine之间传递数据的桥梁，向channel中发送了数据，必须有另外
	// 的goroutine接收数据，否则会产生deadlock
}
*/

// goroutine利用channel接收数据
func chanDemo() {
	c := make(chan int)
	// goroutine接收数据
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	// 往channel中发送数据
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
}

func main() {
	// createChan()
	// 可能存在只打印某些数据，这是因为来不及打印数据，main函数就退出了，可以加上time,Sleep
	chanDemo()
	time.Sleep(time.Millisecond)
}
