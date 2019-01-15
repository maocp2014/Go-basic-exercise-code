package main

import (
	"fmt"
)

type Rect struct {
	Width int
	Height int
}

// 接口变量的赋值
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

// 从上面的输出结果中可以推断出结构体的内存发生了复制，这个复制可能是因为赋值（a = r）也可能是因为类型转换（rx = a.(Rect)），
// 也可能是两者都进行了内存复制。那能不能判断出究竟在接口变量赋值时有没有发生内存复制呢？不好意思，就目前来说我们学到的知识点还办不到。
// 可以使用 unsafe 包来洞悉其中的更多细节。不过答案是是两者都会发生数据内存的复制 —— 浅拷贝。 