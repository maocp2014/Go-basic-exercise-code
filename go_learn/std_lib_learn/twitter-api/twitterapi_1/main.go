package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// start with a string representation of our JSON data
var input = `
{
    "created_at": "Thu May 31 00:00:01 +0000 2012"
}
`

func main() {
	// our target will be of type map[string]interface{}, which is a
	// pretty generic type that will give us a hashtable whose keys
	// are strings, and whose values are of type interface{}

	var val map[string]interface{}

	// var val map[string]time.Time   // 直接定义time.Time类型会报错
	// 出错的原因是字符串格式与 Go 中的时间格式不匹配（因为 Twitter's API 是用 Ruby 写的，其格式跟 Go 不同）。
	// 我们必须定义我们自己的类型来解析时间。encoding/json 在解析时会判断传入 json.Unmarshal 的值是否实现了 json.Unmarshaler 接口：
	 

	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}

	fmt.Println(val)
	for k, v := range val {
		fmt.Println(k, reflect.TypeOf(v))
	}
}
