package nextgreaterelement2

func NextGreaterElements(nums []int) []int {
	n := len(nums)
	stack := []int{}
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}

	// 假装数组长度翻倍
	for i := 0; i < 2*n; i++ {
		// 索引求模
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			res[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return res
}
