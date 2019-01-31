package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 方式1
func eval1(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		// return a / b
		q, _ := div2(a, b)
		return q
	default:
		panic("unsupported operation: " + op)
	}
}

// 方式2：推荐
func eval2(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// return a / b
		q, _ := div2(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: " + op)
	}
}

// 函数返回多个值（这里举例是2个返回值）
// 方式1：
func div1(a, b int) (int, int) {
	return a / b, a % b
}

// 方式2：命名返回值（非常简单的函数推荐，函数复杂时不好区分）
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return // 注意这里直接返回
}

// 函数式编程实现上述功能: 函数的参数也可以是函数，函数是一等公民
func apply(op func(int, int) int, a, b int) int {
	// 利用反射获取函数名
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

// 重写pow函数
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(eval1(3, 4, "*"))
	fmt.Println(div1(13, 3))
	q2, r2 := div2(13, 3)
	fmt.Println(q2, r2)
	fmt.Println(eval1(17, 4, "/"))

	fmt.Println(eval2(3, 2, "x"))
	// 推荐方式
	if result, err := eval2(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// 定义函数，传进去
	fmt.Println(apply(pow, 3, 4))
	// 直接定义匿名函数
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 4, 5))

	// 可变参数列表
	fmt.Println(sum(1, 2, 3, 4, 5))
}
