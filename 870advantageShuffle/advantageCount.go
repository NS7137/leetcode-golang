package advantageshuffle

import "sort"

func AdvantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1) - 1
	// copy nums2
	b := make([]int, len(nums2))
	copy(b, nums2)
	a := nums1
	sort.Ints(a)
	sort.Ints(b)

	dict, index := make(map[int][]int), n
	// 从大往小遍历 能胜过的都记录到map
	for i := n; i >= 0; i-- {
		if b[i] < a[index] {
			dict[b[i]] = append(dict[b[i]], a[index])
			index--
		}
	}
	res := []int{}
	for _, b := range nums2 {
		if arr, ok := dict[b]; ok && len(arr) >= 1 {
			res = append(res, arr[len(arr)-1])
			dict[b] = arr[:len(arr)-1]
		} else {
			// 剩下的都是小于的随便给
			res = append(res, a[index])
			index--
		}
	}

	return res

}
