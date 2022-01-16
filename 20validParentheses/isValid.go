package validparentheses

func IsValid(s string) bool {
	left := []rune{}
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			left = append(left, c)
		case ')', ']', '}':
			if len(left) == 0 || left[len(left)-1] != leftOf(c) {
				return false
			}
			left = left[:len(left)-1]
		}
	}

	return len(left) == 0

}

func leftOf(c rune) rune {
	if c == '}' {
		return '{'
	}
	if c == ')' {
		return '('
	}
	return '['
}
