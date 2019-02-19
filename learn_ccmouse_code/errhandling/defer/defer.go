package main

import (
	"bufio"
	"errors"
	"fmt"
	"go_pratice_code/learn_ccmouse_code/functional/fibonacci/fib"
	"os"
)

// defer 先进后出的栈，在函数返回前执行
// 即使函数中有return语句、panic语句，defer语句也会执行
// 下面输出结果: "3\n2\n1\n"
func tryDefer_1() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	// return   // defer语句遇到return也会执行
	panic("error occurred!") // defer语句遇到panic也会执行
	fmt.Println(4)
}

func tryDefer_2() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many!")
		}
	}
}

func writeFile(filename string) {
	// file, err := os.Create(filename) // 打开写文件

	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	
	// 自定义error
	// err = errors.New("this is a custom error!")

	if err != nil {
		// fmt.Println("file already exists")
		// fmt.Println("Error:", err.Error())
		// fmt.Println("Error:", err)
		// 简单的错误处理
		// 类型断言err.(*os.PathError)，判断是不是PathError
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close() // 关闭文件

	// 直接写文件比较慢，这里用到bufio包装，实现缓冲写
	// 这里只是写到了buffer中，最后需要刷新写到文件
	writer := bufio.NewWriter(file)
	defer writer.Flush() // 写缓冲刷新到文件

	f := fib.Fibonacci() // 斐波那契数列生成器
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// tryDefer_1()
	// tryDefer_2()
	writeFile("D:\\VSCodeWorkSpace\\goWorkspace\\src\\go_pratice_code\\learn_ccmouse_code\\errhandling\\defer\\fib.txt")
}
