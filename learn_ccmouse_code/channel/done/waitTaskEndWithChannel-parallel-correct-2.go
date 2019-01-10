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

	// 大小写分开并行打印
	// 并发打印小写
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for _, worker := range workers {
		<-worker.done
	}

	// 并发打印大写
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	// 大写
	for _, worker := range workers {
		<-worker.done
	}
	// 等待所有任务结束，先向channel发送任务，然后一起接收完成信息
	// 这种情况下，程序可以正常打印小写字母，然后报错
	// 这是因为channel发送是阻塞的，发了之后就要收，导致程序报错deadlock
	// 解决方法参见另外程序
	// for _, worker := range workers {
	// 	<-worker.done
	// 	<-worker.done
	// }
}

func main() {
	chanAsPara()
}
