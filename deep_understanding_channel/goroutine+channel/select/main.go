package main

import (
	"fmt"
	"time"
)

func main() {
	myChan := make(chan string)

	go func() {
		// myChan <- "Message!"
		select {
		case myChan <- "message":
			fmt.Println("sent the message")
		default:
			fmt.Println("no message sent")
		}
	}()

	select {
	case msg := <-myChan:
		fmt.Println(msg)
	default:
		fmt.Println("No Msg")
	}

	// 计时器
	<-time.After(time.Second * 5)

	select {
	case msg := <-myChan:
		fmt.Println(msg)
	default:
		fmt.Println("No Msg")
	}
}
