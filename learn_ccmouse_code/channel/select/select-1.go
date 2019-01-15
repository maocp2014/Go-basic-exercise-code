package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
// select实现非阻塞获取channel的值
func main() {
	var c1, c2 chan int // c1 and c2 = nil
    // 运行程序会输出default
	select {
	case n := <-c1:
		fmt.Println("Received from c1:", n)
	case n := <-c2:
		fmt.Println("Received from c2:", n)
	default:
		fmt.Println("no value received!")
	}
}
*/

// 生成chan int
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()

	w := createWorker(0)

	for {
		select {
		case n := <-c1:
			// fmt.Println("Received from c1:", n)
			w <- n // 阻塞
		case n := <-c2:
			// fmt.Println("Received from c2:", n)
			w <- n  // 阻塞
		}
	}
}
