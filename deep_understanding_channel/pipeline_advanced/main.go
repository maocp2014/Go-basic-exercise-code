package main

import (
	"fmt"
	"sync"
)

// <-chan int 表示单向channel
// 表示数据从管道出来，对于调用者就是得到管道的数据，当然就是输入
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// merge 函数 将多个 channel 转换为一个 channel，它为每一个 inbound channel 启动一个 goroutine，
// 用于将数据拷贝到 outbound channel。
func merge(in ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// 为每一个输入channel in 创建一个 goroutine output
	// output 将数据从 c 拷贝到 out，直到 c 关闭，然后 调用 wg.Done
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(in))

	for _, i := range in {
		go output(i)
	}

	// 启动一个 goroutine，用于所有 output goroutine结束时，关闭 out
	// 该goroutine 必须在 wg.Add 之后启动
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := gen(2, 3)

	// 启动两个 sq 实例，即两个goroutines处理 channel "in" 的数据
	c1 := sq(in)
	c2 := sq(in)

	// merge 函数将 channel c1 和 c2 合并到一起，这段代码会消费 merge 的结果
	out := merge(c1, c2)

	for n := range out {
		fmt.Println(n)
	}
}
