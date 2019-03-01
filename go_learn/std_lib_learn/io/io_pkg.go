package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

// 文件路径常量定义
const filepath = "src\\go_pratice_code\\go_learn\\std_lib_learn\\io\\iowriteAt.txt"

func ioReader_1() {
	reader := strings.NewReader("GO语言学习园地")
	p := make([]byte, 6)
	// 从指定位置读取
	n, err := reader.ReadAt(p, 5)
	// fmt.Println(err)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s, %d\n", p, n)
}

func ioReader_2() {
	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	// 关闭文件
	defer file.Close()
	// 把string写入文件
	n, err := file.WriteString("Golang中文社区 - xxxxxxx")
	if err != nil {
		panic(err)
	}
	fmt.Printf("write string length n = %d\n", n)
	// 在文件指定位置写入字符串
	n, err = file.WriteAt([]byte("Golang学习园地！"), 21)
	if err != nil {
		panic(err)
	}
	fmt.Println("n = ", n)
}

func ioWriter_1() {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(os.Stdout)
	_, err = writer.ReadFrom(file)
	if err != nil {
		panic(err)
	}

	writer.Flush()
}

func ioWriter_2() {
	reader := bytes.NewReader([]byte("Go语言学习园地"))
	reader.WriteTo(os.Stdout)
}

func ioReader_3() {
	reader := strings.NewReader("Go语言学习园地")
	n, err := reader.Seek(-6, os.SEEK_END)
	if err != nil {
		panic(err)
	}
	fmt.Printf("seek n = %d\n", n)

	ch, size, err := reader.ReadRune()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ch = %c\n", ch)
	fmt.Printf("size = %d\n", size)
}

//func main() {
	// ioReader_1()
	// ioReader_2()
	// ioWriter_1()
	// ioWriter_2()
	// ioReader_3()
//}

func main() {
	p:=&person{name:"张三"}
	p.modify() //指针接收者，修改有效
	fmt.Println(p.String())
}

type person struct {
	name string
}

func (p person) String() string{
	return "the person name is "+p.name
}

func (p *person) modify(){
	p.name = "李四"
}