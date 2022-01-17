package removeduplicateletters

// 去重，不能打乱s顺序，字典序最小
func RemoveDuplicateLetters(s string) string {
	stack := []rune{}
	exist := [26]bool{}
	count := [26]int{}

	for _, c := range s {
		count[c-'a']++
	}

	for _, c := range s {
		// 每遍历字符，计数-1
		count[c-'a']--

		if exist[c-'a'] {
			continue
		}

		// 栈中取出字符 大于当前元素，不符合字典序，且入栈的字符后续至少还有1个
		for len(stack) != 0 && stack[len(stack)-1] > c && count[stack[len(stack)-1]-'a'] > 0 {
			// pop
			exist[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, c)

		exist[c-'a'] = true

	}

	return string(stack)

}
