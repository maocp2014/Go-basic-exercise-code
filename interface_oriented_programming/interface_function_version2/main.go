package main

import "fmt"

type Handler interface {
	Do(k, v interface{})
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

// 定义了一个新的类型HandlerFunc，它是一个func(k, v interface{})类型
type HandlerFunc func(k, v interface{})

// 新的HandlerFunc实现了Handler接口
// Do方法的实现是调用HandlerFunc本身，因为HandlerFunc类型的变量就是一个方法
func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

type welcome string

func (w welcome) selfInfo(k, v interface{}) {
	fmt.Printf("%s,我叫%s,今年%d岁\n", w, k, v)
}

func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26

	var w welcome = "大家好"

	// 还是差不多原来的实现，只是把方法名Do改为selfInfo。HandlerFunc(w.selfInfo)不是方法的调用，而是转型，
	// 因为selfInfo和HandlerFunc是同一种类型，所以可以强制转型。
	// 转型后，因为HandlerFunc实现了Handler接口，所以我们就可以继续使用原来的Each方法了。
	Each(persons, HandlerFunc(w.selfInfo))
}
