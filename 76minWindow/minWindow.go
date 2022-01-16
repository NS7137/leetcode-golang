package minwindow

import (
	"math"
)

func MinWindow(s string, t string) string {

	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	if len(t) > len(s) {
		return ""
	}

	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, c := range []byte(t) {
		need[c]++
	}
	// valid 为满足个数
	left, right, valid := 0, 0, 0
	start := 0
	length := math.MaxInt32
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
		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
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

	if length == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+length]
	}

}
