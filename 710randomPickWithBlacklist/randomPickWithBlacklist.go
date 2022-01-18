package randompickwithblacklist

import "math/rand"

type Solution struct {
	size      int
	maps      map[int]int
	blacklist []int
}

func Constructor(n int, blacklist []int) Solution {
	// 最终数组中的元素个数
	size := n - len(blacklist)
	solution := Solution{
		size:      size,
		maps:      map[int]int{},
		blacklist: blacklist,
	}
	// 将bl 加入map，目的仅把键存入
	for _, bl := range solution.blacklist {
		solution.maps[bl] = 666
	}

	// 最后一个元素索引
	last := n - 1
	// bl 索引 映射到最后
	for _, bl := range solution.blacklist {
		// 如果 bl 已经在 [size,n) 直接忽略
		if bl >= solution.size {
			continue
		}
		//跳过所有黑名单中的数字
		_, ok := solution.maps[last]
		for ok {
			last--
			_, ok = solution.maps[last]
		}
		solution.maps[bl] = last
		last--
	}

	return solution

}

// bl 元素 [size,n), wl 元素[0,size)
func (s *Solution) Pick() int {
	index := rand.Int() % s.size
	// 命中黑名单 需要被映射到其他位置
	if val, ok := s.maps[index]; ok {
		return val
	}
	return index
}
