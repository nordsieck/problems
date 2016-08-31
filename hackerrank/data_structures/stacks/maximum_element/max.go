package max

import (
	"fmt"
	"io"
)

type stack []int

func (s *stack) push(n int) { *s = append(*s, n) }
func (s *stack) pop()       { *s = (*s)[:len(*s)-1] }
func (s *stack) max() int {
	max := 0
	for _, s := range *s {
		if s > max {
			max = s
		}
	}
	return max
}

func readFirst(r io.Reader) int {
	n := 0
	fmt.Fscanf(r, "%d\n", &n)
	return n
}

func readLine(r io.Reader) (typ, val int) {
	fmt.Fscan(r, &typ)
	if typ == 2 || typ == 3 {
		return typ, 0
	}
	fmt.Fscan(r, &val)
	return typ, val
}

func process(r io.Reader, w io.Writer) {
	s, n := stack{}, readFirst(r)
	for i := 0; i < n; i++ {
		typ, val := readLine(r)
		switch typ {
		case 1:
			s.push(val)
		case 2:
			s.pop()
		case 3:
			fmt.Fprintln(w, s.max())
		}
	}
}
