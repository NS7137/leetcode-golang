package subarraysumequalsk

func SubarraySumByMap(nums []int, k int) int {
	dict := make(map[int]int)
	dict[0] = 1
	sum := 0
	res := 0
	for _, v := range nums {
		sum += v
		res += dict[sum-k]
		dict[sum]++
	}
	return res
}

func SubarraySumByPreSum(nums []int, k int) int {
	var n = len(nums)
	prefix := make([]int, n+1)
	for i, v := range nums {
		prefix[i+1] = prefix[i] + v
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if prefix[j]-prefix[i] == k {
				res++
			}
		}
	}
	return res
}
