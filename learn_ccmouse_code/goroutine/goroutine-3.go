package main

import (
	"fmt"
	"time"
	"runtime"
)
// go run -race "src\go_pratice_code\learn_ccmouse_code\goroutine\goroutine-3.go" 检查访问数据冲突
// 协程内部直接引用外面的变量会引发运行时错误(runtime error)
/*
func main(){
	var a [10]int
	for i := 0; i < 10; i++ {
		// 这里是闭包的概念，不传参给里面的匿名函数，匿名函数就会引用for循环中的i，当for循环退出时
		// i = 10，匿名函数中a数组引用了a[10]，导致数组越界，引发runtime error
		go func() { // race condition 
			// 死循环
			for {
				a[i]++
				// 手动交出控制权，让别的协程也有机会运行
			    runtime.Gosched()
			}
		}() // 传给匿名函数的参数
	}
	
	time.Sleep(time.Second)
	fmt.Println(a)
}
*/

func main(){
	var a [10]int
	for i := 0; i < 10; i++ {
		// 这里是闭包的概念，不传参给里面的匿名函数，匿名函数就会引用for循环中的i，当for循环退出时
		// i = 10，匿名函数中a数组引用了a[10]，导致数组越界，引发runtime error
		// 为了解决这个问题，需要将每个goroutine的ii固定下来，因此通过传递参数每个goroutine复制一个i
		go func(ii int) {
			// 死循环
			for {
				a[ii]++
				// 手动交出控制权，让别的协程也有机会运行
			    runtime.Gosched()
			}
		}(i) // 传给匿名函数的参数
	}
	
	time.Sleep(time.Second)
	// 上述race condition通过传参解决，但这里还是存在race condition，这是因为a[ii]++在写数据的同时
	// 而fmt.Println同时在读数据导致race condition，这个需要通过channel来解决
	fmt.Println(a)
}