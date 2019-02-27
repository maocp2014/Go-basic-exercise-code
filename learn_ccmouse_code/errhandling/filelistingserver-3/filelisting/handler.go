package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

// 实现error接口的Error方法
func (e userError) Error() string {
	return e.Message()
}

// 实现 Message 方法
func (e userError) Message() string {
	return string(e)
}

// 业务逻辑
func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		// 返回可以给用户看的错误
		return userError(fmt.Sprintf("path %s must start with %s", request.URL.Path, prefix))
	}

	path := request.URL.Path[len(prefix):]

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)

	return nil
}
