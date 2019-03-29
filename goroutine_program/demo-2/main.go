package main

import "fmt"

var (
	R int = 100
	N int = 100
)

/*创建一个包含有「R」个 go-routines 的池。不多也不少，所有对「input」channel 的监听通过闭包属性来引用。
创建 go-routines，它通过在每次循环中检查 ok 参数来判断 channel 是否关闭，如果 channel 关闭则杀死自己。
返回 input channel 来允许调用者函数分配任务给池。
使用「task」参数来允许调用函数定义 go-routines 的主体。
*/
func Workers(task func(int)) chan int { // Point #4
	input := make(chan int)  // Point #1
	for i := 0; i < R; i++ { // Point #1
		go func() {
			for {
				v, ok := <-input // Point #2
				if ok {
					task(v) // Point #4
				} else {
					return // Point #2
				}
			}
		}()
	}
	return input // Point #3
}

func Task(i int) {
	fmt.Println("Box", i)
}

func main() {
	ack := make(chan bool, N)

	workers := Workers(func(a int) { // Point #2
		Task(a)
		ack <- true // Point #1
	})

	for i := 0; i < N; i++ {
		workers <- i
	}
	
	for i := 0; i < N; i++ { // Point #3
		<-ack
	}
}
