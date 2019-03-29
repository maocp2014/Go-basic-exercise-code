package main

import (
	"fmt"
)

func main() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	oreChannel := make(chan string)
	minedOreChan := make(chan string)
	// goroutine完成时，往doneChan中发送数据表示完成
	doneChan := make(chan string)

	// Finder
	go func(mine [5]string) {
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item //在 oreChannel 上发送东西
			}
		}
	}(theMine)

	// Ore Breaker
	go func() {
		// for i := 0; i < 3; i++ {
		// 	foundOre := <-oreChannel //从 oreChannel 上读取
		// 	fmt.Println("From Finder: ", foundOre)
		// 	minedOreChan <- "minedOre" //向 minedOreChan 发送
		// }
		for foundOre := range oreChannel {
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- "minedOre" //向 minedOreChan 发送
		}
	}()

	// Smelter
	go func() {
		// 方式1
		for i := 0; i < 3; i++ {
			minedOre := <-minedOreChan //从 minedOreChan 读取
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}

		// 这里不能用for...range读取
		// for minedOre := range minedOreChan {
		// 	fmt.Println("From Miner: ", minedOre)
		// 	fmt.Println("From Smelter: Ore is smelted")
		// }

		// 如果只有发送方的 go-routine，没有其他的 go-routine。那么会发生死锁，go 程序会检测出死锁并崩溃。
		// 注意：所有以上讲的也都适用于接收方 go-routines。
		close(minedOreChan)
		// 完成时，发送数据到doneChan表示完成
		doneChan <- "I’m all done!"
	}()

	// <-time.After(time.Second * 5) // 计时器 —— 它创建了一个通道，并在 5 秒后发送了一个值。

	// 现在有更好的方法来处理阻塞主函数，直到所有其他的 Go 例程完成。
	// 通常的做法是创建一个主函数在等待读取时阻塞的 done 通道。一旦你完成你的工作，写入这个通道，程序将结束。
	<-doneChan // 阻塞直到 Go 例程发出工作完成的信号
}
