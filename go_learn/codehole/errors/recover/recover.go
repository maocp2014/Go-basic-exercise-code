package main

import (
	"errors"
	"fmt"
)

// var negErr = fmt.Errorf("non positive number")
var negErr = errors.New("non positive number")

// panic 抛出的对象未必是错误对象，而 recover() 返回的对象正是 panic 抛出来的对象，所以它也不一定是错误对象
// 我们经常还需要对 recover() 返回的结果进行判断，以挑选出我们愿意处理的异常对象类型，对于那些不愿意处理的，可以选择再次抛出来，让上层来处理。
// func panic(v interface{})
// func recover() interface{}

func main() {
	// defer + 匿名函数
	// defer 确保 recover() 逻辑在程序异常的时候也可以得到调用
	// recover的作用是防止程序崩溃退出
	defer func() {
		if err := recover(); err != nil {
			if err == negErr{
				fmt.Println("error catched", err) // error catched non positive number
			} else {
				panic(err)  // 重新抛出err
			}
		}
	}()  // 匿名函数的调用

	fmt.Println(fact(5))
	fmt.Println(fact(10))
	fmt.Println(fact(-2))
	fmt.Println(fact(15)) // 不会执行
}

func fact(a int) int {
	if a <= 0 {
		panic(negErr) // 抛出negErr被recover捕获，程序不会崩溃，但后续逻辑也不会执行
	}

	r := 1
	for i := 1; i <= a; i++ {
		r *= i
	}
	return r
}
