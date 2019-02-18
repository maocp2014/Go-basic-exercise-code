package tree

import "fmt"

func (node *Node) Traverse() {
	// 匿名函数作为参数传递
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

// 函数作为参数传递
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
    // 中序遍历
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

// func (node *Node) TraverseWithChannel() chan *Node {
// 	out := make(chan *Node)
// 	go func() {
// 		node.TraverseFunc(func(node *Node) {
// 			out <- node
// 		})
// 		close(out)
// 	}()
// 	return out
// }
