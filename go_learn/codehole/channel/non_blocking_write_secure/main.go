package main

import (
	"fmt"
	"time"
)

func recv(ch chan int, gap time.Duration, name string) {
	for v := range ch {
		fmt.Printf("receive %s %d\n", name, v)
		time.Sleep(gap)
	}
}

func send(ch1, ch2 chan int) {
	i := 0
	for {
		i++
		// select写通道形式，非阻塞形式（如果去掉default形式，则是阻塞形式）
		select {
		case ch1 <- i:
			fmt.Printf("send ch1 %d\n", i)
		case ch2 <- i:
			fmt.Printf("send ch2 %d\n", i)
		// select 语句的 default 分支非常关键，它是决定通道读写操作阻塞与否的关键
		default:
			fmt.Println("execute default branch")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)

	go recv(ch1, time.Millisecond, "ch1")
	go recv(ch2, 2*time.Millisecond, "ch2")

	send(ch1, ch2)
}
