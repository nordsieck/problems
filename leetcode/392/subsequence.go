package subsequence

func isSubsequence(s string, t string) bool {
	if s == `` {
		return true
	}

	var counter int
	for i := 0; i < len(t); i++ {
		if s[counter] == t[i] {
			if counter == len(s)-1 {
				return true
			}
			counter++
		}
	}
	return false
}
