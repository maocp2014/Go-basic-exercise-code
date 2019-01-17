package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 包内变量定义方法（go中没有全局变量的说法），即作用域为包内

// 方式1
// var aa = 3
// var ss = "kkk"
// var bb = true

// 方式2（推荐方式）
var (
	aa = 33
	ss = "sss"
	bb = false
)

// 变量默认零值
func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)           // 这种方式空字符串打印不出来
	fmt.Printf("%d %q\n", a, s) // 打印空字符串的格式化字符
}

// 变量初始化
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// 变量类型自动推导
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// 短变量操作符（推荐的函数内变量定义）
func variableShorter() {
	a, b, c, s := 3, 4, true, "def" // 变量定义 + 赋值
	b = 5                           // 变量赋值
	fmt.Println(a, b, c, s)
}

// 欧拉公式
func euler() {
	// 保留小数点后3位
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1) //1i表示虚数i
}

// 强制类型转换（go中不支持类型隐式转换，只能强制转换）
func triangle() {
	var a, b int = 3, 4
	var c int
	// 强制类型转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 包内常量： 包内作用域
// 不用大写
const aaa = "ddd"  // 
// 函数内常量：函数内作用域
func consts() {
	// 方式1
	// const filename = "test.txt"
	// const a, b = 3, 4

	// 方式2
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	fmt.Println(filename, a, b)

	var c int
	// const数值（常量）可以作为各种类型使用，内部自动转换，不需要显式强制转换
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println("ccc:", c)
}

// 枚举类型
func enums() {
	// 方式1
	// const (
	// 	cpp = 0
	// 	java = 1
	// 	python = 2
	// 	golang = 3  
	// )
	// fmt.Println(cpp, java, python, golang)
	// 方式2
	const (
		// iota表示从0开始自增
		cpp = iota  // 0
		_  // 占位符，该处值为1
		python // 2
		golang // 3
		javascript // 4
	)
	fmt.Println(cpp, python, golang, javascript)
	
	// iota复杂运算
	// 字节单位
	const (
		b = 1 << (10 * iota)  // 1 << 10 * 0 表示1左移位，即1
		kb  // 1 << (10 * 1)  // 左移10位，即2的10次方，下面依次类推
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	fmt.Println(aa, ss, bb)
    // 欧拉公式
	euler()
    // 强制类型转换
	triangle()
    // 常量 
	consts()
    // 枚举类型
	enums()
}