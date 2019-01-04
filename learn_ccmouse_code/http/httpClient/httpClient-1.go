package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main(){
	url := "http://www.imooc.com"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Body)

	s, err := httputil.DumpResponse(resp, true)
	if err !=nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}