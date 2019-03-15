package main

import "fmt"

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

func main() {
	// 设置流水线
	c := gen(2, 3)
	out := sq(c)
	// 消费输出结果
	fmt.Println(<-out)
	fmt.Println(<-out)

	// 由于sq函数的inbound channel和outbound channel类型一样，所以组合任意个sq函数。比如像下面这样使用：
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
	}
}
