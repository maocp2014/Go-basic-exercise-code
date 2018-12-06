package main

import (
	"fmt"
	"io"
	"strings"
)

/*
Read 用数据填充指定的字节 slice，并且返回填充的字节数和错误信息。 在遇到数据流结尾时，返回 io.EOF 错误。
例子代码创建了一个 strings.Reader。 并且以每次 8 字节的速度读取它的输出。
*/
func main() {
	r := strings.NewReader("Hello, Reader!")
    fmt.Println(r)
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}