package main

import (
	"errors"
	"fmt"
)

// var negErr = fmt.Errorf("non positive number")
var negErr = errors.New("non positive number")

func main() {
	fmt.Println(fact(5))
	fmt.Println(fact(10))
	fmt.Println(fact(-2))
	fmt.Println(fact(15)) // 不会执行
}

func fact(a int) int {
	if a <= 0 {
		panic(negErr) // 抛出negErr, 程序崩溃
	}

	r := 1
	for i := 1; i <= a; i++ {
		r *= i
	}
	return r
}
