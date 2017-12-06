package capital

func detectCapitalUse(word string) bool {
	if len(word) < 2 {
		return true
	}
	switch {
	case upper(word[0]) && upper(word[1]):
		for i := 2; i < len(word); i++ {
			if !upper(word[i]) {
				return false
			}
		}
	case upper(word[0]) && lower(word[1]),
		lower(word[0]) && lower(word[1]):
		for i := 2; i < len(word); i++ {
			if !lower(word[i]) {
				return false
			}
		}
	default:
		return false
	}
	return true
}

func upper(c byte) bool { return c >= 'A' && c <= 'Z' }
func lower(c byte) bool { return c >= 'a' && c <= 'z' }
