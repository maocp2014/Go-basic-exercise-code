package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	// // 一行一行读
	// scanner := bufio.NewScanner(file)
	// // 省略起始条件、递增条件
	// // for循环只有结束条件，这种形式的for相当于while循环
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	printFileContents(file) // file文件名直接传递给io.reader
}

// 参数是io.reader
//这样做的好处是：不仅可以打印文件；也可以像打印文件一样打印字符串
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

func main() {
	const filename = "src\\go_pratice_code\\learn_ccmouse_code\\interfaces\\fileOps\\abc.txt"
	printFile(filename)

	s := `abc"d"
	kkkk
	1234
	
	p`
	printFileContents(strings.NewReader(s))
}
