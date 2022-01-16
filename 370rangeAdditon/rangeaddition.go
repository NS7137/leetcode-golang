package rangeadditon

import "sync"

type NumArray struct {
	diff []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{}
	}
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return NumArray{diff: diff}
}

func (da *NumArray) increment(i int, j int, val int) {
	da.diff[i] += val
	if j+1 < len(da.diff) {
		da.diff[j+1] -= val
	}
}

func (da *NumArray) result() []int {
	res := make([]int, len(da.diff))
	res[0] = da.diff[0]
	for i := 1; i < len(da.diff); i++ {
		res[i] = res[i-1] + da.diff[i]
	}
	return res
}

func GetModifiedArray(length int, updates [][]int) []int {
	res := make([]int, length)
	diff := Constructor(res)
	for _, item := range updates {
		diff.increment(item[0], item[1], item[2])
	}
	res = diff.result()
	return res
}

func GetModifiedArrayByGoroutine(length int, updates [][]int) []int {
	res := make([]int, length)

	var wg sync.WaitGroup
	wg.Add(len(updates))
	for idx := range updates {
		go func(index int) {
			defer wg.Done()
			s := updates[index][0]
			e := updates[index][1]
			c := updates[index][2]

			for i := s; i <= e; i++ {
				res[i] += c
			}
		}(idx)
	}

	wg.Wait()

	return res
}
