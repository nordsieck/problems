package binary

func countBinarySubstrings(s string) int {
	var last, curr, sum int
	var one bool

	for _, ch := range s {
		if (ch == '1') != one {
			if last-curr < 0 {
				sum += last
			} else {
				sum += curr
			}
			last, curr, one = curr, 0, !one
		}
		curr++
	}
	if last-curr < 0 {
		return sum + last
	} else {
		return sum + curr
	}
}
