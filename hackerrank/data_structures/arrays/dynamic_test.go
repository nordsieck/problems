package arrays

import (
	"bytes"
	"testing"

	"github.com/nordsieck/defect"
)

func TestProcess(t *testing.T) {
	cases := map[string]string{
		`2 3
1 0 5
1 1 7
1 0 3
`: ``,
		`2 5
1 0 5
1 1 7
1 0 3
2 1 0
2 1 1
		`: `7
3
`,
	}

	for in, expected := range cases {
		in := bytes.NewBufferString(in)
		buffer := &bytes.Buffer{}
		_, err := Process(in, buffer)
		defect.Equal(t, err, nil)
		defect.Equal(t, buffer.String(), expected)
	}
}

func TestNewList(t *testing.T) {
	defect.DeepEqual(t, newList(2, 3), &seqList{array: [][]int{nil, nil, nil}, lastAns: 2})
}

func TestSeqList_One(t *testing.T) {
	s := &seqList{array: [][]int{nil, nil, nil}, lastAns: 5}

	s.one(1, 1)
	defect.DeepEqual(t, s, &seqList{array: [][]int{nil, []int{1}, nil}, lastAns: 5})

	s.one(2, 2)
	defect.DeepEqual(t, s, &seqList{array: [][]int{nil, []int{1, 2}, nil}, lastAns: 5})

	s = &seqList{array: [][]int{nil, nil}}
	s.one(0, 5)
	s.one(1, 7)
	s.one(0, 3)
	defect.DeepEqual(t, s, &seqList{array: [][]int{[]int{5, 3}, []int{7}}})
}

func TestSeqList_Two(t *testing.T) {
	s := &seqList{array: [][]int{nil, []int{1}, nil}, lastAns: 5}

	s.two(1, 1)
	defect.DeepEqual(t, s, &seqList{array: [][]int{nil, []int{1}, nil}, lastAns: 1})

	s = &seqList{array: [][]int{[]int{5, 3}, []int{7}}}

	s.two(1, 0)
	defect.DeepEqual(t, s.lastAns, 7)

	s.two(1, 1)
	defect.DeepEqual(t, s.lastAns, 3)
}

func TestReadFirst(t *testing.T) {
	in := bytes.NewBufferString("2 5\n")
	n, q, err := readFirst(in)
	defect.Equal(t, err, nil)
	defect.Equal(t, n, 2)
	defect.Equal(t, q, 5)
}

func TestReadQuery(t *testing.T) {
	in := bytes.NewBufferString("1 0 5\n")
	typ, x, y, err := readQuery(in)
	defect.Equal(t, err, nil)
	defect.Equal(t, typ, 1)
	defect.Equal(t, x, 0)
	defect.Equal(t, y, 5)
}
