package main

import "fmt"

// 斐波那契数列 F[n]=F[n-1]+F[n-2](n>=3,F[1]=1,F[2]=1)
// fibonacci函数会返回一个返回int的函数。
func fibonacci() func() int {
	a, b := 1, 1
	return func() int {
		temp := a
		a, b = b, a + b
		return temp // go中局部变量也可以返回，不会报错，c++中不允许
	}

}

func main() {
	f := fibonacci() // 类似于python中的生成器
	// fmt.Println(f())
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}