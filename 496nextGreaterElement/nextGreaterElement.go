package nextgreaterelement

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	n1 := len(nums1)
	n2 := len(nums2)
	stack := []int{}
	dict := make(map[int]int)

	for i := 0; i < n2; i++ {
		// 从nums2中取元素压栈，与后续元素比较，前大于后就跳过，后大于前就map记录，并出栈
		for len(stack) != 0 {
			a := stack[len(stack)-1]
			if nums2[i] <= a {
				break
			} else {
				dict[a] = nums2[i]
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, nums2[i])
	}

	res := make([]int, n1)
	for _, key := range nums1 {
		// 在nums1中元素为key 是否存在，存在赋值，不存在 -1
		if v, ok := dict[key]; ok {
			res = append(res, v)
		} else {
			res = append(res, -1)
		}
	}

	return res
}
