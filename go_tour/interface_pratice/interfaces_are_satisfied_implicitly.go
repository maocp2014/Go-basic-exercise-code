package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

// 接口组合
type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer
    // os.Stdout实现了Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, Writer!")
}