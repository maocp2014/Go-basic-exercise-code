package main

import (
	"fmt"
	"go_pratice_code/learn_ccmouse_code/interfaces/retriever/mock"
	"go_pratice_code/learn_ccmouse_code/interfaces/retriever/real"
	"time"
)

// Retriever接口
// 接口变量中有两个值：实现者的类型 和 实现者的值
type Retriever interface {
	Get(url string) string
}

// Poster接口
type Poster interface {
	Post(url string, form map[string]string) string
}

// 常量
const url = "http://www.baidu.com"

func post(poster Poster) {
	poster.Post(url, 
	map[string]string{
		"name" : "xxx",
		"course" : "golang",
	})
}

// 使用者
func download(r Retriever) string {
	return r.Get(url)
}

// 接口组合
type RetrieverPoster interface {
	Retriever
	Poster
}

// 使用组合接口
func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another fake baidu.com",
	})
	return s.Get(url)
}

// 查看interface变量有两种方法：1、type switch 2、type assertion
// 1、type switch
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}

func main() {
	var r Retriever
	// retriever实现了上述两个接口
	retriever := mock.Retriever{Contents: "This is a fake imooc.com"}
	// fmt.Println(download(r))
	r = &retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	// fmt.Println(download(r))

	// 2、Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}
