package main

import (
	"fmt"
	"time"
)

/*
在使用子协程时一定要特别注意保护好每个子协程，确保它们正常安全的运行。
因为子协程的异常退出会将异常传播到主协程，直接会导致主协程也跟着挂掉，然后整个程序就崩溃了。
利用recover保护协程
*/
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("main goroutine error")
		}
	}()

	fmt.Println("run in main goroutine")
	go func() {
		fmt.Println("run in child goroutine")
		go func() {
			fmt.Println("run in grand child goroutine")
			go func() {
				defer func() {
					if err := recover(); err != nil {
						fmt.Println("grand grand child goroutine error")
					}
				}()
				fmt.Println("run in grand grand child goroutine")
				panic("exception exit") // recover保护goroutine
			}()
		}()
	}()

	time.Sleep(time.Second)

	fmt.Println("main goroutine will quit") // 执行这行代码
	panic("panic main goroutine")           // panic后面的逻辑不会执行
}
