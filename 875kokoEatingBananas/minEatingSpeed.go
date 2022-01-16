package kokoeatingbananas

func MinEatingSpeed(piles []int, h int) int {
	left := 1
	right := 1000000000 + 1
	for left < right {
		mid := left + (right-left)/2
		if f(piles, mid) <= h {
			// 吃完小于h说明快，减上限
			right = mid
		} else {
			// 吃完大于h说明慢, 加下限
			left = mid + 1
		}
	}
	return left

}

// 定义速度为x时，需要f(x)小时 吃完所有
func f(piles []int, x int) int {
	hours := 0
	for i := 0; i < len(piles); i++ {
		hours += piles[i] / x
		if piles[i]%x != 0 {
			hours++
		}
	}
	return hours
}
