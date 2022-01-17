package minimuminsertionstobalanceaparenthesesstring

func MinInsertions(s string) int {

	size := len(s)
	res, need := 0, 0

	// 一个左 对应 两个右
	for i := 0; i < size; i++ {
		if s[i] == '(' {
			need += 2
			// 若右需求为奇，结果加1，右需求-1
			if need%2 == 1 {
				res++
				need--
			}
		}

		// 遇右，需要插入一个左，但1左对2右，则对右需求为1
		if s[i] == ')' {
			// 对右需求-1
			need--
			// 右太多，需插入左
			if need == -1 {
				res++
				// 对右需求变为1
				need = 1
			}
		}
	}
	return res + need
}
