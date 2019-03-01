package main

import "fmt"

type Handler interface {
	Do(k, v interface{})
}

// 这个函数的功能，就是迭代传递过来的map参数，然后把map的每个key和value值传递给Handler的Do方法，
// 去做具体的事情，可以是输出，也可以是计算，具体由这个Handler的实现来决定，这也是面向接口编程。
func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

// welcome是我们新定义的类型，对应基本类型string
type welcome string

// welcome实现了Handler接口，打印出自我介绍
func (w welcome) Do(k, v interface{}) {
	fmt.Printf("%s, 我叫%s, 今年%d岁\n", w, k, v)
}

func main() {
	// 创建map
	// 1、方式1
	persons1 := make(map[interface{}]interface{}) // 空
	fmt.Println("person1: ", persons1)
	// 2、方式2
	var persons2 map[interface{}]interface{}
	// 这样创建了nil切片, 不能操作存储键值对的，必须要初始化后才可以，比如使用make函数,为其开启一块可以存储数据的内存，也就是初始化。
	// persons2["xxx"] = "yyyy"  // 直接操作报错 panic: assignment to entry in nil map
	// 初始化
	persons2 = make(map[interface{}]interface{})
	fmt.Println("person2: ", persons2)
	// 3、方式3
	persons3 := map[interface{}]interface{}{
		"张三": 20,
		"李四": 30,
		"王五": 40,
	}

	fmt.Println("person3: ", persons3)

	var w welcome = "大家好"
	// w实现了这个接口，具备这个功能，在这里可以使用
	Each(persons3, w)
}
