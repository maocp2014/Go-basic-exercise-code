package main

import (
	"fmt"
	"time"
)


/*
这种形式比较繁琐，需要为每一种消费来源都单独启动一个汇聚协程。
Go 语言为这种使用场景带来了「多路复用」语法糖，也就是下面要讲的 select 语句，
它可以同时管理多个通道读写，如果所有通道都不能读写，它就整体阻塞，
只要有一个通道可以读写，它就会继续。下面我们使用 select 语句来简化上面的逻辑
*/
// 每隔一会生产一个数
func send(ch chan int, gap time.Duration) {
	i := 0
	for {
		i++
		ch <- i
		time.Sleep(gap)
	}
}

// 将多个原通道内容拷贝到单一的目标通道
func collect(source chan int, target chan int) {
	for v := range source {
		target <- v
	}
}

// 从目标通道消费数据
func recv(ch chan int) {
	for v := range ch {
		fmt.Printf("receive %d\n", v)
	}
}

func main() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)
	var ch3 = make(chan int)
	go send(ch1, time.Second)
	go send(ch2, 2*time.Second)
	go collect(ch1, ch3)
	go collect(ch2, ch3)
	recv(ch3)
}
