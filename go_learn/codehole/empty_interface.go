package main

import (
	"fmt"
)

func main(){
	// 类型推导以及赋值
	var user = map[string]interface{}{
		"age": 30,
        "address": "Beijing Tongzhou",
        "married": true,
	}

	fmt.Println(user)
	// 类型断言，接口类型转换成目标类型
	// user字典中变量的类型是map[string]interface{}，字典中得到的value类型是interface{}，需要通过类型转换才能得到期望的变量
	var age = user["age"].(int)
	var address = user["address"].(string)
	var married = user["married"].(bool)

	fmt.Println(age, address, married)
}