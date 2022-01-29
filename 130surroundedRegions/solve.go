package surroundedregions

/*
将二维映射到一维 (x,y) -> x*n + y (m行，n列)
x * n 当前数之前有几个数
y 当前数在第几个位置

数组范围
[0...m*n-1]
设置虚拟节点dummy索引m*n
把不需要替换的O与dummy连通
*/

func Solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	m := len(board)
	n := len(board[0])
	// 给dummy留一个位置
	uf := NewUF(m*n + 1)
	dummy := m * n
	// 首列与末列O与dummy连通
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			uf.union(i*n, dummy)
		}
		if board[i][n-1] == 'O' {
			uf.union(i*n+n-1, dummy)
		}
	}
	// 首行与末行O与dummy连通
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			uf.union(j, dummy)
		}
		if board[m-1][j] == 'O' {
			uf.union((m-1)*n+j, dummy)
		}
	}

	// 方向数组 上下左右
	d := [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}
	// 内O与上下左右O连通
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' {
				for k := 0; k < 4; k++ {
					x := i + d[k][0]
					y := j + d[k][1]
					if board[x][y] == 'O' {
						uf.union(x*n+y, i*n+j)
					}
				}
			}
		}
	}

	// 替换所有不和dummy连通的O
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if !uf.connected(dummy, i*n+j) {
				board[i][j] = 'X'
			}
		}
	}

}

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

var m, n int

// dfs
func SolveByDFS(board [][]byte) {
	// 最小3x3 否则不可能出现被包围的O
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n = len(board), len(board[0])
	for i := 0; i < m; i++ {
		dfs(board, i, 0)   //第一列
		dfs(board, i, n-1) // 最后一列
	}
	for i := 1; i < n-1; i++ {
		dfs(board, 0, i)   // 第一行
		dfs(board, m-1, i) // 最后一行
	}

	//其余中间的查找替换
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	// 是否越界
	if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
		return
	}
	// 从边界开始，与边界相连的O改#
	board[x][y] = '#'
	// 上下左右
	dfs(board, x+1, y)
	dfs(board, x-1, y)
	dfs(board, x, y+1)
	dfs(board, x, y-1)
}
