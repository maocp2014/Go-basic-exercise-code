package main

import "fmt"

func send(ch chan int) {
	i := 0
	for {
		i++
		ch <- i // channel关闭后继续写导致panic: send on closed channel
	}
}

func recv(ch chan int) {
	value := <-ch
	fmt.Println(value)
	value = <-ch
	fmt.Println(value)
	close(ch) // 接收方关闭了channel
}

func main() {
	var ch = make(chan int, 4)
	go recv(ch)
	send(ch)
}
