package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 我的祖国！"
	fmt.Println(len(s))           // 7个英语字符 + 5个中文 * 3 = 22， utf-8编码
	fmt.Printf("%s\n", []byte(s)) // 打印上述字符串
	fmt.Printf("%X\n", []byte(s)) // 打印上述字符串的字节
	// 打印每个字节
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b) // 48 65 6C 6C 6F 2C 20 E6 88 91 E7 9A 84 E7 A5 96 E5 9B BD EF BC 81
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch) // (0 48) (1 65) (2 6C) (3 6C) (4 6F) (5 2C) (6 20) (7 6211) (10 7684) (13 7956) (16 56FD) (19 FF01)
	}
	fmt.Println()

	// utf8包
	// string中含有多少个rune(类似于其它语言中的char)
	runeCount := utf8.RuneCountInString(s)
	fmt.Println("Rune count: ", runeCount) // Rune count:  12
	
	// 打印每个rune
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch) // H e l l o ,   我 的 祖 国 ！
	}
	fmt.Println()

	// rune切片
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch) // (0 H) (1 e) (2 l) (3 l) (4 o) (5 ,) (6  ) (7 我) (8 的) (9 祖) (10 国) (11 ！)
	}
	fmt.Println()
}
