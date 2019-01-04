package main

import (
	"fmt"
)

type Smellable interface {
	smell()
}

type Eatable interface {
	eat()
}

// 接口组合继承（接口合并），Fruitable包含了smell()和eat()方法，等价于下述接口
// type Fruitable interface {
// 	eat()
// 	smell()
// }
type Fruitable interface {
	Smellable
	Eatable
}

type Apple struct{}

func (a Apple) smell() {
	fmt.Println("apple can smell!")
}

func (a Apple) eat() {
	fmt.Println("apple can eat!")
}

type Flower struct{}

func (f Flower) smell() {
	fmt.Println("flower can smell!")
}

func main() {
	var s1 Smellable
	var s2 Eatable
	var apple Apple = Apple{}    // 等价于var apple Apple
	var flower Flower = Flower{} // 等价于var flower Flower

	s1 = apple
	s1.smell()

	s1 = flower
	s1.smell()

	s2 = apple
	s2.eat()
}
