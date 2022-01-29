package mincosttoconnectallpoints

import "sort"

func MinCostConnectPoints(points [][]int) int {
	n := len(points)
	// 生成所有边及权重
	edges := [][]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			xi, yi := points[i][0], points[i][1]
			xj, yj := points[j][0], points[j][1]
			// 用坐标点在points中的索引表示坐标点
			weight := abs(xi-xj) + abs(yi-yj)
			edges = append(edges, []int{i, j, weight})
		}
	}
	sort.Sort(ByWeight(edges))
	// 执行Kruskal
	mst := 0
	uf := NewUF(n)
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		weight := edge[2]
		// 若产生环，则不能加入mst
		if uf.Connected(u, v) {
			continue
		}
		mst += weight
		uf.Union(u, v)
	}
	if uf.Count() == 1 {
		return mst
	} else {
		return -1
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 按权重升序排序
type ByWeight [][]int

func (a ByWeight) Len() int {
	return len(a)
}
func (a ByWeight) Less(i, j int) bool {
	return a[i][2] < a[j][2]
}
func (a ByWeight) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// 并查集
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
