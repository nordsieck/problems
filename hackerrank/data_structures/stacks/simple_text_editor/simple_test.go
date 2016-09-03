package simple

import (
	"bytes"
	"testing"

	"github.com/nordsieck/defect"
)

func TestProcessLine(t *testing.T) {
	type Case struct {
		in, out string
		buffer  []string
	}

	buffer := []string{""}
	cases := []Case{
		{in: "1 abc\n", out: "", buffer: []string{"", "abc"}},
		{in: "3 3\n", out: "c\n", buffer: []string{"", "abc"}},
		{in: "2 3\n", out: "", buffer: []string{"", "abc", ""}},
		{in: "1 xy\n", out: "", buffer: []string{"", "abc", "", "xy"}},
		{in: "3 2\n", out: "y\n", buffer: []string{"", "abc", "", "xy"}},
		{in: "4\n", out: "", buffer: []string{"", "abc", ""}},
		{in: "4\n", out: "", buffer: []string{"", "abc"}},
		{in: "3 1\n", out: "a\n", buffer: []string{"", "abc"}},
	}

	for _, c := range cases {
		in := bytes.NewBufferString(c.in)
		buff := &bytes.Buffer{}

		buffer = ProcessLine(buffer, in, buff)
		defect.Equal(t, buff.String(), c.out)
		defect.DeepEqual(t, buffer, c.buffer)
	}
}

func TestProcess(t *testing.T) {
	in := bytes.NewBufferString(`8
1 abc
3 3
2 3
1 xy
3 2
4
4
3 1
`)
	out := `c
y
a
`
	buffer := &bytes.Buffer{}
	Process(in, buffer)
	defect.Equal(t, buffer.String(), out)
}
