package tree

import "fmt"

// 结构体：节点定义
type Node struct {
	value       int
	Left, Right *Node
}

// 结构体定义方法
// 参数传递方式：传值
// 值接收者
func (node Node) Print() {
	fmt.Println(node.value)
}

// 指针接收者
func (node *Node) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.value = value // nil不能取值，所有前面需要return
}

// 接收者允许为空指针nil
// nil可以调用方法，因此在遍历时不需要判断节点是否为nil
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	// 中序遍历
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}