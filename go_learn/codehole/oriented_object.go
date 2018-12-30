package main

import "fmt"

type Fruit struct {}

func (f Fruit) eat() {
    fmt.Println("eat fruit")
}

func (f Fruit) enjoy() {
    fmt.Println("smell first")
    f.eat()
    fmt.Println("clean finally")
}

type Apple struct {
    Fruit
}

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
    apple.enjoy()
    banana.enjoy()
}