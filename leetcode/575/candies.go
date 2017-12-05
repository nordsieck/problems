package candies

func distributeCandies(candies []int) int {
	types := map[int]struct{}{}
	for _, c := range candies {
		types[c] = struct{}{}
	}

	if l := len(candies) / 2; l < len(types) {
		return l
	}
	return len(types)
}
