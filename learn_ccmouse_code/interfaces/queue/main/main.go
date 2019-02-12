package main

import (
	"fmt"
	"go_pratice_code/learn_ccmouse_code/interfaces/queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q)
	q.Pop()
	q.Pop()
	fmt.Println(q)
	q.Pop()
	flag := q.IsEmpty()
	fmt.Println(flag)

	q.Push(123)
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
}
