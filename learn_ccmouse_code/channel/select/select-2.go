package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
		time.Sleep(time.Second)  // 演示产生和消耗数据的速度不一致
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	// 版本1: 实现非阻塞，但会一直执行case w <- n语句，
	// 因为前2条语句需要等待，而第3条语句直接执行 
	/*
	var c1, c2 = generator(), generator()

	w := createWorker(0)

	for {
		n := 0
		select {
		case n = <-c1:
		case n = <-c2:
		case w <- n:
		}
	}
	*/

	// 版本2：select中同时实现数据的收和发
	// 存在问题: 生成数据的速度和消耗数据的速度不同，导致n数据丢失，输出结果如下：
	// Worker 0 received 0
    // Worker 0 received 7
	// Worker 0 received 12
	
	// 解决方法见版本3：对收到的数据进行排队
	/*
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	n := 0
	hasValue := false
	
	for {
		var activeWorker chan<- int  // activeWorker = nil
        if hasValue {
			activeWorker = worker
		}

		select {
		case n = <- c1:
			hasValue = true
		case n = <- c2:
			hasValue = true
		// 当activeWorker=nil时，这里会一直阻塞，利用了chan为nil的特性
		case activeWorker <- n:
			hasValue = false
		}
	}
	*/

	// 版本3：解决收发数据速度不一致问题
	// 方法：对收到的数据进行排队
	// 问题：程序不能退出
	/*
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	var values []int
	
	for {
		var activeWorker chan<- int  // activeWorker = nil
		var activeValue int

        if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <- c1:
			values = append(values, n)
		case n := <- c2:
			values = append(values, n)
		// 当activeWorker=nil时，这里会一直阻塞，利用了chan为nil的特性
		case activeWorker <- activeValue:
			values = values[1:]
		}
	}
	*/

	// 版本4：解决程序退出问题以及完善程序
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	var values []int
	// 返回值tm为chan
	// 表示10s后会向tm chan中发送一个时间（具体时间不确定）
	// 程序运行开始一直到目前的时间
	tm := time.After(10 * time.Second)
	// 每隔1s会送一个值
	tick := time.Tick(time.Second)
	
	for {
		var activeWorker chan<- int  // activeWorker = nil
		var activeValue int

        if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <- c1:
			values = append(values, n)
		case n := <- c2:
			values = append(values, n)
		// 当activeWorker=nil时，这里会一直阻塞，利用了chan为nil的特性
		case activeWorker <- activeValue:
			values = values[1:]
		// 每次select化的时间
		// 两次生成数据的时间差超过800ms,则打印超时
		case <- time.After(800 * time.Millisecond):
			fmt.Println("time out")
		// 每隔1s查看队列的长度
		case <- tick:
			fmt.Println("queue len =", len(values))
		// 10s后退出，定时的作用
		case <- tm:
			fmt.Println("bye")
			return
		}
	}
}
