package findthecelebrity

/*
所谓名人定义
所有其他人都认识名人
名人不认识任何其他人
那么 它保证了人群中最多有一个名人

两人关系 4种
你认识我我不认识你，我认识你你不认识我，互相认识，互相不认识

逐一排除掉一个
cand 认识 other 所以cand不是，排除掉一个
other 认识 cand 所以other不是，排除掉一个
互相认识，肯定都不是名人,随便排除掉一个
互相不认识，肯定都不是名人,随便排除掉一个

不断从候选人中选两个出来，然后排除掉一个
*/

// 总人数 n 通过 knows 来查询人和人的关系
func FindCelebrity(n int) int {
	// 暴力解法
	// 遍历每个候选人
	for cand := 0; cand < n; cand++ {
		var other int
		for other = 0; other < n; other++ {
			if cand == other {
				continue
			}
			// 保证其他人都认识cand，且cand不认识其他人
			// 否则cand就不能是名人
			if knows(cand, other) || !knows(other, cand) {
				break
			}
		}
		if other == n {
			// 找到
			return cand
		}
	}
	// 没找到
	return -1
}

// 优化
func FindCelebrity2(n int) int {
	if n == 1 {
		return 0
	}
	// 所有候选入队列
	q := make([]int, n)
	for i := range q {
		q = append(q, i)
	}
	// 一直排除，直到剩下一个候选人停止循环
	for len(q) >= 2 {
		// 每次取2个，排除一个
		cand := q[0]
		other := q[1]
		q = q[2:]
		if knows(cand, other) || !knows(other, cand) {
			// cand 不是，排除
			q = append(q, other)
		} else {
			// other 不是，排除
			q = append(q, cand)
		}
	}
	// 排除得只剩一个候选，判断他是否是名人
	cand := q[0]
	for other := 0; other < n; other++ {
		if other == cand {
			continue
		}
		// 保证其他人都认识cand，且cand不认识其他人
		if !knows(other, cand) || knows(cand, other) {
			return -1
		}
	}
	return cand
}

// 最终
func FindCelebrity3(n int) int {
	// 假设cand是名人
	cand := 0
	for other := 1; other < n; other++ {
		if !knows(other, cand) || knows(cand, other) {
			// cand 不是名人，排除
			// 假设 other 是名人
			cand = other
		}
		// else
		// other 不是名人 排除
		// 什么都不做，继续假设cand是名人
	}

	// cand是排除的最后结果，但不能保证一定是
	for other := 0; other < n; other++ {
		if cand == other {
			continue
		}
		// 保证其他人认识cand，且cand不认识其他人
		if !knows(other, cand) || knows(cand, other) {
			return -1
		}
	}
	return cand
}

func knows(i, j int) bool {
	graph := [][]int{{1, 0, 1}, {1, 1, 0}, {0, 1, 1}} //假设
	// i认识j 认识返回true
	return graph[i][j] == 1
}
