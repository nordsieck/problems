package smallest

import (
	"testing"

	"github.com/nordsieck/defect"
)

func n(letters []byte, target byte) byte {
	return nextGreatestLetter(letters, target)
}

func TestNextGreatestLetter(t *testing.T) {
	defect.Equal(t, n([]byte{'c', 'f', 'j'}, 'a'), byte('c'))
	defect.Equal(t, n([]byte{'c', 'f', 'j'}, 'c'), byte('f'))
	defect.Equal(t, n([]byte{'c', 'f', 'j'}, 'f'), byte('j'))
	defect.Equal(t, n([]byte{'c', 'f', 'j'}, 'j'), byte('c'))
	defect.Equal(t, n([]byte{'c', 'f', 'j'}, 'k'), byte('c'))
}
