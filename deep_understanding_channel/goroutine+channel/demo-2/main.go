package main

import (
	"fmt"
	"time"
)

func main() {
	bufferedChan := make(chan string, 3)

	go func() {
		bufferedChan <- "first"
		fmt.Println("Sent 1st")
		bufferedChan <- "second"
		fmt.Println("Sent 2nd")
		bufferedChan <- "third"
		fmt.Println("Sent 3rd")
	}()

	<-time.After(time.Second * 1)

	go func() {
		firstRead := <-bufferedChan
		fmt.Println("Receiving..")
		fmt.Println(firstRead)
		secondRead := <-bufferedChan
		fmt.Println(secondRead)
		thirdRead := <-bufferedChan
		fmt.Println(thirdRead)
	}()

	<-time.After(time.Second * 5) // 计时器 —— 它创建了一个通道，并在 5 秒后发送了一个值。
}
