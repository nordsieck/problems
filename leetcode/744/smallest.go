package smallest

func nextGreatestLetter(letters []byte, target byte) byte {
	seen := map[byte]struct{}{}
	for i := 0; i < len(letters); i++ {
		seen[letters[i]] = struct{}{}
	}

	for i := uint8(1); i < 26; i++ {
		curr := (uint8(target)-uint8('a')+i)%26 + uint8('a')
		if _, ok := seen[curr]; ok {
			return curr
		}
	}

	// should never happen
	return 0
}
