package main

import "fmt"

// 定义一维数组的3种方法及遍历数组
func define_linear_array() {
	var arr1 [5]int // 初始化当前类型的默认零值, 5个int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr1, arr2, arr3) // [0 0 0 0 0] [1 2 3] [4 5 6 7]

	// 遍历数组方法1
	for i := 0; i < len(arr3); i++ {
		fmt.Println(i, arr3[i])
	}

	// 遍历数组方法2
	for i := range arr3 {
		fmt.Println(i, arr3[i])
	}

	// 遍历数组方法3
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// 遍历数组方法4
	for _, v := range arr3 {
		fmt.Println(v)
	}
}

// 定义二维数组的方法
func define_two_dimensional_array() {
	var grid [4][5]int // 4个长度为5的int
	fmt.Println(grid)
}

// 数组是值类型（好好体会）
// 数组作为参数传递时，会复制原数组
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 指向数组的指针
func updateArraybypointer(arr *[5]int) {
	// (*arr)[0] = 200
	arr[0] = 200 // go中的指针很灵活，这里arr[0] = 200 等价于 (*arr)[0] = 200，两者都可以
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 切片实现更改数组的值
func updateArraybyslice(arr []int) {
	arr[0] = 200
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	define_linear_array()
	define_two_dimensional_array()

	var arr1 [5]int // 初始化当前类型的默认零值, 5个int
	arr2 := [3]int{10, 20, 30}
	arr3 := [...]int{40, 50, 60, 70, 80}

	fmt.Println("数组是值类型测试：")
	printArray(arr1)
	printArray(arr3)
	// printArray(arr2)  // 报错，[3]int与[5]int是不同的数据类型，不能互相传递参数
	fmt.Println("数组传递：测试是否修改了原数组的数据：")
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	fmt.Println("参数通过数组指针传递")
	updateArraybypointer(&arr1)
	updateArraybypointer(&arr3)

	fmt.Println("指针传递：测试是否修改了原数组的数据：")
	fmt.Println(arr1)
	fmt.Println(arr3)

	fmt.Println("切片参数传递实现更改数组的值")
	updateArraybyslice(arr1[:])
	fmt.Println(arr1)
	updateArraybyslice(arr3[:])
	fmt.Println(arr3[:])
}
