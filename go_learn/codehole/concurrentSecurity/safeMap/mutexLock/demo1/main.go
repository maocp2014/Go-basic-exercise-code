package main

import (
	"fmt"
	"sync"
	"time"
)

// 线程安全map
type SafeDict struct {
	data map[string]int
	// sync.Mutex 是一个结构体对象，这个对象在使用的过程中要避免被复制 —— 浅拷贝。
	// 复制将会导致锁被「分裂」了（操作的不是同一个锁），也就起不到保护的作用。所以在平时的使用中要尽量使用它的指针类型
	mutex *sync.Mutex // 互斥锁
}

// 生成线程安全的map
func NewSafeDict(data map[string]int) *SafeDict {
	return &SafeDict{
		data:  data,
		mutex: &sync.Mutex{},
	}
}

func (d *SafeDict) Len() int {
	d.mutex.Lock() // 加锁
	// defer 语句总是要推迟到函数尾部运行，所以如果函数逻辑运行时间比较长，这会导致锁持有的时间较长，
	// 这时使用 defer 语句来释放锁未必是一个好主意
	defer d.mutex.Unlock() // 释放锁
	return len(d.data)
}

func (d *SafeDict) Put(key string, value int) (int, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	old_value, ok := d.data[key]
	d.data[key] = value
	return old_value, ok
}

func (d *SafeDict) Get(key string) (int, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	old_value, ok := d.data[key]
	return old_value, ok
}

func (d *SafeDict) Delete(key string) (int, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
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
	fmt.Println(d.Delete("apple"))
	time.Sleep(time.Second)
}