package main

import (
	"fmt"
	"os"
)

const fsrc = "src\\go_pratice_code\\go_learn\\codehole\\errors\\defer\\source.txt"
const fdst = "src\\go_pratice_code\\go_learn\\codehole\\errors\\defer\\target.txt"

func main() {
	fsrc, err := os.Open(fsrc)
	if err != nil {
		fmt.Println("open source file failed")
		return
	}
	defer func() {
		fmt.Println("close source file")
		fsrc.Close()
	}()

	fdes, err := os.Open(fdst)
	if err != nil {
		fmt.Println("open target file failed")
		return
	}
	defer func() {
		fmt.Println("close target file")
		fdes.Close()
	}()

	fmt.Println("do something here")
}
