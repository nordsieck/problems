package balanced

import (
	"bytes"
	"testing"

	"github.com/nordsieck/defect"
)

func TestStack_LastMatch(t *testing.T) {
	cases := map[string]bool{
		"(":      false,
		"()":     true,
		"((({}":  true,
		"([{}])": false,
	}

	for in, out := range cases {
		defect.Equal(t, stack(in).lastMatch(), out)
	}
}

func TestMatching(t *testing.T) {
	cases := map[string]bool{
		"(":      false,
		"()":     true,
		"((({}":  false,
		"([{}])": true,
	}

	for in, out := range cases {
		buf := bytes.NewBufferString(in)
		defect.Equal(t, matching(buf), out)
	}
}

func TestProcess(t *testing.T) {
	in := bytes.NewBufferString(`3
{[()]}
{[(])}
{{[[(())]]}}
`)
	out := `YES
NO
YES
`

	buf := &bytes.Buffer{}
	process(in, buf)
	defect.Equal(t, buf.String(), out)
}
