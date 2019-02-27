package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccmouse@gmail.com"
email1 is abc@def.org
email2 is kkk@qq.com
email3 is kkkk@hnu.com.cn
`

func main() {
	// 方式1
	// re, err := regexp.Compile("ccmouse@gmail.com")
	// if err != nil {
	// 	panic(err)
	// }

	// 方式2
	// re := regexp.MustCompile(".+@.+\\..+") // 方式1
	// re := regexp.MustCompile(`.+@.+\..+`)  // 方式2  ``中的字符串不会发生转义

	// re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`) // 方式3

	// 子匹配，相当于分组
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9]+)`) // 方式4

	// match := re.FindString(text) // 找单个
	// match := re.FindAllString(text, -1) // 找多个  输出：[ccmouse@gmail.com abc@def.org kkk@qq.com]
	//fmt.Println(match)

	// 子匹配，提取相应字段
	matchs := re.FindAllStringSubmatch(text, -1)
	for _, match := range matchs {
		fmt.Println(match)
	}
	/*
	output:
		[ccmouse@gmail.com ccmouse gmail .com]
		[abc@def.org abc def .org]
		[kkk@qq.com kkk qq .com]
		[kkkk@hnu.com.cn kkkk hnu.com .cn]
	*/
}
