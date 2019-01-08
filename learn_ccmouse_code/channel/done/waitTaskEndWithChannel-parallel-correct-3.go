package main

import (
	"fmt"
	"sync"
)

// 接收数据并打印
// 增加一个channel通知外面worker是否打印完成
func doWorker(id int, in chan int, wg *sync.WaitGroup) {
	for n := range in {
		fmt.Printf("Worker %d received %c\n", id, n)
		// 通知外面打印完成
		wg.Done()
	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

// 返回channel，作用是接收数据
func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(id, w.in, w.wg)
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
