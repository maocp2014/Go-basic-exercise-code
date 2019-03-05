package main

import "fmt"
import "time"

func main() {
	fmt.Println("run in main goroutine")
	// 当 n 值大于 3 时，主协程将没有机会得到运行，而如果 n 值为 3、2、1，主协程依然可以每秒输出一次
    n := 3
    for i:=0; i<n; i++ {
        go func() {
            fmt.Println("dead loop goroutine start")
            for {}  // 死循环
        }()
    }
    for {
        time.Sleep(time.Second)
        fmt.Println("main goroutine running")
    }
}
