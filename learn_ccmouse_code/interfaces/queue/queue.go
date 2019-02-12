package queue

// 扩展已有类型 方法2： 类型重命名
// 支持任意类型(interface{}表示任意类型)
type Queue []interface{}

// 在函数里面限制int类型
// push
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v.(int)) // 强制类型转换
}

// pop
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	(*q) = (*q)[1:]
	return head.(int) // 强制类型转换
}

// 队列为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// 注意
// 上述函数都是指针接收者，会改变原Queue的值，每次的Queue都不同，这点与其它面向对象编程有所区别
