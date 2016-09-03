package rectangle

import (
	"fmt"
	"io"
)

type list []int

const max = 1<<31 - 1

func (l list) smallestElement() (index, val int) {
	index, val = -1, max
	for i, v := range l {
		if v < val {
			index, val = i, v
		}
	}
	return index, val
}

func (l list) largestRectangle() int {
	if len(l) == 0 {
		return 0
	}
	idx, val := l.smallestElement()

	left := l[:idx].largestRectangle()
	right := l[idx+1:].largestRectangle()
	whole := len(l) * val

	if left > right {
		if left > whole {
			return left
		}
		return whole
	}
	if right > whole {
		return right
	}
	return whole
}

func process(r io.Reader, w io.Writer) error {
	n, buff := 0, 0
	if _, err := fmt.Fscan(r, &n); err != nil {
		return err
	}

	l := make(list, n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Fscan(r, &buff); err != nil {
			return err
		}
		l[i] = buff
	}
	_, err := fmt.Fprintln(w, l.largestRectangle())
	return err
}
