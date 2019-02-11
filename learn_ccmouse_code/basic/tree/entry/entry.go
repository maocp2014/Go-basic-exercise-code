package main

import "go_pratice_code/learn_ccmouse_code/basic/tree"

// 扩展已有类型 方法1: 组合
type myTreeNode struct {
	node *tree.Node
}

// 后序遍历
func (myNode *myTreeNode) postOrder(){
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	root := tree.CreateNode()
	// fmt.Println(root.left.right, root.left, root, root.right, root.right.left)
	// root.print()

	// root.setValue(10)
	// root.print()

	// var pRoot *treeNode
	// pRoot.setValue(200)  // nil指针也可以调用方法
	// pRoot = root
	// pRoot.setValue(300)
	// pRoot.print()

	// 中序遍历
	root.Traverse()
}
