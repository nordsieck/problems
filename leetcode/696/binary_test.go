package binary

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestCountBinarySubstrings(t *testing.T) {
	defect.Equal(t, c(``), 0)
	defect.Equal(t, c(`1`), 0)
	defect.Equal(t, c(`00`), 0)
	defect.Equal(t, c(`10`), 1)
	defect.Equal(t, c(`11110`), 1)
	defect.Equal(t, c(`00001`), 1)
	defect.Equal(t, c(`00110011`), 6)
}

func c(s string) int { return countBinarySubstrings(s) }
