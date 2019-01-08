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
		done <- true
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

	// 依次串行打印，上1个任务完成了才能开始下1个任务
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		// 等待单个任务结束，任务依次进行
		<-workers[i].done
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		// 等待单个任务结束，任务依次进行
		<-workers[i].done
	}
}

func main() {
	chanAsPara()
}

/*
Worker 0 received a
Worker 1 received b
Worker 2 received c
Worker 3 received d
Worker 4 received e
Worker 5 received f
Worker 6 received g
Worker 7 received h
Worker 8 received i
Worker 9 received j
Worker 0 received A
Worker 1 received B
Worker 2 received C
Worker 3 received D
Worker 4 received E
Worker 5 received F
Worker 6 received G
Worker 7 received H
Worker 8 received I
Worker 9 received J
*/
