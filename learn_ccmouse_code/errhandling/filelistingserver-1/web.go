package main

import (
	"go_pratice_code/learn_ccmouse_code/errhandling/filelistingserver-2/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

// 统一出错处理
// 利用了函数式编程：输入为函数，输出为函数（思想：将输入函数包装后再返回新的函数）
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// 业务处理
		err := handler(w, r)
		// 出错处理
		if err != nil {
			// log.Warn("Error handling request: %s", err.Error()) // log.Warn需要安装第3方库
			log.Printf("Error occurred handling request %s:", err.Error()) // 标准库
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				// http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}

func main() {
	// 包装输入函数
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

// 浏览器输入http://localhost:8888/list/src/go_pratice_code/learn_ccmouse_code/errhandling/defer/fib.txt
