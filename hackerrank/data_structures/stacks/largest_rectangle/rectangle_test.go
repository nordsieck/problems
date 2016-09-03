package rectangle

import (
	"bytes"
	"testing"

	"github.com/nordsieck/defect"
)

func TestList_SmallestElement(t *testing.T) {
	type Case struct {
		list       list
		index, val int
	}

	cases := []Case{
		{list: list{1, 2, 3, 4, 5}, index: 0, val: 1},
		{list: list{5, 1, 4, 2, 3}, index: 1, val: 1},
	}

	for _, c := range cases {
		i, v := c.list.smallestElement()
		defect.Equal(t, i, c.index)
		defect.Equal(t, v, c.val)
	}
}

func TestList_LargestRectangle(t *testing.T) {
	cases := map[int]list{
		9: list{1, 2, 3, 4, 5},
		6: list{5, 1, 4, 2, 3},
	}

	for expected, input := range cases {
		defect.Equal(t, input.largestRectangle(), expected)
	}
}

func TestProcess(t *testing.T) {
	cases := map[string]string{
		`5
1 2 3 4 5
`: `9
`,
	}

	for in, expected := range cases {
		input := bytes.NewBufferString(in)
		output := &bytes.Buffer{}

		err := process(input, output)
		defect.Equal(t, err, nil)
		defect.Equal(t, output.String(), expected)
	}
}
