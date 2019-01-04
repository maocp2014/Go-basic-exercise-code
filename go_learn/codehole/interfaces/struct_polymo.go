package main

import (
	"fmt"
)

// 接口变量
type Fruitable interface {
	eat()
}

// Fruit结构体
type Fruit struct {
	Name      string
	Fruitable // 内嵌匿名接口变量
}

// 结构体方法
func (f Fruit) want() {
	fmt.Printf("I like ")
	f.eat()
}

type Apple struct{}

func (a Apple) eat() {
	fmt.Println("eating apple!")
}

type Banana struct{}

func (b Banana) eat() {
	fmt.Println("eating banana!")
}

func main() {
	// 方式1
	// var f1 Fruit = Fruit{"Apple", Apple{}}
	// var f2 Fruit = Fruit{"Banana", Banana{}}

	// 方式2: 内嵌匿名变量传值方法
	var f1 Fruit = Fruit{Name: "Apple", Fruitable: Apple{}}
	var f2 Fruit = Fruit{Name: "Banana", Fruitable: Banana{}}

	fmt.Printf("%T\n", Apple{})
	fmt.Printf("%T\n", Banana{})

	f1.want()
	f2.want()
}
