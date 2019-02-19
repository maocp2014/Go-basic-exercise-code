package fib

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
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int { // 这里不能替换成别名
		a, b = b, a+b
		return a
	}
}
