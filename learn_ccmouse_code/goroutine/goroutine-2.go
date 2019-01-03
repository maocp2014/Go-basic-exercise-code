package main

import (
	"fmt"
	"time"
	"runtime"
)

// 举例说明goroutine是“抢占式多任务”
// fmt.Printf是IO操作，会进行协程的自动切换
// 算术运算a[i]++不会进行协程的自动切换，这个程序卡死在第一个协程里面，不会有任何输出
/*
func main(){
	var a [10]int
	for i := 0; i < 10; i++ {
		// go + 匿名函数 构成go的协程goroutine
		go func(i int) {
			// 死循环
			for {
				// fmt.Printf("Hello from goroutine %d\n", i)
				a[i]++
			}
		}(i) // 传给匿名函数的参数
	}
	// main函数也是一个goroutine
	// 由于第1个协程执行自加加运行，不会主动交出控制权，所以程序不会执行到这里
	time.Sleep(time.Second)
	fmt.Println(a)
}
*/

func main(){
	var a [10]int
	for i := 0; i < 10; i++ {
		// go + 匿名函数 构成go的协程goroutine
		go func(i int) {
			// 死循环
			for {
				// fmt.Printf("Hello from goroutine %d\n", i)
				a[i]++
				// 手动交出控制权，让别的协程也有机会运行
			    runtime.Gosched()
			}
		}(i) // 传给匿名函数的参数
	}
	
	time.Sleep(time.Second)
	fmt.Println(a)
}