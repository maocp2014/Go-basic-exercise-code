package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
	// t := template.New("index")
	// t.Parse("<div id='templateTextDiv'>Hi, {{.name}}, {{.someStr}}</div>")

	t, err := template.ParseFiles("src\\go_pratice_code\\goWeb\\demo03\\index.html")
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
	http.HandleFunc("/", myWeb)
	fmt.Println("服务器开启，访问地址 http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器开启错误：", err)
	}
}
