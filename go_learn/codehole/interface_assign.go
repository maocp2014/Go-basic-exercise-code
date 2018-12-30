package main

import (
	"fmt"
)

type Rect struct {
	Width int
	Height int
}

func main() {
	var a interface {}
	var r Rect = Rect {
		Width : 50,
		Height : 50,
	}

	// a = r
	a = &r // 接口变量a指向结构体指针

	// var rx = a.(Rect) // 类型转换语法，将接口变量a转换成结构体变量

	var rx = a.(*Rect) // 转换成结构体指针

	r.Height = 100
	r.Width = 100
	// fmt.Println(rx)  // 这里的值还是（50,50）
	fmt.Println(*rx) // 这里的值变为（100,100）
}