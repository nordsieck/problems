package palindrome

func isPalindrome(x int) bool {
	top := 1

	if x < 0 {
		return false
	}

	for ; top <= x; top *= 10 {
	}

	for bot, top := 1, top/10; bot < top; bot, top = bot*10, top/10 {
		high := x / top % 10
		low := x / bot % 10

		if high != low {
			return false
		}
	}

	return true
}
