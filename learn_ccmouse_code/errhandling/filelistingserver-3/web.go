package main

import (
	"go_pratice_code/learn_ccmouse_code/errhandling/filelistingserver-3/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

// 统一出错处理
// 利用了函数式编程：输入为函数，输出为函数（思想：将输入函数包装后再返回新的函数）
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// recover
		// recover服务器内部错误，增强页面友好度
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		// 业务处理
		err := handler(w, r)
		// 出错处理
		if err != nil {
			// log.Warn("Error handling request: %s", err.Error()) // log.Warn需要安装第3方库
			log.Printf("Error occurred handling request: %s", err.Error()) // 标准库

			// 判断是不是用户可见的错误
			// type assertion
			if userErr, ok := err.(userError); ok {
				http.Error(w, userErr.Message(), http.StatusBadRequest)
				return
			}

			// 用户不可见的错误处理
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

// 定义接口
// 区分可以给用户看的错误以及不能给用户看的错误
type userError interface {
	error            // 给系统看的错误
	Message() string // 给用户看的错误
}

func main() {
	// 包装输入函数
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

// 浏览器输入http://localhost:8888/list/src/go_pratice_code/learn_ccmouse_code/errhandling/defer/fib.txt
