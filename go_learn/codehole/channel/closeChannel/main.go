package main

import "fmt"

/*
对于已关闭的通道读取要分两种情况，如果通道是有缓冲的，会将缓存的值读出来，如果没有缓存值了再继续读就是零值；
对于非缓冲通道关闭，读出的是零值。
怎样判断通道关闭了，可以对通道进行读取，返回两个值，通过第二个值是否为false来判断通道是否关闭
*/
func main() {
	// 缓冲型通道
	var ch1 chan int = make(chan int, 4)
	ch1 <- 1
	ch1 <- 2
	close(ch1) // 关闭通道

	value := <-ch1
	fmt.Println(value) // 1
	value = <-ch1
	fmt.Println(value) // 2
	value = <-ch1
	fmt.Println(value) // 0
	value = <-ch1
	fmt.Println(value) // 0

	// 非缓冲型通道
	ch2 := make(chan int)
	close(ch2) // 关闭通道
	value1 := <-ch2
	fmt.Println("xxx:", value1) // 0
	value1 = <-ch2
	fmt.Println("xxx:", value1) // 0
}
