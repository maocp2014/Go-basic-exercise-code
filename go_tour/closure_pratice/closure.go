package main

import (
	"fmt"
	"math"
)

// 函数也是值
func func_is_value() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
    fmt.Println(hypot)  // hypot是指向函数的指针？
	fmt.Println(hypot(3, 4))
}

// 闭包, 返回值是函数，该函数被绑定在各自的sum变量上
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	func_is_value()

	pos, neg := adder(), adder()

	for i :=0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2 * i))
	}
}