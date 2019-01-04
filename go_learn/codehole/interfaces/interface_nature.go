package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s interface{} // 空接口，值为nil
	fmt.Println(s)  // nil
	fmt.Println(unsafe.Sizeof(s))  // 16

	var arr [10]int = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(unsafe.Sizeof(arr)) // 80

	s = arr
	fmt.Println(s, unsafe.Sizeof(s)) // [1 2 3 4 5 6 7 8 9 10] 16
}