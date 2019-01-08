package main

import (
	"fmt"
	"sync"
)

// 接收数据并打印
// 增加一个channel通知外面worker是否打印完成
func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		// 通知外面打印完成
		w.done()
	}
}

type worker struct {
	in   chan int
	done func() // 函数式编程，抽象程序更好，done交由调用者自己定义，更灵活，更抽象
}

// 返回channel，作用是接收数据
func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		// 调用者自己定影done，更抽象，更灵活
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func chanAsPara() {
	var wg sync.WaitGroup
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	// 添加20个任务
	wg.Add(20)

	// 大小写分开并行打印
	// 并发打印小写
	for i, worker := range workers {
		worker.in <- 'a' + i
		// 每次添加一个任务也行，这里知道是20个任务，所以一次性添加了
		// wg.Add(1)
	}

	// 并发打印大写
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanAsPara()
}
