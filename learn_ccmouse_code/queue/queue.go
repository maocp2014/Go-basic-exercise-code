package queue

// 扩展已有类型 方法2： 类型重命名
type Queue []int

// push
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// pop
func (q *Queue) Pop() int {
	head := (*q)[0]
	(*q) = (*q)[1:]
	return head
}

// 队列为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// 注意
// 上述函数都是指针接收者，会改变原Queue的值，每次的Queue都不同，这点与其它面向对象编程有所区别
