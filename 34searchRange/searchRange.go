package searchrange

func SearchRange(nums []int, target int) []int {
	return []int{leftBound(nums, target), rightBound(nums, target)}
}

func leftBound(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	// 左右闭区间
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			// 收缩右边界
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	// 检查left越界
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func rightBound(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			// 收缩左边界
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	// 检查right越界
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
