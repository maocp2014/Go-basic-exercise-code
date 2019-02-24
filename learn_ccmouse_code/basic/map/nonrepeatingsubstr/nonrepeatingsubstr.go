package main

import "fmt"

// 寻找最长不含重复字符的子串
// 版本1（不支持中文）
func lengthOfNonRepeatingSubStr_1(s string) int {
	// 字母最后出现的位置map
	lastOccurred := make(map[byte]int)
	// 起始位置
	start := 0
	// 最大长度
	maxLength := 0
	// 将字符串转换成字节切片（1个字节表示1个字符）
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		// 更新长度
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		// 更新字符最后出现的位置
		lastOccurred[ch] = i
	}
	return maxLength
}

// 版本2（支持中英文）
func lengthOfNonRepeatingSubStr_2(s string) int {
	// 字母最后出现的位置map
	lastOccurred := make(map[rune]int)
	// 起始位置
	start := 0
	// 最大长度
	maxLength := 0
	// 将字符串转换成字节切片（1个字节表示1个字符）
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		// 更新长度
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		// 更新字符最后出现的位置
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	// 寻找最大不重复子串
	fmt.Println(lengthOfNonRepeatingSubStr_1("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr_2("bbbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr_2("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr_2(""))
	fmt.Println(lengthOfNonRepeatingSubStr_2("b"))
	fmt.Println(lengthOfNonRepeatingSubStr_2("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr_2("我爱中国长城"))
	fmt.Println(lengthOfNonRepeatingSubStr_2("一二三三二一"))
}
