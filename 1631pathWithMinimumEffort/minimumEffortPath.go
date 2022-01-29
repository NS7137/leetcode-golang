package pathwithminimumeffort

import (
	"container/heap"
	"math"
)

// 左上到右下角 最小体力消耗
func MinimumEffortPath(heights [][]int) int {
	return dijkstra(heights)
}

type pair struct {
	x, y int
}

// 方向数组
var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 返回坐标 (x,y) 的上下左右相邻坐标
func adj(matrix [][]int, x, y int) []pair {

	m, n := len(matrix), len(matrix[0])
	// 存储相邻节点
	neighbors := []pair{}
	for _, dir := range dirs {
		nx := x + dir.x
		ny := y + dir.y
		if nx >= m || nx < 0 || ny >= n || ny < 0 {
			//越界
			continue
		}
		neighbors = append(neighbors, pair{x: nx, y: ny})
	}
	return neighbors
}

func dijkstra(heights [][]int) int {
	m := len(heights)
	n := len(heights[0])
	// 定义从(0,0)到(i,j)的最小体力消耗是effortTo[i][j]
	effortTo := make([][]int, m)
	for i := range effortTo {
		effortTo[i] = make([]int, n)
		for j := range effortTo[i] {
			effortTo[i][j] = math.MaxInt64
		}
	}

	// base case 起点到起点的最小消耗就是 0
	effortTo[0][0] = 0

	// 优先级队列 effortFromStart 小的排前面
	pq := &hp{}
	// 从起点 (0,0) 开始 BFS
	heap.Push(pq, NewState(0, 0, 0))

	for pq.Len() > 0 {
		curState := heap.Pop(pq).(State)
		curX := curState.x
		curY := curState.y
		curEffortFromStart := curState.effortFromStart

		// 到达终点提前结束
		if curX == m-1 && curY == n-1 {
			return curEffortFromStart
		}

		if curEffortFromStart > effortTo[curX][curY] {
			continue
		}
		// 将 (curX,curY) 相邻坐标 入队列
		neighbors := adj(heights, curX, curY)
		for _, neighbor := range neighbors {
			nextX := neighbor.x
			nextY := neighbor.y
			// 计算 (curX,curY) 到 (nextX,nextY) 的消耗
			effortToNextNode := max(
				effortTo[curX][curY],
				abs(heights[curX][curY]-heights[nextX][nextY]),
			)

			// 更新 dp table
			if effortTo[nextX][nextY] > effortToNextNode {
				effortTo[nextX][nextY] = effortToNextNode
				heap.Push(pq, NewState(nextX, nextY, effortToNextNode))
			}
		}
	}
	return -1
}

type State struct {
	x, y            int // 矩阵中的一个位置
	effortFromStart int // 从起点(0,0) 到当前位置的最小体力消耗
}

func NewState(x, y int, effortFromStart int) State {
	return State{x: x, y: y, effortFromStart: effortFromStart}
}

type hp []State

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].effortFromStart < h[j].effortFromStart }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(State)) }
func (h *hp) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
