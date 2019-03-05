package main

import "fmt"

func send(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	// 那如何确保呢？Go 语言并不存在一个内置函数可以判断出通道是否已经被关闭。即使存在这样一个函数，当你判断时通道没有关闭，
	// 并不意味着当你往通道里写数据时它就一定没有被关闭，并发环境下，它是可能被其它协程随时关闭的。
    // 确保通道写安全的最好方式是由负责写通道的协程自己来关闭通道，读通道的协程不要去关闭通道。
	close(ch)  // 写通道方自己关闭channel
}

func recv(ch chan int) {
	// for...range循环读取channel数据
	for value := range ch{
		fmt.Println(value)
	}
}

func main() {
	var ch = make(chan int, 1)
	go send(ch)
	recv(ch)
}
