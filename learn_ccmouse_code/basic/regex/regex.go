package main

import (
	"fmt"
	"regexp"
)

const text = "My email is ccmouse@gmail.com"

func main() {
	// 方式1
	// re, err := regexp.Compile("ccmouse@gmail.com")
	// if err != nil {
	// 	panic(err)
	// }

	// 方式2
	re := regexp.MustCompile("ccmouse@gmail.com")
	match := re.FindString(text)
	fmt.Println(match)

}
