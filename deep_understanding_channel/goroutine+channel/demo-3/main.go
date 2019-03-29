package main

import (
	"fmt"
	"time"
)

func main() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	oreChannel := make(chan string)
	minedOreChan := make(chan string)
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
		for i := 0; i < 3; i++ {
			foundOre := <-oreChannel //从 oreChannel 上读取
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- "minedOre" //向 minedOreChan 发送
		}
	}()

	// Smelter
	go func() {
		for i := 0; i < 3; i++ {
			minedOre := <-minedOreChan //从 minedOreChan 读取
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
	}()

	<-time.After(time.Second * 5) // 计时器 —— 它创建了一个通道，并在 5 秒后发送了一个值。
}
