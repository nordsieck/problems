package dividing

func selfDividingNumbers(left int, right int) []int {
	var result []int
	for i := left; i <= right; i++ {
		if selfDividing(i) {
			result = append(result, i)
		}
	}
	return result
}

func digits(i int) map[int]struct{} {
	result := map[int]struct{}{}
	for ; i > 0; i /= 10 {
		result[i%10] = struct{}{}
	}
	return result
}

func selfDividing(i int) bool {
	ds := digits(i)
	for d := range ds {
		if d == 0 || i%d != 0 {
			return false
		}
	}
	return true
}
