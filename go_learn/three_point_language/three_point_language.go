package main

import "fmt"

// '...'语法糖用于函数中，表示可以接受可变参数（任意个指定类型的参数）
func variableParameterTest(args ...string) {
	for k, v := range args {
		fmt.Println(k, v)
	}
}

// 值传递，但是是传递的切片头 —— 指向底层数组的指针ptr、长度len、 容量cap，其中包含了指向切片数据的指针，所以s1, s2的地址跟main函数中的地址一样
// 切片的结构是指向数据的指针，长度和容量。复制切片，或者在切片上创建新切片，切片中的指针都指向相同的数据内存区域。
func appendFunc(s0, s2 []int) []int {
	s0 = append(s0, s2...)
	return s0
}

func main() {
	var str = []string{"aaa", "bbb", "ccc", "ddd", "eee"}
	// '...'语法糖用于函数中传递参数，表示将slice打散进行传递（形式上），slice可以为任意长度
	variableParameterTest(str...)

	// '...'应用于append函数中
	var s0 []int = make([]int, 10, 20)
	fmt.Printf("s0: %v %p\n", s0, s0)

	var s1 = []int{1, 2, 3, 4, 5}
	// 切片复制
	copy(s0, s1)

	var s2 = []int{6, 7, 8, 9, 10}

	// 切片赋值也是切片头的拷贝
	s3 := appendFunc(s0, s2)
	fmt.Printf("s3: %v %p", s3, s3)
}
