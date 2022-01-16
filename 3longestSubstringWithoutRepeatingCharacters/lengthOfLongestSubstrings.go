package longestsubstringwithoutrepeatingcharacters

func LengthOfLongestSubstring(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	window := make(map[byte]int)

	left, right, res := 0, 0, 0

	for right < len(s) {
		c := s[right]
		right++
		// 进行窗口内数据的更新
		window[c]++
		// 左侧是否要收缩 出现重复
		for window[c] > 1 {
			d := s[left]
			left++

			window[d]--
		}

		res = max(res, right-left)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
