package movezeroes

func MoveZeroes(nums []int) {

	// 返回去0的长度
	p := removeElement(nums, 0)

	// p之后的所有元素赋值0
	for ; p < len(nums); p++ {
		nums[p] = 0
	}

}

// 复用27移除元素的解法

func removeElement(nums []int, val int) int {

	numsLen := len(nums)
	slow, fast := 0, 0
	for fast < numsLen {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func MoveZeroesBySwap(nums []int) {
	n := 0
	for i := 0; i < len(nums); i++ {
		// 为0就跳过，不为0，就与下标n值互换，n++
		if nums[i] != 0 {
			nums[n], nums[i] = nums[i], nums[n]
			n++
		}
	}
}
