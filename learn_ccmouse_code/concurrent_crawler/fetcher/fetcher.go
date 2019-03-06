package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// 返回channel
var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	// 限制访问频率
	// 如果是100个worker, 则相当于每秒10个request
	<-rateLimiter

	resp, err := http.Get(url)
	if err != nil {
		return nil, err // 如果出错，把出错返回出去
	}
	// 关闭
	defer resp.Body.Close()
	// 检查http code
	if resp.StatusCode != http.StatusOK {
		// fmt.Println("Error: status code ", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	// 确定页面编码
	e := determineEncoding(bodyReader)
	// 转换编码，解决中文问题
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	// 读取页面内容
	// all, err := ioutil.ReadAll(utf8Reader)
	// if err != nil {
	// 	panic(err)
	// }
	return ioutil.ReadAll(utf8Reader) // 直接返回，因为两者的返回一样
}

// 猜测页面编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// 抽取1024
	bytes, err := r.Peek(1024)
	if err != nil {
		// panic(err)
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	// 这里不严格限制，只返回编码类型
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
