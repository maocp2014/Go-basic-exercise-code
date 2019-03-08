package main

import (
	"fmt"
	"time"
)

func send(ch chan int, gap time.Duration) {
	i := 0
	for {
		i++
		ch <- i
		time.Sleep(gap)
	}
}

// 多路复用 select 语句的读通道形式， 阻塞形式
func recv(ch1 chan int, ch2 chan int) {
	for {
		// select读通道形式，阻塞形式
		select {
		case v := <-ch1:
			fmt.Printf("recv %d from ch1\n", v)
		case v := <-ch2:
			fmt.Printf("recv %d from ch2\n", v)
		}
	}
}

func main() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	go send(ch1, time.Second)
	go send(ch2, 2*time.Second)
	recv(ch1, ch2)
}
