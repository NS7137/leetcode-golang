package populatingnextrightpointersineachnode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 主函数
func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTowNode(root.Left, root.Right)
	return root
}

// 辅助函数 连接跨父节点的两个相邻节点
func connectTowNode(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}

	// 将传入的两个节点连接
	node1.Next = node2

	// 连接相同父节点的两个子节点
	connectTowNode(node1.Left, node1.Right)
	connectTowNode(node2.Left, node2.Right)

	// 连接跨父节点的两个子节点
	connectTowNode(node1.Right, node2.Left)
}
