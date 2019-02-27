package main

import (
	"fmt"
	"sort"
)

// 定义Person结构体
type Person struct {
	Name string
	Age  int
}

// 为Person定义方法，实现string接口
func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// 类型重命名
// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []Person

// 下面3个方法是为了实现sort interface
//为ByAge定义方法，实现sort interface接口中的Len方法
func (a ByAge) Len() int {
	return len(a)
}

//为ByAge定义方法，实现sort interface接口中的Swap方法
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

//为ByAge定义方法，实现sort interface接口中的Less方法
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	// fmt.Println(ByAge(people))
	// ByAge(people) 是类型转换
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
