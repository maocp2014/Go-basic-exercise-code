package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {}

func (c Cat) Speak() string {
	return "Meow!"
}

type L1ama struct {}

func (l L1ama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

type TestInterface struct {}

func main() {
	// var animals []Animal = []Animal {Dog{}, Cat{}, L1ama{}, JavaProgrammer{}, TestInterface{}} // 报错，TestInterface{}没有实现Speak()方法
	// 实现了Animal接口中Speak()方法的都可以认为是Animal类型（duck typing）
	var animals []Animal = []Animal {Dog{}, Cat{}, L1ama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
