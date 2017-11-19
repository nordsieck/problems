package validation

func validUtf8(data []int) bool {
	mask := []int{
		3 << 6,  // 11
		1 << 7,  // 1
		7 << 5,  // 111
		15 << 4, // 1111
		31 << 3, // 11111
	}
	test := []int{
		2 << 6,  // 10
		0,       // 0
		6 << 5,  // 110
		14 << 4, // 1110
		30 << 3, // 11110
	}

	if len(data) == 0 {
		return false
	}

	for len(data) > 0 {
		switch {
		case len(data) >= 1 && data[0]&mask[1]^test[1] == 0:
			data = data[1:]
		case len(data) >= 2 && data[0]&mask[2]^test[2] == 0 &&
			data[1]&mask[0]^test[0] == 0:
			data = data[2:]
		case len(data) >= 3 && data[0]&mask[3]^test[3] == 0 &&
			data[1]&mask[0]^test[0] == 0 &&
			data[2]&mask[0]^test[0] == 0:
			data = data[3:]
		case len(data) >= 4 && data[0]&mask[4]^test[4] == 0 &&
			data[1]&mask[0]^test[0] == 0 &&
			data[2]&mask[0]^test[0] == 0 &&
			data[3]&mask[0]^test[0] == 0:
			data = data[4:]
		default:
			return false
		}
	}
	return true
}
