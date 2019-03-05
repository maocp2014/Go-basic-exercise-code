package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.NumGoroutine())
	for i := 0; i < 10; i++ {
		go func() {
			for {
				time.Sleep(time.Second)
			}
		}()
	}
	fmt.Println(runtime.NumGoroutine())
}
