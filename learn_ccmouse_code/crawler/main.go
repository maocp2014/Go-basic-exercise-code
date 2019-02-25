package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("https://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	// 关闭
	defer resp.Body.Close()
	// 检查http code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code ", resp.StatusCode)
		return
	}
	// 确定页面编码
	e := determineEncoding(resp.Body)
	// 转换编码，解决中文问题
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	// 读取页面内容
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	// 抽取1024
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	// 这里不严格限制，只返回编码类型
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
