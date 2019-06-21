package main

import (
	"fmt"
	"net/http"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for k, v := range r.URL.Query() {
		fmt.Println("key: ", k, "value: ", v)
	}

	for k, v := range r.PostForm {
		fmt.Println("key:", k, ", value:", v[0])
	}

	fmt.Fprintf(w, "这是一个开始")
}

func main() {
	http.HandleFunc("/", myWeb)
	fmt.Println("服务器开启，访问地址 http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器开启错误：", err)
	}
}
