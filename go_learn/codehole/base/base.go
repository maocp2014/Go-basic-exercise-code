package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x      int
	y      int
	Radius int
}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}

func (c Circle) expand_by_value_receiver() {
	c.Radius *= 2
}

func (c *Circle) expand_by_pointer_receiver() {
	c.Radius *= 2
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * float64(c.Radius)
}

func struct_method() {
	var c Circle = Circle{Radius: 50}
	area1 := c.Area()

	circum1 := c.Circumference()
	println(area1, circum1)

	var pc *Circle = &c
	area2 := pc.Area()
	circum2 := pc.Circumference()
	println(area2, circum2)
}

// 内嵌结构体
type Point struct {
	x int
	y int
}

func (p Point) show() {
	fmt.Println(p.x, p.y)
}

type CircleLn struct {
	loc Point
	Radius int
}
func struct_inline() {
	var c CircleLn = CircleLn {
		loc : Point {
			x : 100,
			y : 100,
		},
		Radius : 50,
	}
	fmt.Printf("%+v\n", c)
	fmt.Printf("%+v\n", c.loc)
	fmt.Printf("%d %d\n", c.loc.x, c.loc.y)

	c.loc.show()
}

func main() {
	fmt.Println("welcome to go!")

	// struct_method()
	struct_inline()
}
