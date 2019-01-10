package main

import "fmt"

/*
Go 语言不是面向对象语言在于它的结构体不支持多态，它不能算是一个严格的面向对象语言。多态是指父类定义的方法可以调用子类实现的方法，
不同的子类有不同的实现，从而给父类的方法带来了多样的不同行为。

Go 语言的结构体明确不支持这种形式的多态，外结构体的方法不能覆盖内部结构体的方法。

面向对象的多态性需要通过 Go 语言的接口特性来模拟
*/

// 空结构体
type Fruit struct {}

// 给结构体定义eat方法
func (f Fruit) eat() {
    fmt.Println("eat fruit")
}

// 给结构体定义enjoy方法
func (f Fruit) enjoy() {
    fmt.Println("smell first")
    f.eat()
    fmt.Println("clean finally")
}

// Apple结构体，组合自Fruit结构体
type Apple struct {
    Fruit
}

// Apple结构体方法eat
func (a Apple) eat() {
    fmt.Println("eat apple")
}

type Banana struct {
    Fruit
}

func (b Banana) eat() {
    fmt.Println("eat banana")
}

func main() {
    var apple = Apple {}
    var banana = Banana {}
    // enjoy 方法调用的 eat 方法还是 Fruit 自己的 eat 方法，它没能被外面的结构体方法覆盖掉。
    apple.enjoy()
    banana.enjoy()
}