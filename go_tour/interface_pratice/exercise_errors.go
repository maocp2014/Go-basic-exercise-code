/*
题目：
从先前的练习中复制 Sqrt 函数，并修改使其返回 error 值。

由于不支持负数，当 Sqrt 接收到一个负数时，应当返回一个非 nil 的错误值。

创建一个新类型  type ErrNegativeSqrt float64 为其实现方法（实际是实现接口的方法，重写Error()）
func (e ErrNegativeSqrt) Error() string

使其成为一个error， 该方法就可以让 ErrNegativeSqrt(-2).Error() 返回 “cannot Sqrt negative number: -2”

提示：
定义float64的类型ErrNegativeSqrt ，并重写Error()方法
*/

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} 
	return math.Sqrt(x), nil
}

// 别名自定义类型
type ErrNegativeSqrt float64

// 为自定义类型实现接口方法
func (e ErrNegativeSqrt) Error() string {
	// return fmt.Sprintf("cannot Sqrt negative number: %v", e) // 报错
	// 注意： 在 Error 方法内调用 fmt.Sprint(e) 将会让程序陷入死循环。可以通过先转换 e 来避免这个问题：`fmt.Sprint(float64(e))`。请思考这是为什么呢？
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e)) // 正确方法
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
