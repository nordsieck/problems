package max

import (
	"bytes"
	"testing"

	"github.com/nordsieck/defect"
)

func TestStack_Push(t *testing.T) {
	s := stack{}
	s.push(5)
	s.push(13)

	defect.DeepEqual(t, s, stack{5, 13})
}

func TestStack_Pop(t *testing.T) {
	s := stack{1, 2, 3, 4, 5}
	s.pop()
	s.pop()
	s.pop()

	defect.DeepEqual(t, s, stack{1, 2})
}

func TestStack_Max(t *testing.T) {
	s := stack{1, 5, 2, 4, 3}

	defect.Equal(t, s.max(), 5)
}

func TestReadLine(t *testing.T) {
	in := bytes.NewBufferString("1 97\n")
	typ, val := readLine(in)
	defect.Equal(t, typ, 1)
	defect.Equal(t, val, 97)

	in = bytes.NewBufferString("2\n")
	typ, val = readLine(in)
	defect.Equal(t, typ, 2)
	defect.Equal(t, val, 0)
}

func TestProcess(t *testing.T) {
	in := bytes.NewBufferString(`10
1 97
2
1 20
2
1 26
1 20
2
3
1 91
3
`)
	expected := `26
91
`
	out := &bytes.Buffer{}

	process(in, out)
	defect.Equal(t, out.String(), expected)
}
