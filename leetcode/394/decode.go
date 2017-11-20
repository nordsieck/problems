package decode

import (
	"regexp"
	"strconv"
)

func decodeString(s string) string {
	if s == `` {
		return ``
	}
	return parse(s)
}

func parse(s string) string {
	var result string

	for {
		next, rest := getNext(s)
		result += next
		s = rest
		if s == `` {
			break
		}

	}

	return result
}

func getNext(s string) (next, rest string) {
	letter := regexp.MustCompile(`^[a-z]+`)
	number := regexp.MustCompile(`^([0-9]+)\[`)

	if match := letter.FindStringSubmatch(s); len(match) > 0 {
		return match[0], s[len(match[0]):]
	} else if match := number.FindStringSubmatch(s); len(match) > 0 {
		num, _ := strconv.Atoi(match[1])
		i := len(match[0])
		brackets := 1
		for ; i < len(s) && brackets > 0; i++ {
			if s[i] == `[`[0] {
				brackets++
			} else if s[i] == `]`[0] {
				brackets--
			}
		}

		internal := s[len(match[0]) : i-1]
		parsed := parse(internal)

		result := ""
		for n := 0; n < num; n++ {
			result += parsed
		}
		return result, s[i:]
	}
	return `error`, `error` // error
}
