package main

import "fmt"

// 值传递，不改变原变量的值
func swap1(a, b int) {
	a, b = b, a
}

// 指针传递，改变原变量的值
func swap2(a, b *int) {
	*a, *b = *b, *a
}

// 更好的交换两数的写法
func swap3(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	swap1(a, b)
	fmt.Println(a, b)

	swap2(&a, &b)
	fmt.Println(a, b)

	a, b = swap3(a, b)
	fmt.Println(a, b)
}
