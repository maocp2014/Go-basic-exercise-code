package main

import (
	"fmt"
	"time"
)

// channel作为函数返回值
// 该函数本身只是创建goroutine并返回，不做具体事情
// 具体事情交由另1个goroutine去做
/*
func createWorker(id int) chan int {
	c := make(chan int)
    // 收数据交给另1个goroutine,否则死循环
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <- c)
		}
	}()
	return c
}
*/

// 方便复用
func chanWorker(id int, c chan int) {
	// 方式1: 判断channel中是否有数据
	/*
	for {
		// 判断channel是否还有数据，ok = true表示还有数据，否则没有数据
		n, ok := <- c
		fmt.Println("ok:", ok)
		if !ok {
			break
		}
		// fmt.Printf("Worker %d received %d\n", id, <-c)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
	*/
	// 方式2: 判断channel中是否有数据
	// 这种方式既可以用在关闭channel的情况，也可以用在不关闭channel的情况（goroutine会一直收数据，直到某一方退出）
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

// 对上述createWork进一步优化
// 告诉使用者channel的作用，这里channel是发数据的(向该channel发数据)，自己本身是收数据的
func createWorker(id int) chan<- int {
	c := make(chan int)
	// 收数据交给另1个goroutine,否则死循环
	go chanWorker(id, c)
	return c
}

func chanAsParameter() {
	// var channels [10]chan int // 1等公民，10个channel int数组
	var channels [10]chan<- int // 这里相应修改

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
}

func bufferedChannel() {
	c := make(chan int, 3) //缓冲区，提升性能
	// 开协程收数据
	go chanWorker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	// 开协程收数据
	go chanWorker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c) // 关闭channel后，goroutine也会收到数据，只不过数据是0
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen") // channel是一等公民
	// 可能存在只打印某些数据，这是因为来不及打印数据，main函数就退出了，可以加上time,Sleep
	chanAsParameter()
	fmt.Println("buffered channel") 
	// bufferedChannel()
	fmt.Println("channel close and range")  // 关闭channel和range方式收数据
	channelClose()
}
