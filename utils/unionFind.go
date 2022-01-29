package utils

type UF struct {
	count  int   // 记录连通分量
	parent []int // 节点x的根节点是parent[x]
	size   []int // 记录树的重量
}

func NewUF(n int) UF {
	parent := make([]int, n)
	size := make([]int, n)
	// 重量初始化为1，父节点指向自己
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	uf := UF{
		count:  n,
		parent: parent,
		size:   size,
	}
	return uf
}

func (uf *UF) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP == rootQ {
		return
	}
	// 小树连接到大树下
	if uf.size[rootP] > uf.size[rootQ] {
		uf.parent[rootQ] = rootP
		uf.size[rootP] += uf.size[rootQ]
	} else {
		uf.parent[rootP] = rootQ
		uf.size[rootQ] += uf.size[rootP]
	}
	uf.count--
}

// 路径压缩
func (uf *UF) Find(x int) int {
	for uf.parent[x] != x {
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

func (uf *UF) Connected(p, q int) bool {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	return rootP == rootQ
}

func (uf *UF) Count() int {
	return uf.count
}
