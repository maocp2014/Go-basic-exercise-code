package main

import (
	"fmt"
	"math"
	"time"
	"unsafe"
)

func pointerfun() {
	// 自动推导
	// i := 42
	// p := &i

	var i int = 42
	var p *int = &i
	fmt.Println(p, &i, *p)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func sig(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	} else {
		return 0
	}
}

// 值匹配
func switch_value(score float32) string {
	switch score / 10 {
	case 0, 1, 2, 3, 4, 5:
		return "差"
	case 6, 7:
		return "及格"
	case 8:
		return "良"
	default:
		return "优"
	}
}

// 表达式匹配
func switch_expression(score float32) string {
	switch {
	case score < 60:
		return "差"
	case score < 80:
		return "及格"
	case score < 90:
		return "良"
	default:
		return "优"
	}
}

func gotype() {
	var i int8 = 1
	var s string = "hello, string!"
	var ui uint8 = 1
	var f float32 = 3.2
	fmt.Println(i, s, ui, f)
}

// 替代其它语言中的while
func for_type1() {
	// for {
	// 	fmt.Println("hello for loop!")
	// }

	for true {
		fmt.Println("hello for loop")
		time.Sleep(2000)
	}
}

func for_loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("hello %d\n", i)
	}

}

func array_1() {
	var arr [6]int
	fmt.Println(arr)

	var a = [3]int{1, 2, 3}
	var b [4]int = [4]int{1, 2, 3, 4}
	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func array_2() {
	var squares [9]int

	for i := 0; i < len(squares); i++ {
		squares[i] = (i + 1) * (i + 1)
	}
	fmt.Println(squares)
}

// 数组越界检查（原理是在编译后的代码中插入下标检测的逻辑）
func array_3() {
	var arr [3]int = [3]int{1, 2, 3}

	// arr[10] = 10

	var b int = 10
	arr[b] = 10

	fmt.Println(arr)
}

// 只有元素类型和长度相同的数组才能互相赋值，浅拷贝
func array_assign() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	var b [5]int
	// var b [6]int  // 属于不同类型，不能相互赋值

	b = a
	a[0] = 100

	fmt.Println(a)
	fmt.Println(b)
}

func array_traversing() {
	var a [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 方式1
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }

	// 方式2
	for i := range a {
		fmt.Println(i, a[i])
	}

	// 方式3
	for index, value := range a {
		fmt.Println(index, value)
	}
}

func slice_create() {
	var s1 []int = make([]int, 8) //cap 与 len 相同
	// var s1 = make([]int, 8)
	var s2 []int = make([]int, 5, 8)
	// s2 := make([]int, 5, 8)

	var s3 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3, len(s3), cap(s3))
}

func empty_slice() {
	var s1 []int                  // nil切片
	var s2 []int = []int{}        //空切片
	var s3 []int = make([]int, 0) //空切片

	fmt.Println(s1, s2, s3)
	fmt.Println(len(s1), len(s2), len(s3))
	fmt.Println(cap(s1), cap(s2), cap(s3))
}

func slice_assign() {
	var s1 []int = make([]int, 5, 8)

	for i := 0; i < len(s1); i++ {
		s1[i] = i + 1
	}

	var s2 []int = s1
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	s1[0] = 100
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
}

func slice_traversing() {
	var s1 []int = make([]int, 5, 8)

	for i := 0; i < len(s1); i++ {
		s1[i] = i + 1
	}

	// 方式1
	fmt.Println("subscript of for ")
	for i := 0; i < len(s1); i++ {
		fmt.Println(i, s1[i])
	}

	// 方式2
	fmt.Println("range 1 of for")
	for index := range s1 {
		fmt.Println(index, s1[index])
	}

	// 方式3
	fmt.Println("range 2 of for")
	for index, value := range s1 {
		fmt.Println(index, value)
	}

	// 方式4
	fmt.Println("rang 3 of for")
	for _, value := range s1 {
		fmt.Println(value)
	}
}

func slice_append() {
	var s1 []int = []int{1, 2, 3, 4, 5}
	fmt.Printf("%d %d %d %p\n", s1, len(s1), cap(s1), &s1)

	var s2 []int = append(s1, 6)
	fmt.Printf("%d %d %d %p\n", s1, len(s1), cap(s1), &s1)
	fmt.Printf("%d %d %d %p\n", s2, len(s2), cap(s2), &s2)

	var s3 []int = append(s2, 7)
	fmt.Printf("%d %d %d %p\n", s2, len(s2), cap(s2), &s2)
	fmt.Printf("%d %d %d %p\n", s3, len(s3), cap(s3), &s3)
}

