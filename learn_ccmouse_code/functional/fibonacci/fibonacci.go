package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 版本1
// 闭包
// 斐波那契数列生成器
/*
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	fib := fibonacci() // fib是生成器
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
}
*/

// 版本2
// 为函数实现io.Reader接口

// 闭包（返回函数进行了重命名）
func fibonacci() intGen {
	a, b := 0, 1
	return func() int { // 这里不能替换成别名
		a, b = b, a+b
		return a
	}
}

type intGen func() int // 函数类型（类型重命名，实现接口）

// 为函数类型添加方法
// 为函数类型intGen实现io.Reader接口（好好体会这里）
// TODO：incorrect if p is to small!
// 解决方法：构建struct存储intGen和strings.NewReader(s)
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	// 实现退出条件
	if next > 10000 {
		return 0, io.EOF
	}
	// 直接实现Read方法较麻烦（需要返回读入的字节数）
	// 这里借助别的方法，将next转换成字符串
	s := fmt.Sprintf("%d\n", next) // 转换成字符串
	// 将s字符串写入p
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	// 将f当文件读（实现了io.Reader接口）
	printFileContents(f)
}
