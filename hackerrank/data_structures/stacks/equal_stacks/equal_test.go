package equal

import (
	"bytes"
	"testing"

	"github.com/nordsieck/defect"
)

func TestReadStacks(t *testing.T) {
	in := bytes.NewBufferString(`5 3 4
3 2 1 1 1
4 3 2
1 1 4 1
`)
	out := stacks{
		{stack: []int{1, 1, 1, 2, 3}},
		{stack: []int{2, 3, 4}},
		{stack: []int{1, 4, 1, 1}},
	}

	s := readStacks(in)
	defect.DeepEqual(t, s, out)
}

func TestStack_Next(t *testing.T) {
	type Case struct {
		s   stack
		sum int
		end bool
	}

	cases := []Case{
		{s: stack{top: 0}, sum: 0, end: true},
		{s: stack{top: 0, stack: []int{0}}, sum: 0, end: false},
		{s: stack{top: 1, stack: []int{0}}, sum: 0, end: true},
		{s: stack{top: 0, stack: []int{1}}, sum: 1, end: false},
		{s: stack{top: 1, sum: 1, stack: []int{1, 1}}, sum: 2, end: false},
	}

	for _, c := range cases {
		sum, end := c.s.next()
		defect.Equal(t, sum, c.sum)
		defect.Equal(t, end, c.end)
	}
}

func TestStack_Inc(t *testing.T) {
	s := stack{top: 1, sum: 3, stack: []int{3, 2, 1}}
	s.inc()
	defect.Equal(t, s.sum, 5)
}

func TestStacks_Next(t *testing.T) {
	type Case struct {
		result int
		s      stacks
	}

	cases := []Case{
		{result: -1, s: stacks{}},
		{result: -1, s: stacks{{top: 1, stack: []int{0}}, {top: 1, stack: []int{0}}, {top: 1, stack: []int{0}}}},
		{result: 0, s: stacks{{top: 0, stack: []int{0}}, {top: 0, stack: []int{1}}, {top: 0, stack: []int{2}}}},
		{result: 0, s: stacks{
			{top: 0, stack: []int{20}},
			{top: 3, sum: 9, stack: []int{1, 3, 5}},
			{top: 2, sum: 15, stack: []int{15, 0}},
		}},
	}

	for _, c := range cases {
		defect.Equal(t, c.s.next(), c.result)
	}
}

func TestStacks_Top(t *testing.T) {
	type Case struct {
		s   stacks
		top int
	}

	cases := []Case{
		{top: 0, s: stacks{}},
		{top: 0, s: stacks{{}, {}, {stack: []int{0}}}},
		{top: 0, s: stacks{{}, {}, {stack: []int{1}}}},
		{top: 1, s: stacks{{stack: []int{0, 0, 1, 15}}, {stack: []int{0, 1}}, {stack: []int{1, 20}}}},
	}

	for _, c := range cases {
		defect.Equal(t, c.s.top(), c.top)
	}
}

func TestProcess(t *testing.T) {
	in := bytes.NewBufferString(`5 3 4
3 2 1 1 1
4 3 2
1 1 4 1
`)
	buffer := &bytes.Buffer{}
	process(in, buffer)
	defect.Equal(t, buffer.String(), "5\n")
}
