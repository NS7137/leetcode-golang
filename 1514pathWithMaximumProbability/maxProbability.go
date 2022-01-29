package pathwithmaximumprobability

import "container/heap"

// 输入无向图，边上权重代表概率，返回从start 到 end 最大的概率
func MaxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := buildGraph(n, edges, succProb)
	return dijkstra(start, end, graph)
}

type State struct {
	id            int     // 图节点的id
	probFromStart float64 // 从start节点到当前节点的概率
}

func newState(id int, probFromStart float64) State {
	return State{id: id, probFromStart: probFromStart}
}

// 构造邻接表
func buildGraph(n int, edges [][]int, succProb []float64) [][]State {
	graph := make([][]State, n)
	for i := range graph {
		graph[i] = []State{}
	}

	for i, edge := range edges {
		from := edge[0]
		to := edge[1]
		weight := succProb[i]
		// 无向图就是双向图
		graph[from] = append(graph[from], newState(to, weight))
		graph[to] = append(graph[to], newState(from, weight))
	}
	return graph
}

func dijkstra(start, end int, graph [][]State) float64 {
	// 定义probTo[i] 的值就是节点start到节点i的最大概率
	probTo := make([]float64, len(graph))
	// dp table初始化为取不到的最小值
	for i := range probTo {
		probTo[i] = -1
	}
	// base case start到start的概率就是1
	probTo[start] = 1
	// 优先级队列 probFromStart 大的排在前面
	pq := &hp{}
	heap.Push(pq, newState(start, 1))

	for pq.Len() > 0 {
		curState := heap.Pop(pq).(State)
		curNodeID := curState.id
		curProbFromStart := curState.probFromStart

		// 遇到终点提前返回
		if curNodeID == end {
			return curProbFromStart
		}

		if curProbFromStart < probTo[curNodeID] {
			// 已经有一条概率更大的路径到达 当前节点
			continue
		}
		// 将当前节点的相邻节点入队列
		for _, neighbor := range graph[curNodeID] {
			nextNodeID := neighbor.id
			// 从当前curNode到nextNode的概率是否会更大
			probToNextNode := probTo[curNodeID] * neighbor.probFromStart
			if probTo[nextNodeID] < probToNextNode {
				probTo[nextNodeID] = probToNextNode
				heap.Push(pq, newState(nextNodeID, probToNextNode))
			}
		}
	}
	// 如果到达这里，说明从start开始无法到达end，返回 0
	return 0
}

type hp []State

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].probFromStart > h[j].probFromStart }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(State)) }
func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
