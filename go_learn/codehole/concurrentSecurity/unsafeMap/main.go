package main

import "fmt"

func write(d map[string]int) {
	d["fruit"] = 2
}

func read(d map[string]int) {
	fmt.Println(d["fruit"])
}

// 线程不安全的map
// Go 语言内置了数据结构「竞态检查」工具来帮我们检查程序中是否存在线程不安全的代码。
// 当我们在运行代码时，打开 -race 开关，程序就会在内置的通用数据结构中进行埋点检查
// 竞态检查工具是基于运行时代码检查，而不是通过代码静态分析来完成的。
// 这意味着那些没有机会运行到的代码逻辑中如果存在安全隐患，它是检查不出来的。
func main() {
	d := map[string]int{}
	go read(d)
	write(d)
}
