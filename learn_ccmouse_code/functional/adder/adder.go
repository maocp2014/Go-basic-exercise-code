package main

import "fmt"

// go中的闭包：返回值是函数
func adder1() func(int) int {
	// 自由变量
	sum := 0
	// 返回闭包（函数 + sum变量的引用）
	// v是局部变量
	return func(v int) int {
		fmt.Println("before add: sum = ", sum)
		sum += v
		fmt.Println("after add: sum = ", sum)
		return sum
	}
}

// 正统的闭包实现累加和
// 不用自由变量
type iAdder func(int) (int, iAdder)  // 函数定义别名

// 闭包
func adder2(base int) iAdder{
	return func(v int) (int, iAdder){
		return base + v, adder2(base + v)
	}
}

func main() {
	// 1、adder1函数求和
	// a := adder1()
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	// }

	// 2、adder2函数求和
	a := adder2(0)
	for i := 0; i < 10; i++{
		var s int
		s, a = a(i)
        fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}
}
