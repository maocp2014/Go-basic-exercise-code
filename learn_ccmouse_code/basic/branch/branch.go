package main

import (
	"fmt"
	"io/ioutil"
)

// switch方式1：switch后没有表达式
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score! %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	}
	return g
}

// switch方式2: switch后有表达式
func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+" :
		result = a + b
	case "-" :
		result = a - b
	case "*" :
		result = a * b
	case "/" :
		result = a / b
	default :
	    panic(fmt.Sprintf("unsupported operator: %s", op))
	}
	return result
}

func main() {
	const filename = "src\\go_pratice_code\\learn_ccmouse_code\\basic\\branch\\abc.txt"
	// 方式1:
	/*
		contents, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s\n", contents)
		}
	*/

	// 方式2(推荐)

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	// 注意：contents的作用域限定在if语句中，在if外引用会报错
	// fmt.Println(contents) // contents未定义

	// switch后无参数
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(70),
		grade(85),
		grade(95),
		// grade(101),
		// grade(-3),
	)

    // switch后有参数
	res := eval(32, 13, "+")
	fmt.Println(res)
}
