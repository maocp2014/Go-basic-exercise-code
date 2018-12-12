package main

import (
    "fmt"
)

// Can I convert a []T to an []interface{}?

// Not directly. It is disallowed by the language specification because the two types do not have the same representation in memory. It is necessary to copy the elements individually to the destination slice. This example converts a slice of int to a slice of interface{}:

// t := []int{1, 2, 3, 4}
// s := make([]interface{}, len(t))
// for i, v := range t {
//     s[i] = v
// }

/*
// 程序报错，names不能转换为[]interface{}类型
func PrintAll(vals []interface{}) {
    for _, val := range vals {
	    fmt.Println(val)
    }
}

func main() {
    names := []string{"stanley", "david", "oscar"}
    PrintAll(names)
}
*/

// 正确做法
func PrintAll(vals []interface{}) {
    for _, val := range vals {
	    fmt.Println(val)
    }
}

func main() {
	names := []string{"stanley", "david", "oscar"}
	vals := make([]interface{}, len(names))
    for i, v := range names {
		vals[i] = v
	}
    PrintAll(vals)
}