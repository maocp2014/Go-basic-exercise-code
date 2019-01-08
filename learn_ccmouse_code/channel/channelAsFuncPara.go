package main

import (
	"fmt"
	"time"
)

// channel和函数一样，在go中都是一等公民
// channel作为函数参数
/*
func worker(c chan int) {
	for {
		n := <-c
		fmt.Println(n)
	}
}
*/

func worker(id int, c chan int) {
	for {
		// 方式1
		// n := <- c
		// fmt.Printf("Worker %d received %d\n", id, n)

		// 方式2
		// fmt.Printf("Worker %d received %d\n", id, <- c)

		fmt.Printf("Worker %d received %c\n", id, <- c)
	}
}

func chanAsPara() {
	// 1、共用1个channel
	/*
	c := make(chan int)
	goroutine接收数据
	go worker(c)
	for i := 0; i < 10; i++ {
		go worker(i, c)
	}

	// 往c中发送数据
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	*/

	// 2、单独channel
	var channels [10]chan int // 1等公民，10个channel int数组
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}
	
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	} 
	
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	} 
}

func main() {
	// 可能存在只打印某些数据，这是因为来不及打印数据，main函数就退出了，可以加上time,Sleep
	chanAsPara()
	time.Sleep(time.Millisecond)
}
