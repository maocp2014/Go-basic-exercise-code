package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 读通道收数据
func recv(ch chan int) {
	for {
		value := <-ch
		fmt.Printf("recv %d\n", value)
		// time.Sleep(time.Second)
	}
}

func send(ch chan int) {
	for {
		value := rand.Intn(100)
		ch <- value
		fmt.Printf("send %d\n", value)
	}
}

func main() {
	// var ch chan int = make(chan int, 1)
	ch := make(chan int, 1)
	// 开协程收数据
	go recv(ch)
	// 主协程发数据
	go send(ch)

	// time.Sleep(time.Minute)
	for{
		time.Sleep(time.Second)
	}
}
