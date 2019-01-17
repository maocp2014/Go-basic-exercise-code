package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	if n <= 0 {
		panic(fmt.Sprintf("unsupported convert: %d", n))
	}

	// 省略起始条件
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	// 一行一行读
	scanner := bufio.NewScanner(file)
	// 省略起始条件、递增条件
	// for循环只有结束条件，这种形式的for相当于while循环
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever() {
	// 省略初始条件、递增条件、结束条件，相当于while true
	for {
		fmt.Println("abc")
	}
}
func main() {
	fmt.Println(
		convertToBin(13),
		convertToBin(5),
		// convertToBin(-1),
	)

	const filename = "src\\go_pratice_code\\learn_ccmouse_code\\basic\\branch\\abc.txt"
	printFile(filename)
}
