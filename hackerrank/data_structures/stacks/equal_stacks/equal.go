package equal

import (
	"fmt"
	"io"
)

type stack struct {
	top, sum int
	stack    []int
}

func readStack(r io.Reader, length int) stack {
	s := stack{stack: make([]int, length)}

	for i := length - 1; i >= 0; i-- {
		fmt.Fscan(r, &s.stack[i])
	}
	return s
}

func (s *stack) next() (sum int, end bool) {
	if s.top == len(s.stack) {
		return s.sum, true
	}
	return s.sum + s.stack[s.top], false
}

func (s *stack) inc() {
	if s.top == len(s.stack) {
		return
	}
	s.sum += s.stack[s.top]
	s.top += 1
}

type stacks [3]stack

func readStacks(r io.Reader) stacks {
	var lens [3]int
	fmt.Fscanf(r, "%d %d %d\n", &lens[0], &lens[1], &lens[2])

	s := stacks{
		readStack(r, lens[0]),
		readStack(r, lens[1]),
		readStack(r, lens[2]),
	}
	return s
}

func (s stacks) next() int {
	sums, ends := [3]int{}, [3]bool{}
	for i := 0; i < len(s); i++ {
		sums[i], ends[i] = s[i].next()
	}

	if sums[0] < sums[1] || ends[1] {
		if sums[0] < sums[2] || ends[2] {
			if ends[0] {
				return -1
			}
			return 0
		}
		if ends[2] {
			return -1
		}
		return 2
	}
	if sums[1] < sums[2] || ends[2] {
		if ends[1] {
			return -1
		}
		return 1
	}
	if ends[2] {
		return -1
	}
	return 2
}

func (s stacks) top() int {
	top := 0

	for n := s.next(); n != -1; n = s.next() {
		s[n].inc()
		if s[0].sum == s[1].sum && s[1].sum == s[2].sum {
			top = s[0].sum
		}
	}
	return top
}

func process(r io.Reader, w io.Writer) {
	s := readStacks(r)
	fmt.Fprintln(w, s.top())
}
