package main

import "fmt"

func main() {
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	close(ch)

	// for range 语法我们已经见了很多次了，它是多功能的，除了可以遍历数组、切片、字典，还可以遍历通道，取代箭头操作符。
	// 当通道空了，循环会暂停阻塞，当通道关闭时，阻塞停止，循环也跟着结束了。当循环结束时，我们就知道通道已经关闭了。
	for value := range ch {
		fmt.Println(value)
	}
}
