package main

import (
	"encoding/json"
	"errors"
	"fmt"
	myerr "go_pratice_code/custom_error/errors" // 包重命名
)

// error返回值是内置的接口
// type error interface {
//     Error() string
// }
func myErr() error {
	err := myerr.NewError(2, "err test")
	return err
}

func staErr() error {
	m := make(map[string]string)

	err := json.Unmarshal([]byte(""), m)
	if err != nil {
		return err
	}

	return errors.New("aaaa")
}

func main() {
	err1 := staErr()
	fmt.Println("---sta err: ---", err1.Error())

	err2 := myErr()
	fmt.Println("--- my err: ---", err2.Error())

}
