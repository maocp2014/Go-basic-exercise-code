package main

import (
	"fmt"
	"time"
)

// 普通版：不加go关键字的情况
/*
func main(){
	for i := 0; i < 10; i++ {
		// 匿名函数
		func(i int) {
			// 死循环，无退出条件
			// 一直打印: "Hello from goroutine 0"
			for {
				fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i) // 传给匿名函数的参数
	}
}
*/

// goroutine版：加go关键字情况
// 主程序开1000个协程打印内容，主程序直接往下执行
func main(){
	for i := 0; i < 1000; i++ {
		// go + 匿名函数 构成go的协程goroutine
		go func(i int) {
			// 死循环
			for {
				fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i) // 传给匿名函数的参数
	}
	// 这里main程序需要等待goroutine执行完，否则什么不打印(main直接退出，goroutine来不及打印结果)
	time.Sleep(time.Second)
}