func slice_cut() {
	var s1 []int = []int{1, 2, 3, 4, 5, 6, 7}
	var s2 []int = s1[:5]
	var s3 []int = s2[3:]
	var s4 []int = s3[:]

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
	fmt.Println(s4, len(s4), cap(s4))
}

func array_to_slice() {
	var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var s []int = arr[2:6]
	fmt.Println(arr, len(arr))
	fmt.Println(s, len(s), cap(s))

	arr[4] = 100
	fmt.Println(arr, len(arr))
	fmt.Println(s, len(s), cap(s))

	s[0] = 1000
	fmt.Println(arr, len(arr))
	fmt.Println(s, len(s), cap(s))
}

func slice_copy() {
	var s []int = make([]int, 5, 8)
	for i := 0; i < len(s); i++ {
		s[i] = i + 1
	}
	fmt.Println(s)

	var d []int = make([]int, 2, 6)
	n := copy(d, s)
	fmt.Println(n, d)
}

func slice_expansion() {
	s1 := make([]int, 6)
	s2 := make([]int, 1024)
	s1 = append(s1, 1)
	s2 = append(s2, 2)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
}

func map_create() {
	var m1 map[int]string
	var m2 map[int]string = make(map[int]string)
	var m3 map[string]int = map[string]int{
		"优秀": 80,
		"良好": 76, // 字符串只能用双引号或反引号，不能用单引号
		"及格": 60, // 此处逗号不能省略，否则报语法错误
	}

	fmt.Println(m1, len(m1))
	fmt.Println(m2, len(m2))
	fmt.Println(m3, len(m3))
}

func map_ops() {
	m := map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}
	fmt.Println(m, len(m))
	// 读取元素
	price, ok := m["pear"]
	if ok {
		fmt.Println(price)
	} else {
		fmt.Println("apple key does not exist!")
	}

	// 增加或修改元素
	m["pear"] = 0
	price, ok = m["pear"]
	if ok {
		fmt.Println(price)
	} else {
		fmt.Println("pear key does not exist!")
	}
	fmt.Println(m, len(m))

	// 删除元素
	delete(m, "pear")
	fmt.Println(m, len(m))
}

func map_traversing() {
	var m = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}
	fmt.Println(m, len(m))

	for name, price := range m {
		fmt.Println(name, price)
	}

	for _, price := range m {
		fmt.Println(price)
	}

	var names []string = make([]string, 0, len(m)) // 空切片，注意参数0
	var prices []int = make([]int, 0, len(m))      // 空切片

	fmt.Println(names, len(names))
	fmt.Println(prices, len(prices))

	for name, price := range m {
		names = append(names, name)
		prices = append(prices, price)
	}

	fmt.Println(names, len(names))
	fmt.Println(prices, len(prices))
}

// map占用的内存大小，一个机器字
func map_sizeof() {
	var m = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}
	fmt.Println(m, len(m), unsafe.Sizeof(m))
}

