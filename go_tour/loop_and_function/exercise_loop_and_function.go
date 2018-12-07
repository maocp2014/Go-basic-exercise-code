/*
题目：
用牛顿法实现开方函数。
在这个例子中，牛顿法是通过选择一个初始点 z 然后重复这一过程求 Sqrt(x) 的近似值： 

为了做到这个，只需要重复计算 10 次，并且观察不同的值（1，2，3，……）是如何逐步逼近结果的。 然后，修改循环条件，
使得当值停止改变（或改变非常小）的时候退出循环。观察迭代次数是否变化。结果与 math.Sqrt 接近吗？

提示：
定义并初始化一个浮点值，向其提供一个浮点语法或使用转换：
z := float64(1)
z := 1.0
*/ 

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// 定义初始值
	z := float64(1)
	// 存储z的临时值
	temp := float64(0)

	for {
		// 牛顿公式
		z = z - (z*z - x)/(2*z)
		fmt.Println("z:", z)
		if math.Abs(z - temp) < 0.000000000000001 {
			// 当值改变非常小时，退出循环
			break
		} else {
			temp = z
		}
	}
	return z 
}

func main() {
	fmt.Println("牛顿法：",Sqrt(2))
	fmt.Println("math.Sqrt(2):", math.Sqrt(2))
}