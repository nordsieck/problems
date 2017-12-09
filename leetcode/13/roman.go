package roman

func romanToInt(s string) int {
	val := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var curr, sum int

	for i := len(s) - 1; i >= 0; i-- {
		if digit := val[s[i]]; digit >= curr {
			curr = digit
			sum += digit
		} else {
			sum -= digit
		}
	}
	return sum
}
