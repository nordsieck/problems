package dynamic

import (
	"fmt"
	"io"
)

type seqList struct {
	array   [][]int
	lastAns int
}

func Process(r io.Reader, w io.Writer) (*seqList, error) {
	n, q, err := readFirst(r)
	if err != nil {
		return nil, err
	}

	list := newList(0, n)
	for i := 0; i < q; i++ {
		typ, x, y, err := readQuery(r)
		if err != nil {
			return nil, err
		}
		switch typ {
		case 1:
			list.one(x, y)
		case 2:
			list.two(x, y)
			fmt.Fprintln(w, list.lastAns)
		}
	}
	return list, nil
}

func (s *seqList) index(n int) int { return (n ^ s.lastAns) % len(s.array) }

func (s *seqList) one(x, y int) {
	idx := s.index(x)
	s.array[idx] = append(s.array[idx], y)
}

func (s *seqList) two(x, y int) {
	idx := s.index(x)
	s.lastAns = s.array[idx][y%len(s.array[idx])]
}

func newList(lastAns, n int) *seqList {
	return &seqList{array: make([][]int, n), lastAns: lastAns}
}

func readFirst(r io.Reader) (n, q int, err error) {
	_, err = fmt.Fscanf(r, "%d %d\n", &n, &q)
	return n, q, err
}

func readQuery(r io.Reader) (typ, x, y int, err error) {
	_, err = fmt.Fscanf(r, "%d %d %d\n", &typ, &x, &y)
	return typ, x, y, err
}
