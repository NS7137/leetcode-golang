package shipwithindays

func ShipWithinDays(weights []int, days int) int {
	left := 0
	right := 1
	// 计算承载上下限
	for _, w := range weights {
		left = max(left, w)
		right += w
	}

	for left < right {
		mid := left + (right-left)/2
		if f(weights, mid) <= days {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 定义 运载力为x时，需要f(x)天运完
func f(weights []int, x int) int {
	days := 0
	for i := 0; i < len(weights); {
		// 尽可能多装
		cap := x
		for i < len(weights) {
			if cap < weights[i] {
				break
			} else {
				cap -= weights[i]
			}
			i++
		}
		days++
	}
	return days
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
