package smallestsubsequenceofdistinctcharacters

import "container/list"

//与 316. 去除重复字母 一模一样
func SmallestSubsequence(s string) string {
	stack := list.New()
	exist := [26]bool{}
	count := [26]int{}

	for _, c := range s {
		count[c-'a']++
	}

	for _, c := range s {
		count[c-'a']--

		if exist[c-'a'] {
			continue
		}

		for stack.Len() != 0 && stack.Back().Value.(rune) > c {
			if count[stack.Back().Value.(rune)-'a'] == 0 {
				break
			}
			exist[stack.Back().Value.(rune)-'a'] = false
			stack.Remove(stack.Back())
		}

		stack.PushBack(c)
		exist[c-'a'] = true
	}

	res := []rune{}
	for stack.Len() > 0 {
		v := stack.Remove(stack.Front())
		res = append(res, v.(rune))
	}

	return string(res)
}
