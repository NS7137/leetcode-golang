package rangesumquery

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{[]int{}}
	}
	tmp := make([]int, len(nums)+1)
	tmp[0], tmp[1] = 0, nums[0]
	for i := 2; i < len(tmp); i++ {
		tmp[i] = tmp[i-1] + nums[i-1]
	}
	return NumArray{arr: tmp}

}

func (na *NumArray) SumRange(left int, right int) int {
	return na.arr[right+1] - na.arr[left]
}
