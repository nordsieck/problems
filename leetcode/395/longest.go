package longest

import "strings"

func longestSubstring(s string, k int) int {
	count := map[rune]int{}
	for _, r := range s {
		count[r]++
	}

	for r, c := range count {
		if c < k {
			max := 0
			candidates := strings.Split(s, string(r))

			for _, c := range candidates {
				if result := longestSubstring(c, k); result > max {
					max = result
				}
			}
			return max
		}
	}
	return len(s)
}
