package main

import (
	"fmt"
	"os"
)

func main() {
	// 文件路径
	filePath := "E:\\VSCodeWorkspace\\GoWorkspace\\src\\go_pratice_code\\go_learn\\codehole\\errors\\read_file\\test.txt"
	// 打开文件，返回文件句柄指针和错误信息
	// go中函数可以返回多个值，一般返回两个：其中一个是错误信息error
	// 除函数之外，字典中也会返回两个值，用来判断字典中该key是否存在：var score, ok := scores["apple"]
	file, err := os.Open(filePath)
	// 错误处理， err=nil表示打开文件正常，否则打开文件异常
	if err != nil {
		fmt.Println("open file failed reason :", err.Error()) // err.Error() 与 err打印效果一样
		return
	}
	// 关闭文件，defer会在函数退出时执行，确保文件关闭
	// defer关键字，它将文件的关闭调用推迟到当前函数的尾部执行，即使后面的代码抛出了异常，文件关闭也会确保被执行，相当于python中的finally
	defer file.Close()

	// fmt.Println(file)

	// 保存文件内容的字节切片
	var content = []byte{} // 等价于 var content []byte   // content = []
	// fmt.Println(content)
	// 读缓冲切片，len和cap都为100字节
	var buf = make([]byte, 100)

	// buf不能按如下方式初始化
	// len=0，则死循环，一直读取0字节，导致永远读取不完文件
	// var buf = make([]byte, 0, 100)

	for {
		// 1次读取buf大小的数据
		// 返回读取的字节数和发生的错误
		// 读到文件尾EOF, n = 0
		n, err := file.Read(buf)
		fmt.Println("***n:", n)
		// 错误处理
		if err != nil {
			fmt.Println("---n:", n)
			break // 读取完或读错误退出for循环
		}

		if n > 0 {
			fmt.Println("###n: ", n)
			/*
				append 函数可以将单个元素追加到切片中，其实 append 函数可以一次性追加多个元素，它的参数数量是可变的:
				var s = []int{1,2,3,4,5}
				s = append(s,6,7,8,9)
			*/

			// 但是读文件的代码中需要将整个切片的内容追加到另一个切片中，这时候就需要 … 操作符，
			// 它的作用是将切片参数的所有元素展开后传递给 append 函数。你可能会担心如果切片里有成百上千的元素，展开成元素再传递会不会非常耗费性能。这个不必担心，展开只是形式上的展开，在实现上其实并没有展开，传递过去的参数本质上还是切片
			content = append(content, buf[:n]...)
		}
	}

	// 类型转换：[]byte直接转换成string类型
	// string 不能直接和byte数组转换
	// string 可以直接和byte切片转换
	str := string(content)
	length := len(str)

	fmt.Println("str:", str)
	fmt.Println("length:", length)
}
