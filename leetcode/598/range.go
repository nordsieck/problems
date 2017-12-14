package rang

func maxCount(m int, n int, ops [][]int) int {
	lowA, lowB := m, n

	for _, op := range ops {
		if op[0] < lowA {
			lowA = op[0]
		}
		if op[1] < lowB {
			lowB = op[1]
		}
	}
	return lowA * lowB
}
