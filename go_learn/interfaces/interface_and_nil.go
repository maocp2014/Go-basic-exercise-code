package main

import (
	"fmt"
	"reflect"
)

type State struct{}

func testnil1(a, b interface{}) bool {
	return a == b
}

func testnil2(a *State, b interface{}) bool {
	return a == b
}

func testnil3(a interface{}) bool {
	return a == nil
}

func testnil4(a *State) bool {
	return a == nil
}

func testnil5(a interface{}) bool {
	v := reflect.ValueOf(a)
	return !v.IsValid() || v.IsNil()
}

func main() {
	var a *State
	fmt.Println(testnil1(a, nil))
	fmt.Println(testnil2(a, nil))
	fmt.Println(testnil3(a))
	fmt.Println(testnil4(a))
	fmt.Println(testnil5(a))
}

/*
false
false
false
true
true

解释：
一个interface{}类型的变量包含了2个指针，一个指针指向值的类型，另外一个指针指向实际的值
对一个interface{}类型的nil变量来说，它的两个指针都是0；但是var a *State传进去后，指向的类型的指针不为0了，因为有类型了， 
所以比较为false。 interface 类型比较， 要是两个指针都相等，才能相等。
*/