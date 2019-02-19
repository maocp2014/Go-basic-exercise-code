package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len("/list/"):]

		file, err := os.Open(path)
		if err != nil {
			// panic(err)
			// 错误会显示到web页面
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		all, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}

		writer.Write(all)
	})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

// 浏览器输入http://localhost:8888/list/src/go_pratice_code/learn_ccmouse_code/errhandling/defer/fib.txt
