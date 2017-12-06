package ones

func findMaxConsecutiveOnes(nums []int) int {
	var max, current int

	for _, n := range nums {
		if n == 1 {
			current++
			continue
		}
		if current > max {
			max = current
		}
		current = 0
	}
	if current > max {
		max = current
	}
	return max
}
