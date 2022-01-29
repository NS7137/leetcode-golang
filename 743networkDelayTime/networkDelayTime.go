package networkdelaytime

import (
	"container/heap"
	"math"
)

/*
1...n个节点， 从k发出信号，多久所有节点都收到信号
Dijkstra 条件，有向加权，无负权重
*/

// 从k出发到其他节点的最短路径
func NetworkDelayTime(times [][]int, n int, k int) int {
	graph := buildGraph(times, n)
	distTo := dijkstra(k, graph)
	// 找到最长的那条最短路径
	res := 0
	for i := 1; i < len(distTo); i++ {
		if distTo[i] == math.MaxInt64 {
			// 有节点不可达，返回 -1
			return -1
		}
		res = max(res, distTo[i])
	}
	return res
}

type State struct {
	id            int // 图节点的id
	distFromStart int // 从start节点到当前节点的距离
}

func NewState(id, distFromStart int) State {
	return State{id: id, distFromStart: distFromStart}
}

func buildGraph(times [][]int, n int) [][]State {
	// 编号 1...n
	graph := make([][]State, n+1)
	for i := 1; i <= n; i++ {
		graph[i] = []State{}
	}

	for _, edge := range times {
		from := edge[0]
		to := edge[1]
		weight := edge[2]
		graph[from] = append(graph[from], NewState(to, weight))
	}
	return graph
}

// 输入start，计算从start到其他节点的最短距离
func dijkstra(start int, graph [][]State) []int {
	// 定义 distTo[i] 的值就是起点start到节点i的最短路径权重
	distTo := make([]int, len(graph))
	// 初始化 无穷大
	for i := range distTo {
		distTo[i] = math.MaxInt64
	}

	// base case, start 到 start 最短距离 0
	distTo[start] = 0

	// 优先级队列 distFromStart 小的排前面
	pq := &QueueState{}
	// 从起点start 开始 BFS
	heap.Push(pq, NewState(start, 0))

	for pq.Len() != 0 {
		curState := heap.Pop(pq).(State)
		curNodeID := curState.id
		curDistFromStart := curState.distFromStart

		if curDistFromStart > distTo[curNodeID] {
			continue
		}
		// 将curNode的相邻节点入队列
		for _, neighbor := range graph[curNodeID] {
			nextNodeID := neighbor.id
			distToNextNode := distTo[curNodeID] + neighbor.distFromStart

			// 更新 dp table
			if distTo[nextNodeID] > distToNextNode {
				distTo[nextNodeID] = distToNextNode
				heap.Push(pq, NewState(nextNodeID, distToNextNode))
			}
		}
	}
	return distTo
}

type QueueState []State

func (pq *QueueState) Len() int {
	return len(*pq)
}

func (pq *QueueState) Less(i, j int) bool {
	return (*pq)[i].distFromStart < (*pq)[j].distFromStart
}

func (pq *QueueState) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *QueueState) Push(x interface{}) {
	*pq = append(*pq, x.(State))
}

func (pq *QueueState) Pop() interface{} {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
