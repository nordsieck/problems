package atoms

import (
	"regexp"
	"sort"
	"strconv"
)

var (
	atom   = regexp.MustCompile(`([A-Z][a-z]*)([0-9]*)`)
	number = regexp.MustCompile(`([0-9]*)`)
)

func countOfAtoms(formula string) string {
	count := Process(formula)

	keys := make(sort.StringSlice, 0, len(count))
	for k := range count {
		keys = append(keys, k)
	}
	keys.Sort()

	var result string
	for _, k := range keys {
		result = result + k + itoa(count[k])
	}
	return result
}

type Result struct {
	freq map[string]int
	rest string
}

func Atom(formula string) Result {
	match := atom.FindStringSubmatch(formula)
	return Result{map[string]int{match[1]: atoi(match[2])}, string(formula[len(match[0]):])}
}

func Paren(formula string) Result {
	lparen, rparen := `(`[0], `)`[0]

	if len(formula) == 0 || formula[0] != lparen {
		return Result{}
	}

	i, l, r := 0, 0, 0
	for i = 0; i < len(formula); i++ {
		if formula[i] == lparen {
			l++
		} else if formula[i] == rparen {
			r++
		}
		if l == r {
			break
		}
	}

	freq := Process(string(formula[1:i]))
	formula = string(formula[i+1:])

	match := number.FindStringSubmatch(formula)
	freq = Mul(freq, atoi(match[1]))

	return Result{freq, formula[len(match[0]):]}
}

func Process(formula string) map[string]int {
	var result map[string]int
	var fn func(string) Result

	for len(formula) > 0 {
		if formula[0] == "("[0] {
			fn = Paren
		} else {
			fn = Atom
		}
		r := fn(formula)
		result = Add(result, r.freq)
		formula = r.rest
	}
	return result
}

func Add(maps ...map[string]int) map[string]int {
	result := map[string]int{}

	for _, m := range maps {
		for k, v := range m {
			if _, ok := result[k]; !ok {
				result[k] = v
			} else {
				result[k] += v
			}
		}
	}

	return result
}

func Mul(m map[string]int, num int) map[string]int {
	for k, v := range m {
		m[k] = v * num
	}
	return m
}

func atoi(s string) int {
	if s == `` {
		return 1
	}
	i, _ := strconv.Atoi(s)
	return i
}

func itoa(i int) string {
	if i == 1 {
		return ""
	}
	return strconv.Itoa(i)
}
