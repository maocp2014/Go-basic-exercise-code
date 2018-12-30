package main

import (
	"fmt"
	"unsafe"
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

	var user map[string]interface{} = map[string]interface{}{
		"age":     30,
		"address": "Tongzhou Beijing",
		"married": true,
	}
	fmt.Println(user)

	// for key,value := range user {
	// 	fmt.Println(key, value)
	// }

	// age := user["age"]
	// address := user["address"]
	// married := user["married"]
	var age = user["age"]
	var address = user["address"]
	var married = user["married"]
	fmt.Println(age, address, married)
	fmt.Printf("%T\n%T\n%T\n", age, address, married)

	var s interface{} // 空接口，值为nil
	fmt.Println(s)
	fmt.Println(unsafe.Sizeof(s))

	var arr [10]int = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(unsafe.Sizeof(arr))

	s = arr
	fmt.Println(s, unsafe.Sizeof(s))
}
