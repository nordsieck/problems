package paren

func isValid(s string) bool {
	match := map[byte]byte{'{': '}', '(': ')', '[': ']'}
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if _, ok := match[s[i]]; ok {
			stack = append(stack, s[i])
		} else if len(stack) == 0 {
			return false
		} else if match[stack[len(stack)-1]] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
