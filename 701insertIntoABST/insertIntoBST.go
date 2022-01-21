package insertintoabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func InsertIntoBST(root *TreeNode, val int) *TreeNode {
	// 找到空位置插入新节点
	if root == nil {
		return &TreeNode{Val: val}
	}
	// BST中一般不会插入已存在的元素

	if val < root.Val {
		root.Left = InsertIntoBST(root.Left, val)
	}

	if val > root.Val {
		root.Right = InsertIntoBST(root.Right, val)
	}
	return root
}

// 迭代
func InsertIntoBSTIter(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	cur := root
	newNode := &TreeNode{Val: val}

	for cur != nil {
		if val < cur.Val {
			if cur.Left == nil {
				cur.Left = newNode
				break
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = newNode
				break
			}
			cur = cur.Right
		}
	}
	return root
}
