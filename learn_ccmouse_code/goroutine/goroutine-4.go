package main

import (
	"fmt"
	"time"
)

func main(){
	// 程序生成1000个goroutine，验证系统开启了多个线程
	// 线程数 = cpu核数 
	for i := 0; i < 1000; i++ {
		go func(ii int) {
			for {
				fmt.Printf("Hello from goroutine %d\n", ii)
			}
		}(i) // 传给匿名函数的参数
	}
	
	time.Sleep(time.Minute)
}