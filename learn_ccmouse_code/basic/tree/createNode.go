package tree

// struct中没有构造函数
// 自定义工厂函数(构造函数)生成treenode
func nodeFactoryFunc(value int) *Node {
	// &treeNode{value: value}是局部变量地址，go中可以返回局部变量供外面使用
	return &Node{value: value}
}

// struct赋值方法
func CreateNode() *Node {
	// var root treeNode
	// fmt.Println(root) // {0 <nil> <nil>}
	root := Node{value: 3}
	root.Left = &Node{} // root.left是1个指针，这里取地址赋值
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node) // new是内建函数，这里新建了一个空节点
	// &{0 <nil> <nil>} {3 0xc000048400 0xc000048420} &{5 0xc000048440 <nil>} &{0 <nil> <nil>}
	// fmt.Println(root.left, root, root.right, root.right.left)

	// struct切片
	// nodes := []treeNode{
	// 	{value: 3},
	// 	{},
	// 	{5, nil, &root}, // ","不能少
	// }
	// fmt.Printf("root address is : %p\n", &root)
	// fmt.Println(nodes)
	// 调用工厂函数生成节点
	root.Left.Right = nodeFactoryFunc(2)
	return &root
}
