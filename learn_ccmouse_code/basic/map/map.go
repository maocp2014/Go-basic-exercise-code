package main

import "fmt"

// 创建map
func createMap() {
	// 方式1
	m1 := map[string]string{
		"name":   "xxx",
		"age":    "18",
		"gender": "male",
	}
	fmt.Println(m1)

	// 方式2
	m2 := make(map[string]int)
	fmt.Println(m2) // m2 == empty map

	// 方式3
	var m3 map[string]int
	fmt.Println(m3) // m3 == nil
}

func traverseMap() {
	// map的key是无序的
	m := map[string]string{
		"name":   "xxx",
		"age":    "18",
		"gender": "male",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

// map中获取值
func getValues(){
	m := map[string]string{
		"name":   "eric",
		"age":    "18",
		"gender": "male",
	}

	userName, ok := m["name"]
	fmt.Println("userName, ok: ", userName, ok)
	// 当获取map中不存在的key时，会获取到zero value(map中为空窜)
	causeName, ok := m["xxx"]
	fmt.Println("causeName, ok: ", causeName, ok)

	// 正确获取map中key对应value的方法
	if gender, ok := m["gender"]; ok {
		fmt.Println(gender)
	} else {
		fmt.Println("key does not exist")
	}
}

// 删除map中的元素
func deleteElement(){
	m := map[string]string{
		"name":   "eric",
		"age":    "18",
		"gender": "male",
	}

	name, ok := m["name"]
	fmt.Println("name, ok: ", name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println("name, ok: ", name, ok)
}

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
		if i - start +1 > maxLength {
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
		if i - start +1 > maxLength {
			maxLength = i - start + 1
		}
        // 更新字符最后出现的位置
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	createMap()
	traverseMap()
	getValues()
	deleteElement()
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
