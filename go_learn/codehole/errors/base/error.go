package main

import (
	"errors"
	"fmt"
)

/*
// 内置的错误接口，虽然接口名error是小写，但它是全局可见的
type error interface {
  Error() string
}
*/

// 定义错误内容的结构体
type SomeError struct {
	Reason string
}

// 实现内置error接口的Error方法
func (err SomeError) Error() string {
	return err.Reason
}

func main() {
	// sError := SomeError{
	// 	Reason : "some error happened!",
	// }

	// 接口赋值
	var sError error = SomeError{"some error happened!"}
	// 打印接口直接输出错误信息？？？？
	fmt.Println(sError)

	// 包外调用errors包，利用New生成错误信息，返回值err是接口
	var err = errors.New("some error occured!")
	fmt.Println(err)

	// 错误输出定制参数的用法
	var thing = "other error"
	var error = fmt.Errorf("%s happened!", thing)
	fmt.Println(error)
}

/*
// go中的错误处理通用包：errors
package errors

func New(text string) error {
    return &errorString{text}
}

// 注意这不是全局结构体，对外不可见，包外需通过New函数调用
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
*/
