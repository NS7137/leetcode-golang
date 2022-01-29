package numberofconnectedcomponentsinanundirectedgraph

type UF struct {
	count  int   // 记录连通分量
	parent []int // 节点x的根节点是parent[x]
}

func Constructor(n int) UF {
	// 父节点指针初始指向自己
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	// 一开始互不连通
	uf := UF{
		count:  n,
		parent: parent,
	}
	return uf
}

// 连接 p 和 q
// 一个节点的根节点连接到另一个节点的根节点上
func (uf *UF) union(p, q int) {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	if rootP == rootQ {
		return
	}
	// 合并
	uf.parent[rootP] = rootQ
	// 两个分量合二为一
	uf.count--
}

// 返回节点x的根节点
func (uf *UF) find(x int) int {
	// 父节点指针没有指向自己 就往上找
	for uf.parent[x] != x {
		x = uf.parent[x]
	}
	return x
}

// 判断 p 和 q 是否连通
func (uf *UF) connected(p, q int) bool {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	return rootP == rootQ
}

// 平衡性优化，在union过程，小树接到大树下面，避免头重脚轻
type UFBalance struct {
	count  int   // 记录连通分量
	parent []int // 节点x的根节点是parent[x]
	size   []int // 记录树的重量
}

func NewUF(n int) UFBalance {
	parent := make([]int, n)
	size := make([]int, n)
	// 重量初始化为1，父节点指向自己
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	ufb := UFBalance{
		count:  n,
		parent: parent,
		size:   size,
	}
	return ufb
}

func (ufb *UFBalance) union(p, q int) {
	rootP := ufb.find(p)
	rootQ := ufb.find(q)
	if rootP == rootQ {
		return
	}
	// 小树连接到大树下
	if ufb.size[rootP] > ufb.size[rootQ] {
		ufb.parent[rootQ] = rootP
		ufb.size[rootP] += ufb.size[rootQ]
	} else {
		ufb.parent[rootP] = rootQ
		ufb.size[rootQ] += ufb.size[rootP]
	}
	ufb.count--
}

// 路径压缩
func (ufb *UFBalance) find(x int) int {
	for ufb.parent[x] != x {
		ufb.parent[x] = ufb.parent[ufb.parent[x]]
		x = ufb.parent[x]
	}
	return x
}

func (ufb *UFBalance) connected(p, q int) bool {
	rootP := ufb.find(p)
	rootQ := ufb.find(q)
	return rootP == rootQ
}
