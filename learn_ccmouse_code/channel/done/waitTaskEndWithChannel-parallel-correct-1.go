package main

import (
	"fmt"
)

// 接收数据并打印
// 增加一个channel通知外面worker是否打印完成
func doWorker(id int, in chan int, done chan bool) {
	for n := range in {
		fmt.Printf("Worker %d received %c\n", id, n)
		// 通知外面打印完成
		// done <- true

		// 解决deadlock问题，true由另一个goroutine发送
		go func() {
			done <- true
		}()
	}
}

type worker struct {
	in   chan int
	done chan bool
}

// 返回channel，作用是接收数据
func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func chanAsPara() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// 收两次done，为解决deadlock，done交由另一个goroutine发送，避免阻塞导致deadlock
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

func main() {
	chanAsPara()
}
