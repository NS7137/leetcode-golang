package utils

import "container/heap"

/*
切分，将图分为两个 不重叠且非空的节点集合
每次切分都能找到最小生成树的一条边，然后又可以进行新一轮切分，直到找到最小生成树的所有边为止
cut({A,B,C}) = cut({A,B}) + cut({C})
*/

type Prim struct {
	pq        hp       // 存储横切边的优先级队列
	inMST     []bool   // 记录哪些点已经成为最小生成数的一部分
	weightSum int      // 最小生成树的权重和
	graph     [][]edge // 邻接表表示图，graph[s]记录节点s所有相邻的边,三元组 from,to,weight 表示一条边
}

type edge struct {
	From, To, Weight int
}

type hp []edge

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].Weight < h[j].Weight }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(edge)) }
func (h *hp) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func NewPrim(graph [][]edge) Prim {
	p := Prim{}
	n := len(graph)
	p.graph = graph
	p.inMST = make([]bool, n)

	// 随便从一个节点开始切分，不妨从0开始
	p.inMST[0] = true
	p.cut(0)
	// 不断进行切分，向最小生成树中添加边
	for p.pq.Len() > 0 {
		edge := heap.Pop(&p.pq).(edge)
		to := edge.To
		weight := edge.Weight
		if p.inMST[to] {
			// 已经在最小生成树中，跳过，否则会产生环
			continue
		}
		// 将边加入最小生成树
		p.weightSum += weight
		p.inMST[to] = true
		// 节点to加入后，进行新一轮切分
		p.cut(to)
	}
	return p
}

// 将s的横切边加入优先级队列
func (p *Prim) cut(s int) {
	// 遍历s的邻边
	for _, edge := range p.graph[s] {
		to := edge.To
		if p.inMST[to] {
			// 相邻节点 to 已经在最小生成树中，跳过
			// 否则会产生环
			continue
		}
		// 加入横切边队列
		heap.Push(&p.pq, edge)
	}
}

// 判断最小生成树是否包含图中所有节点
func (p *Prim) AllConnected() bool {
	for i := 0; i < len(p.inMST); i++ {
		if !p.inMST[i] {
			return false
		}
	}
	return true
}
