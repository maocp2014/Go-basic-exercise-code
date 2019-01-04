package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main(){
	url := "http://www.imooc.com"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	// 请求头
	header := "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1"
	// 请求中加入请求头
	request.Header.Add("User-Agent", header)

	// 发起请求
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	// 关闭
	defer resp.Body.Close()
	fmt.Println(resp.Body)
    // 响应内容处理 
	s, err := httputil.DumpResponse(resp, true)
	if err !=nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}