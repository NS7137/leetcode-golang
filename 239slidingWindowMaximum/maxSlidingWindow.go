package slidingwindowmaximum

import "math"

func MaxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	if len(nums) == 0 {
		return res
	}

	for i := 0; i <= len(nums)-k; i++ {
		max := math.MinInt32
		for j := i; j < k+i; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		res = append(res, max)
	}
	return res
}

type MonotonicQueue struct {
	q []int
}

func (mq *MonotonicQueue) Push(n int) {
	// 将小于自己的元素都删除
	for len(mq.q) > 0 && n > mq.q[len(mq.q)-1] {
		mq.q = mq.q[:len(mq.q)-1]
	}
	mq.q = append(mq.q, n)
}

func (mq *MonotonicQueue) Max() int {
	// 队头元素最大
	return mq.q[0]
}

func (mq *MonotonicQueue) Pop(n int) {
	// 判断是否不存在，就不用删除
	if n == mq.q[0] {
		mq.q = mq.q[1:]
	}
}

func MaxSlidingWindowByMonoicQueue(nums []int, k int) []int {
	mq := &MonotonicQueue{}
	// res 长度0 容量nums.length-k+1
	res := make([]int, 0, len(nums)-k+1)
	for i, v := range nums {
		if i < k-1 {
			// 前 k-1 填满
			mq.Push(v)
		} else {
			// 窗口向前滑动
			// 移入新元素
			mq.Push(v)
			// 将窗口中最大元素计入结果集
			res = append(res, mq.Max())
			// 移除最后的元素
			mq.Pop(nums[i-k+1])
		}
	}
	return res
}
