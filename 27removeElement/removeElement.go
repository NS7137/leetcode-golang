package removeelement

func RemoveElement(nums []int, val int) int {

	numsLen := len(nums)
	// 还是快慢指针，快指针跳过需要去除的元素
	slow, fast := 0, 0
	for fast < numsLen {
		if nums[fast] != val {
			// 先赋值再slow++,保证返回不包含val元素
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
