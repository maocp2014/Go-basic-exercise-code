package main

import (
	"fmt"
	"time"
)

/*
main 函数运行在主协程(main goroutine)里面，例子中我们在主协程里面启动了一个子协程，子协程又启动了一个孙子协程，孙子协程又启动了一个曾孙子协程。
这些协程之间似乎形成了父子、子孙、关系，但是实际上协程之间并不存在这么多的层级关系，在 Go 语言里只有一个主协程，其它都是它的子协程，
子协程之间是平行关系。

值得注意的是这里的 go 关键字语法和前面的 defer 关键字语法是一样的，它后面跟了一个匿名函数，然后还要带上一对()，表示对匿名函数的调用。
*/

func main() {
	fmt.Println("run in main goroutine")
	go func() {
		fmt.Println("run in child goroutine")
		go func() {
			fmt.Println("run in grand child goroutine")
			go func() {
				fmt.Println("run in grand grand child goroutine")
			}()
		}()
	}()
	// 主协程睡眠了 1s，等待子协程们执行完毕。如果将睡眠的这行代码去掉，将会看不到子协程运行的痕迹
	time.Sleep(time.Second)
	fmt.Println("main goroutine will quit")
}
