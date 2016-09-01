package balanced

import (
	"fmt"
	"io"
)

type stack []byte

func (s *stack) push(b byte) { *s = append(*s, b) }

func (s stack) lastMatch() bool {
	if len(s) < 2 {
		return false
	}
	switch (s)[len(s)-2] {
	case '(':
		return s[len(s)-1] == ')'
	case '[':
		return s[len(s)-1] == ']'
	case '{':
		return s[len(s)-1] == '}'
	}
	return false
}

func matching(r io.Reader) bool {
	buffer, s := make([]byte, 1), stack{}
	for _, err := r.Read(buffer); err == nil && buffer[0] != '\n'; _, err = r.Read(buffer) {
		s.push(buffer[0])
		if s.lastMatch() {
			s = s[:len(s)-2]
		}
	}
	return len(s) == 0
}

func process(r io.Reader, w io.Writer) {
	lines := 0
	fmt.Fscanf(r, "%d\n", &lines)

	for i := 0; i < lines; i++ {
		if m := matching(r); m {
			fmt.Fprintln(w, "YES")
		} else {
			fmt.Fprintln(w, "NO")
		}
	}
}
