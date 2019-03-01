package main

import "fmt"

type Handler interface {
	Do(k, v interface{})
}

// 定义了一个新的类型HandlerFunc，它是一个func(k, v interface{})类型
type HandlerFunc func(k, v interface{})

// 新的HandlerFunc实现了Handler接口（函数实现Handler接口）
// Do方法的实现是调用HandlerFunc本身，因为HandlerFunc类型的变量就是一个方法
func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

// 新增了一个EachFunc函数，帮助调用者强制转型，调用者就不用自己做了。
func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
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
