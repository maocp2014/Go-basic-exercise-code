package main

import "fmt"

// 更新切片
func updateSlice(arr []int) {
	arr[0] = 200
}

// 对数组进行切片
func arraySlice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// 切片都是半开半闭，即前闭后开（包括前面不包括后面）
	fmt.Println("arr[2:6] = ", arr[2:6]) // 下标2~5
	fmt.Println("arr[:6] = ", arr[:6])   // 下标0~5
	fmt.Println("arr[2:] = ", arr[2:])   // 下标 2~7
	fmt.Println("arr[:] = ", arr[:])     // 下标0~7

	s1 := arr[2:]
	fmt.Println("s1 = ", s1)
	s2 := arr[:]
	fmt.Println("s2 = ", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)  // [200 3 4 5 6 7]
	fmt.Println(arr) // [0 1 200 3 4 5 6 7]

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)  // [200 1 200 3 4 5 6 7]
	fmt.Println(arr) // [200 1 200 3 4 5 6 7]

	fmt.Println("Reslice") // 在slice基础上slice
	s2 = arr[:]
	fmt.Println("s2 = ", s2)
	s2 = s2[:5]
	fmt.Println("s2 = ", s2)
	s2 = s2[2:]
	fmt.Println("s2 = ", s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	s3 := arr[2:6]
	s4 := s3[3:5]
	fmt.Printf("s3 = %v, len(s3) = %d, cap(s3) = %d\n", s3, len(s3), cap(s3)) // s3 =  [2 3 4 5]
	fmt.Println("s4 = ", s4)                                                  // s4 =  [5 6]
	// fmt.Println("s3[4] = ", s3[4]) // panic: runtime error: index out of range

	s5 := append(s4, 10)
	s6 := append(s5, 11)
	s7 := append(s6, 12)
	fmt.Println("s5, s6, s7 = ", s5, s6, s7) // s5, s6, s7 =  [5 6 10] [5 6 10 11] [5 6 10 11 12]
	fmt.Println("arr = ", arr)               // arr =  [0 1 2 3 4 5 6 10]
}

func printSlice(s []int) {
	fmt.Printf("v = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}

// 创建slice
func createSlice() {
	// 方法1
	var s []int    // Zero value for slice is nil
	fmt.Println(s) // []
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	// 方法2
	s1 := []int{2, 4, 6, 8} // 底层数组的view
	printSlice(s1)
	// 方法3
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)
}

// 切片复制
func copySlice() {
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	copy(s2, s1)
	printSlice(s2)
}

// 删除切片中的元素
func deleteElementsFromSlice() {
	s := []int{2, 4, 6, 8, 0, 0, 0, 0, 0, 0}
	s = append(s[:3], s[4:]...) // "..."表示可变参数
	printSlice(s)

	// 删除首元素
	front := s[0]
	s = s[1:]
	// 删除尾元素
	tail := s[len(s)-1]
	s = s[:len(s)-1]
	fmt.Println(front, tail)
	printSlice(s)
}

func main() {
	fmt.Println("slice of array")
	arraySlice()

	fmt.Println("Creating slice")
	createSlice()

	fmt.Println("Copying slice")
	copySlice()

	fmt.Println("Deleting elements from slice")
	deleteElementsFromSlice()
}
