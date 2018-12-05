package main

import (
	"fmt"
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	c := strings.Fields(s)  // 返回[]string
	
	for _, v := range c {
		fmt.Println(v)
		m[v] += 1  // m[v]初始值为0，即使键值v不存在
	}
	return m
}

func main() {
	// WordCount("welcome to go!")
	wc.Test(WordCount)
}