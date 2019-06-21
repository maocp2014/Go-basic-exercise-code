package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var (
	STATIC = ".\\src\\go_pratice_code\\goWeb\\demo04\\static\\"
	JS     = STATIC + "js\\"
	HTML   = STATIC + "html\\"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
	// t := template.New("index")
	// t.Parse("<div id='templateTextDiv'>Hi, {{.name}}, {{.someStr}}</div>")

	t, err := template.ParseFiles(HTML + "index.html")
	if err != nil {
		panic(err)
	}

	data := map[string]string{
		"name":    "zeta",
		"someStr": "这是一个开始",
	}

	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", myWeb) //  方法形式（myweb）

	// 指定相对路径static为文件服务路径，http.Dir()为类型转换
	// staticHandle := http.FileServer(http.Dir("src\\go_pratice_code\\goWeb\\demo04\\static"))
	//将/js/路径下的请求匹配到 ./static/js/下
	// http.Handle("/js/", staticHandle) // 接口形式（staticHandle实现了ServeHTTP方法，即Handler接口）

	http.Handle("/js/", http.FileServer(http.Dir(STATIC))) // 对应到 STATIC/js   这样所有请求的路径都必须匹配一个static目录下的子目录

	// http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir(JS)))) // 对应到./js 这样，浏览器中访问/js/时，直接对应到./static目录下，不需要再加一个/js/子目录。

	fmt.Println("服务器即将开启，访问地址 http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器开启错误: ", err)
	}
}
