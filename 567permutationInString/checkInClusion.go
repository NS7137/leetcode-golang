package permutationinstring

func CheckInclusion(key string, s string) bool {

	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, c := range []byte(key) {
		need[c]++
	}
	// valid 为满足个数
	left, right, valid := 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		// 窗口内数据更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// valid满足时收缩
		for right-left >= len(key) {
			if valid == len(need) {
				return true
			}
			// 移除的字符
			d := s[left]
			left++
			// 窗口内数据更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
