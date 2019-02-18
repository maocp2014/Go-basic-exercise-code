package main

import (
	"fmt"
	"go_pratice_code/learn_ccmouse_code/functional/tree"
)

// 扩展已有类型方式：组合
// 实现后序遍历
type myTreeNode struct {
	node *tree.Node
}

// 实现后序遍历方法
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
    // 后序遍历 
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	fmt.Print("In-order traversal: ")
	root.Traverse()

	fmt.Print("My own post-order traversal: ")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

	nodeCount := 0
    // 匿名函数
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	// c := root.TraverseWithChannel()
	// maxNodeValue := 0
	// for node := range c {
	// 	if node.Value > maxNodeValue {
	// 		maxNodeValue = node.Value
	// 	}
	// }
	// fmt.Println("Max node value:", maxNodeValue)
}
