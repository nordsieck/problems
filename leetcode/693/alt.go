package alt

func hasAlternatingBits(n int) bool {
	full := 0
	for ; full < n; full = full<<1 + 1 {
	}

	return n^(n>>1) == full
}
