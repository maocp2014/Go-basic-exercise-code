package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("run in main goroutine")

	i := 1
	// for循环
	for i < 100000 {
		go func() {
			// 死循环
			for {
				time.Sleep(time.Minute)
			}
		}()

		if i%10000 == 0 {
			fmt.Printf("%d goroutine started\n", i)
		}
		i++
	}
}
