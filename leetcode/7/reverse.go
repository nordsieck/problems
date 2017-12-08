package reverse

func reverse(x int) int {
	neg := x < 0
	if neg {
		x *= -1
	}

	result := 0
	for base := 1; base <= x; base *= 10 {
		result *= 10
		result += x / base % 10
	}

	if result > 1<<31 {
		return 0
	}

	if neg {
		result *= -1
	}
	return result
}
