package main

import (
	"fmt"
)

func tryRecover() {
	// 匿名函数，recover只在defer调用中使用
	defer func() {
		// recover()返回的是interface{}任意类型
		r := recover()
		// 类型断言
		if err, ok := r.(error); ok {
			fmt.Println("Error Occurred: ", err)
		} else { // 如果r不是error类型，重新panic
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}() // ()表示调用该匿名函数

	// 测试1
	// panic(errors.New("this is an error!"))

	// 测试2
	// recover发挥作用，程序不会挂掉
	// 输出：Error Occurred:  runtime error: integer divide by zero
	// b := 0
	// a := 5 / b
	// fmt.Println(a)

	// 测试3
	// 因为panic不是error，所以这里会重新panic
	// panic: 123 [recovered]
	// panic: I don't know what to do: 123
	// ...
	panic(123)
}

func main() {
	tryRecover()
}
