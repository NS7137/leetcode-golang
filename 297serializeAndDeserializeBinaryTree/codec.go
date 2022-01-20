package serializeanddeserializebinarytree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	nq []*TreeNode
	sq []string
}

func Constructor() Codec {
	return Codec{
		nq: []*TreeNode{},
		sq: []string{},
	}
}

// Serializes a tree to a single string.
func (c *Codec) Serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	c.preBuildString(root, sb)
	return strings.Trim(sb.String(), SEP)
}

// Deserializes your encoded data to tree.
func (c *Codec) Deserialize(sq string) *TreeNode {
	c.sq = strings.Split(sq, SEP)
	return c.preBuildTree()
}

const SEP string = ","
const NULL string = "#"

func (c *Codec) preBuildString(root *TreeNode, sb *strings.Builder) {
	if root == nil {
		sb.WriteString(NULL)
		sb.WriteString(SEP)
		return
	}
	sb.WriteString(strconv.Itoa(root.Val))
	sb.WriteString(SEP)
	c.preBuildString(root.Left, sb)
	c.preBuildString(root.Right, sb)
}

func (c *Codec) preBuildTree() *TreeNode {
	if c.sq[0] == NULL {
		c.sq = c.sq[1:]
		return nil
	}
	val, _ := strconv.Atoi(c.sq[0])
	c.sq = c.sq[1:]
	t := &TreeNode{Val: val, Left: nil, Right: nil}
	t.Left = c.preBuildTree()
	t.Right = c.preBuildTree()
	return t
}

func (c *Codec) postBuildString(root *TreeNode, sb *strings.Builder) {
	if root == nil {
		sb.WriteString(NULL)
		sb.WriteString(SEP)
		return
	}
	c.postBuildString(root.Left, sb)
	c.postBuildString(root.Right, sb)
	sb.WriteString(strconv.Itoa(root.Val))
	sb.WriteString(SEP)
}

func (c *Codec) postBuildTree() *TreeNode {
	if len(c.sq) == 0 {
		return nil
	}
	last := c.sq[len(c.sq)-1]
	if last == NULL {
		c.sq = c.sq[0 : len(c.sq)-1]
		return nil
	}
	val, _ := strconv.Atoi(last)
	c.sq = c.sq[0 : len(c.sq)-1]
	t := &TreeNode{Val: val, Left: nil, Right: nil}
	// 先从右树开始构造 再左树
	t.Right = c.postBuildTree()
	t.Left = c.postBuildTree()
	return t
}

// 只能序列化，无法反序列，因为无法得到root节点的索引
func (c *Codec) inBuildString(root *TreeNode, sb *strings.Builder) {
	if root == nil {
		sb.WriteString(NULL)
		sb.WriteString(SEP)
		return
	}
	c.postBuildString(root.Left, sb)
	sb.WriteString(strconv.Itoa(root.Val))
	sb.WriteString(SEP)
	c.postBuildString(root.Right, sb)
}

func (c *Codec) levelBuildString(root *TreeNode, sb *strings.Builder) {
	if root == nil {
		sb.WriteString(NULL)
		sb.WriteString(SEP)
		return
	}
	c.nq = append(c.nq, root)

	for len(c.nq) != 0 {
		cur := c.nq[0]
		// pop
		c.nq = c.nq[1:]

		if cur == nil {
			sb.WriteString(NULL)
			sb.WriteString(SEP)
			continue
		}
		sb.WriteString(strconv.Itoa(cur.Val))
		sb.WriteString(SEP)

		c.nq = append(c.nq, cur.Left)
		c.nq = append(c.nq, cur.Right)
	}
}

func (c *Codec) levelBuildTree() *TreeNode {
	if len(c.sq) == 0 {
		return nil
	}
	// 第一个元素就是root
	rootVal, _ := strconv.Atoi(c.sq[0])
	root := &TreeNode{Val: rootVal, Left: nil, Right: nil}

	// 记录父节点 将root入队列
	c.nq = append(c.nq, root)

	for i := 1; i < len(c.sq); {
		// 队列中存的都是父节点
		parent := c.nq[0]
		c.nq = c.nq[1:]
		// 父节点对应的左子节点的值
		left := c.sq[i]
		i++
		leftVal, _ := strconv.Atoi(left)
		if left != NULL {
			parent.Left = &TreeNode{Val: leftVal, Left: nil, Right: nil}
			c.nq = append(c.nq, parent.Left)
		}
		// 父节点独赢右侧子节点的值
		right := c.sq[i]
		i++
		rightVal, _ := strconv.Atoi(right)
		if right != NULL {
			parent.Right = &TreeNode{Val: rightVal, Left: nil, Right: nil}
			c.nq = append(c.nq, parent.Right)
		}
	}
	return root
}
