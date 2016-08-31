package left

import (
	"bytes"
	"testing"

	"github.com/nordsieck/defect"
)

func TestReadArray(t *testing.T) {
	n := 5
	in := bytes.NewBufferString("1 2 3 4 5\n")
	expected := []int{1, 2, 3, 4, 5}

	out, err := readArray(in, n)
	defect.Equal(t, err, nil)
	defect.DeepEqual(t, out, expected)
}

func TestWriteSlice(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	expected := "1 2 3 4 5"

	out := &bytes.Buffer{}
	writeSlice(out, in)

	defect.Equal(t, out.String(), expected)
}

func TestWriteRotated(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	rotated := 4
	expected := "5 1 2 3 4\n"

	out := &bytes.Buffer{}
	writeRotated(out, in, rotated)

	defect.Equal(t, out.String(), expected)
}

func TestProcess(t *testing.T) {
	in := bytes.NewBufferString(`5 4
1 2 3 4 5
`)
	expected := `5 1 2 3 4
`

	out := &bytes.Buffer{}
	process(in, out)

	defect.Equal(t, out.String(), expected)
}
