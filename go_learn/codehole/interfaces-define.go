package main

import "fmt"

// 接口：可以闻
type Smellable interface {
  smell()
}

// 接口：可以吃
type Eatable interface {
  eat()
}

// 定义苹果的结构体
type Apple struct {}

// 苹果结构体实现Smellable接口的smell方法
func (a Apple) smell() {
  fmt.Println("apple can smell")
}

// 实现Eatable接口的eat方法
func (a Apple) eat() {
  fmt.Println("apple can eat")
}

// 定义花结构体
type Flower struct {}

// 实现Smellable接口的smell方法
func (f Flower) smell() {
  fmt.Println("flower can smell")
}

func main() {
  var s1 Smellable  // 接口变量
  var s2 Eatable
  var apple = Apple{}
  var flower = Flower{}
  s1 = apple
  s1.smell()
  s1 = flower
  s1.smell()
  s2 = apple
  s2.eat()
}