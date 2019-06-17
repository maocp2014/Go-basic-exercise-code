package main

import "fmt"

// 接口型函数：指的是用函数实现接口，这样在调用的时候就会非常简便，这种函数为接口型函数，这种方式适用于只有一个函数的接口。

type Handler interface {
	Do(k, v interface{})
}

// 定义了一个新的类型HandlerFunc，它是一个func(k, v interface{})类型
// 这个类型只定义了函数的参数列表，函数参数列表与接口定义的方法一致
type HandlerFunc func(k, v interface{})

// 新的HandlerFunc实现了Handler接口（函数实现Handler接口）
// Do方法的实现是调用HandlerFunc本身，因为HandlerFunc类型的变量就是一个方法
func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v) // 接口的实现中调用自己本身。这样就使得可以用函数来实现接口功能，而不是定义类型并实现接口来实现接口功能
}

// 第1种方法
func Each(m map[interface{}]interface{}, h Handler) { // 传入一个实现了Handler接口的类型的实例
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

// 第2种方法
// 新增了一个EachFunc函数，帮助调用者强制转型，调用者就不用自己做了。
// 第二种方式可以只传入一个函数，只要求参数列表一致，函数名字可随便起，类型也不用新定义，用起来很方便。
func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) { // 传入了一个参数列表为接口所需实现函数的参数
	Each(m, HandlerFunc(f)) // HandlerFunc(f)还是类型转换
}

// EachFunc函数接收的是一个func(k, v interface{})类型的函数，没有必要实现Handler接口了，所以我们新的类型可以去掉不用了。
func selfInfo(k, v interface{}) {
	fmt.Printf("大家好，我叫%s,今年%d岁\n", k, v)
}

// 去掉了自定义类型welcome之后，整个代码更简洁
func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26

	EachFunc(persons, selfInfo)
}

/*
延伸:
以上关于函数型接口就写完了，如果我们仔细留意，发现和我们自己平时使用的http.Handle方法非常像，其实接口http.Handler就是这么实现的。
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

func Handle(pattern string, handler Handler) {
	DefaultServeMux.Handle(pattern, handler)
}

func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}
*/
