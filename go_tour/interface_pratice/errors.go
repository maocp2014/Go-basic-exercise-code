/*
自定义error错误信息
*/
package main

import (
	"fmt"
	"time"
)

// 错误信息结构体
type MyError struct {
	when time.Time
	what string
}

// 内置error接口
//type error interface {
//    Error() string
//}
// 实现error接口中Error()方法，与fmt.Stringer类似，fmt包在输出时也会试图匹配error接口。
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}

// 返回值为error接口类型
func run() error {
	return &MyError{
		time.Now(),
        "it don't work!",
	}
}

func main() {
	if err := run(); err != nil {
		// 因为run()方法返回的一个error，故err参数的值就为Error()方法的返回值，即一个string字符串
		fmt.Println(err)
	}
}