package main

import "fmt"

// 创建map
func createMap() {
	// 方式1
	m1 := map[string]string{
		"name":   "xxx",
		"age":    "18",
		"gender": "male",
	}
	fmt.Println(m1)

	// 方式2
	m2 := make(map[string]int)
	fmt.Println(m2) // m2 == empty map

	// 方式3
	var m3 map[string]int
	fmt.Println(m3) // m3 == nil
}

func traverseMap() {
	// map的key是无序的
	m := map[string]string{
		"name":   "xxx",
		"age":    "18",
		"gender": "male",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

// map中获取值
func getValues() {
	m := map[string]string{
		"name":   "eric",
		"age":    "18",
		"gender": "male",
	}

	userName, ok := m["name"]
	fmt.Println("userName, ok: ", userName, ok)
	// 当获取map中不存在的key时，会获取到zero value(map中为空窜)
	causeName, ok := m["xxx"]
	fmt.Println("causeName, ok: ", causeName, ok)

	// 正确获取map中key对应value的方法
	if gender, ok := m["gender"]; ok {
		fmt.Println(gender)
	} else {
		fmt.Println("key does not exist")
	}
}

// 删除map中的元素
func deleteElement() {
	m := map[string]string{
		"name":   "eric",
		"age":    "18",
		"gender": "male",
	}

	name, ok := m["name"]
	fmt.Println("name, ok: ", name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println("name, ok: ", name, ok)
}

func main() {
	createMap()
	traverseMap()
	getValues()
	deleteElement()
}
