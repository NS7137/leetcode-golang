package removeduplicatesfromsortedarray

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow, fast := 0, 0
	// 快慢指针，相等就fast前移，不等就赋值slow前移
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			// 维护nums[0...slow]无重复
			nums[slow] = nums[fast]
		}
		fast++
	}

	// 数组长为索引加1
	return slow + 1

}
