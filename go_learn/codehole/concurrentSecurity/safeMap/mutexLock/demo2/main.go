package main

import (
	"fmt"
	"sync"
)

type SafeDict struct {
	data map[string]int
	// 匿名锁字段，外部结构体可以自动继承匿名内部结构体的所有方法。如果将上面的 SafeDict 结构体进行改造，将锁字段匿名，就可以稍微简化一下代码
	*sync.Mutex
}

func NewSafeDict(data map[string]int) *SafeDict {
	return &SafeDict{data, &sync.Mutex{}}
}

func (d *SafeDict) Len() int {
	d.Lock()
	defer d.Unlock()
	return len(d.data)
}

func (d *SafeDict) Put(key string, value int) (int, bool) {
	d.Lock()
	defer d.Unlock()
	old_value, ok := d.data[key]
	d.data[key] = value
	return old_value, ok
}

func (d *SafeDict) Get(key string) (int, bool) {
	d.Lock()
	defer d.Unlock()
	old_value, ok := d.data[key]
	return old_value, ok
}

func (d *SafeDict) Delete(key string) (int, bool) {
	d.Lock()
	defer d.Unlock()
	old_value, ok := d.data[key]
	if ok {
		delete(d.data, key)
	}
	return old_value, ok
}

func write(d *SafeDict) {
	d.Put("banana", 5)
}

func read(d *SafeDict) {
	fmt.Println(d.Get("banana"))
}

func main() {
	d := NewSafeDict(map[string]int{
		"apple": 2,
		"pear":  3,
	})
	go read(d)
	write(d)
}
