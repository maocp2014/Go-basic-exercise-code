package main

import (
	"fmt"
	"sync"
)

/*
日常应用中，大多数并发数据结构都是读多写少的，对于读多写少的场合，可以将互斥锁换成读写锁，可以有效提升性能。
sync 包也提供了读写锁对象 RWMutex，不同于互斥锁只有两个常用方法 Lock() 和 Unlock()，读写锁提供了四个常用方法，
分别是写加锁 Lock()、写释放锁 Unlock()、读加锁 RLock() 和读释放锁 RUnlock()。写锁是拍他锁，加写锁时会阻塞其它协程再加读锁和写锁，
读锁是共享锁，加读锁还可以允许其它协程再加读锁，但是会阻塞加写锁。

读写锁在写并发高的情况下性能退化为普通的互斥锁

下面我们将上面代码中 SafeDict 的互斥锁改造成读写锁。
*/
type SafeDict struct {
	data          map[string]int
	*sync.RWMutex // 读写锁
}

func NewSafeDict(data map[string]int) *SafeDict {
	return &SafeDict{data, &sync.RWMutex{}}
}

func (d *SafeDict) Len() int {
	d.RLock()
	defer d.RUnlock()
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
	d.RLock()
	defer d.RUnlock()
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
