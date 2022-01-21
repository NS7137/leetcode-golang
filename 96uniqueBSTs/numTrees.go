package uniquebsts

func numTrees(n int) int {
	return count(1, n)
}

// 计算闭区间[lo,hi]组成的bst个数
func count(lo, hi int) int {
	// base case
	if lo > hi {
		// 空节点也是一种情况
		return 1
	}
	res := 0
	for i := lo; i <= hi; i++ {
		// i的值作为根节点root
		left := count(lo, i-1)
		right := count(i+1, hi)
		// 左右子树的组合数乘积是bst的总数
		res += left * right
	}
	return res
}

// 递归
func numTreesRecursion(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	count := 0
	for i := 0; i <= n-1; i++ {
		count += numTreesRecursion(i) * numTreesRecursion(n-i-1)
	}
	return count
}

// dp
func numTreesDP(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}

// 记忆递归
func numTreesMemo(n int) int {
	memo := make([]int, n+1)
	return helper(n, &memo)
}

func helper(n int, memo *[]int) int {
	if n == 1 || n == 0 {
		return 1
	}
	if (*memo)[n] > 0 {
		return (*memo)[n]
	}
	count := 0
	for i := 0; i <= n-1; i++ {
		count += helper(i, memo) * helper(n-i-1, memo)
	}
	(*memo)[n] = count
	return count
}
