package main

import (
	// "golang.org/x/tour/pic"
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	var a [][]uint8 = make([][]uint8, dy)
	// fmt.Println("a:", a, len(a))
    for x := range a {
		fmt.Println("x:", x)
		b := make([]uint8, dx)
		// fmt.Println("b:", b, len(b))
		for y := range b {
			fmt.Println("y:", y)
			b[y] = uint8 (x * y)
			fmt.Printf("b[%d] = %d\n", y, b[y])
		}

		a[x] = b
		fmt.Printf("a[%d] = %d\n", x, a[x])
	}

	return a
}

func main() {
	Pic(4, 4)
	// pic.Show(Pic)
}