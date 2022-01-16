package rangesumquery2d

type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}

	m := len(matrix)
	n := len(matrix[0])

	sum := make([][]int, m+1)
	for i := 0; i < len(sum); i++ {
		sum[i] = make([]int, n+1)
	}

	sum[1][1] = matrix[0][0]
	for i := 2; i < len(sum[0]); i++ {
		sum[1][i] = sum[1][i-1] + matrix[0][i-1]
	}

	for i := 2; i < len(sum); i++ {
		sum[i][1] = sum[i-1][1] + matrix[i-1][0]
	}

	for i := 2; i < len(sum); i++ {
		for j := 2; j < len(sum[0]); j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	return NumMatrix{sum: sum}

}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return nm.sum[row2+1][col2+1] - nm.sum[row2+1][col1] - nm.sum[row1][col2+1] + nm.sum[row1][col1]
}