// 字符串按字节遍历
func string_traversing_by_byte() {
	var s string = "嘻哈china"

	fmt.Println(s, len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

// 字符串按rune遍历
func string_traversing_by_rune() {
	var s string = "嘻哈china"

	for codepoint, runeValue := range s {
		// fmt.Println(codepoint, runeValue)
		fmt.Printf("%d %s %d\n", codepoint, int32(runeValue))
	}
}

// 字符串是只读的
func string_is_readonly() {
	var s string = "hello"

	//s[0] = 'H'  // 不能下标赋值，字符串是只读的
	fmt.Println(s)
}

// 字符串切割
func string_split() {
	var s string = "hello world!"

	var s1 = s[2:8]
	// s1[2] = 'x'   // 不能下标赋值，字符串是只读的
	s2 := s[3:]
	s3 := s[:7]

	s4 := s1[2:6]
	fmt.Printf("%s\n%s\n%s\n%s\n%s", s, s1, s2, s3, s4)
}

func byte_string_conversion() {
	// 字符串与字节切片不共享底层字节数组
	var s string = "hello world!" // 只读
	var b []byte = []byte(s)      // 可以修改

	ss := string(b)

	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(ss)

	b[1] = 111 // 验证字节切片可修改
	fmt.Println(b)
}

func struct_create() {
	// 结构体定义
	type Circle struct {
		x      int
		y      int
		Radius int
	}

	// 创建结构体变量
	// 方式1: kv形式（map）
	// var c1 Circle = Circle {
	// 	x : 100,
	// 	y : 100,
	// 	Radius : 50,
	// }
	// fmt.Printf("%+v\n", c1)

	// var c2 Circle = Circle { Radius : 50 }  //这个不需要加逗号
	// var cc2 Circle = Circle {
	// 	Radius : 50,  //这样必须要加逗号
	// }

	// fmt.Printf("%+v\n", c2)

	// var c3 Circle = Circle {} // 零值结构体，所有字段都初始化成相应字段的零值
	// fmt.Printf("%+v\n", c3)

	// 方式2: 顺序形式，初始化值依次赋值且都不能缺省

	// var c4 Circle = Circle {100, 100, 50}
	// fmt.Printf("%+v\n", c4)

	// 方式3: 全局new函数

	// 结构体的指针类型
	// var c5 *Circle = &Circle {100, 100, 50}
	// fmt.Printf("%+v\n", c5)

	// var c6 *Circle = new(Circle)  // 零值结构体，所有字段都初始化成相应字段的零值
	// fmt.Printf("%+v\n", c6)

	// 方式4
	// var c7 Circle   // 零值结构体，所有字段都初始化成相应字段的零值
	// fmt.Printf("%+v\n", c7)

	// 综上，有3种零值结构体
	// var c3 Circle = Circle {}
	// var c6 Circle = new(Circle)
	// var c7 Circle
}

func zero_nil_struct() {
	type Circle struct {
		x      int
		y      int
		Radius int
	}

	var c1 *Circle = nil // nil指针表示无指向的指针
	var c2 Circle

	fmt.Println(c1, unsafe.Sizeof(c1))
	fmt.Println(c2, &c2, unsafe.Sizeof(c2))
	fmt.Printf("%+v", &c2, unsafe.Sizeof(&c2))
}

func struct_copy() {
	type Circle struct {
		x      int
		y      int
		Radius int
	}

	var c1 Circle = Circle{Radius: 50} // 逗号可以不加
	var c2 Circle = c1
	fmt.Printf("%+v\n", c1)
	fmt.Printf("%+v\n", c2)

	c1.Radius = 100
	fmt.Printf("%+v\n", c1) // {x:0 y:0 Radius:100}
	fmt.Printf("%+v\n", c2) // {x:0 y:0 Radius:50}

	var c3 *Circle = &Circle{Radius: 50}
	var c4 *Circle = c3
	fmt.Printf("%+v\n", *c3)
	fmt.Printf("%+v\n", *c4)

	c3.Radius = 100
	fmt.Printf("%+v\n", *c3) // {x:0 y:0 Radius:100}
	fmt.Printf("%+v\n", *c4) // {x:0 y:0 Radius:100}

}

func struct_array_slice() {
	type ArrayStruct struct {
		value [10]int
	}

	type SliceStruct struct {
		value []int
	}

	var arr ArrayStruct = ArrayStruct{[...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	var sli SliceStruct = SliceStruct{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	fmt.Println(unsafe.Sizeof(arr), unsafe.Sizeof(sli))
}

type Circle struct {
	x      int
	y      int
	Radius int
}

func struct_expandByValue(c Circle) {
	c.Radius *= 2
}

func Struct_expandByPointer(c *Circle) {
	c.Radius *= 2
}

// 函数也是值
func func_is_value() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot) // hypot是指向函数的指针？
	fmt.Println(hypot(3, 4))
}

// 闭包, 返回值是函数，该函数被绑定在各自的sum变量上
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	fmt.Println("welcome to go!")
	/*
		fmt.Println("hello go!")
		pointerfun()
		gotype()

		max := max(4, 3)
		print("max:", max)

		min := min(32, 87)
		print("\nmin:", min)

		sig := sig(0)
		print("\nsig:", sig)

		flag := switch_value(70)
		print("\nscore level:", flag)

		flag = switch_expression(89)
		print("\nscore level:", flag)

		for_type1()

		print("\n")
		for_loop()

		array_1()

		array_2()

		array_3()

		array_assign()

		array_traversing()

		slice_create()

		empty_slice()

		slice_assign()

		slice_traversing()

		slice_append()

		slice_cut()

		array_to_slice()

		slice_copy()

		slice_expansion()

		map_create()

		map_ops()

		map_traversing()

		map_sizeof()

		string_traversing_by_byte()
		string_traversing_by_rune()

		string_split()

		byte_string_conversion()

		struct_create()

		zero_nil_struct()

		struct_copy()

		struct_array_slice()*/

	// var c = Circle { Radius : 50 }
	// struct_expandByValue(c)
	// fmt.Println(c)

	// Struct_expandByPointer(&c)
	// fmt.Println(c)

	func_is_value()

	// pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		// fmt.Println(pos(i), neg(-2 * i))
	}

	// v := &Vertex{3, 4}
	// fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2) // Sqrt2 = 1.41421356237309504880168872420969807856967187537694807317667974 // https://oeis.org/A002193
	fmt.Println(f)
	fmt.Println(f.Abs())
}
