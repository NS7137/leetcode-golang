package findmedianfromdatastream

import "container/heap"

// 倒三角 再一分为二 大顶堆元素 小于 小顶堆元素
type MedianFinder struct {
	small *maxheap // 大顶堆
	large *minheap // 小顶堆
}

func Constructor() MedianFinder {
	return MedianFinder{
		small: &maxheap{},
		large: &minheap{},
	}
}

// 维护large和small的元素个数之差不超过1，large顶元素大于等于small顶元素
// 往large里添加元素，不能直接添加，而是要先往small里添加，再把small的堆顶元素加到large,small添加同理
func (mf *MedianFinder) AddNum(num int) {
	small, large := mf.small, mf.large
	if small.Len() >= large.Len() {
		heap.Push(small, num)
		tmp := heap.Pop(small)
		heap.Push(large, tmp)
	} else {
		heap.Push(large, num)
		tmp := heap.Pop(large)
		heap.Push(small, tmp)
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	small, large := *mf.small, *mf.large
	// 元素不一样多，多的那个堆顶元素就是中位数
	if small.Len() > large.Len() {
		return float64(small[0])
	} else if large.Len() > small.Len() {
		return float64(large[0])
	}
	return float64(small[0]+large[0]) / 2
}

type maxheap []int
type minheap []int

func (h maxheap) Len() int {
	return len(h)
}

func (h maxheap) Less(i int, j int) bool {
	return h[i] > h[j]
}

func (h maxheap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxheap) Pop() interface{} {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[0 : n-1]
	return res
}

func (h minheap) Len() int {
	return len(h)
}

func (h minheap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h minheap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minheap) Pop() interface{} {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[0 : n-1]
	return res
}
