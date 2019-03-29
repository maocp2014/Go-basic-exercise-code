package main

import (
	"fmt"
	"time"
)

func main() {
	theMine := [5]string{"ore1", "ore2", "ore3"}
	oreChan := make(chan string)

	// 寻矿者
	go func(mine [5]string) {
		for _, item := range mine {
			oreChan <- item //send
		}
	}(theMine)

	// 矿工
	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChan //接收
			fmt.Println("Miner: Received " + foundOre + " from finder")
		}
	}()

	// 主程序实际上是在自己的 Go 例程中运行的！更重要的是要知道，一旦主函数返回，它将关闭其它所有正在运行的例程
	<-time.After(time.Second * 5) // 计时器 —— 它创建了一个通道，并在 5 秒后发送了一个值。
}
