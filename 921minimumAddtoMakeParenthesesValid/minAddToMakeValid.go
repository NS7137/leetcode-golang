package minimumaddtomakeparenthesesvalid

func MinAddToMakeValid(s string) int {
	// left 记录左括号需求
	left := 0
	// right 记录右括号需求
	right := 0

	size := len(s)
	for i := 0; i < size; i++ {
		if s[i] == '(' {
			right++
		}

		if s[i] == ')' {
			right--
			if right == -1 {
				right = 0
				// 插入一个左括号
				left++
			}
		}
	}
	return left + right
}
