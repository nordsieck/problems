package random

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestSolution_Pick(t *testing.T) {
	s := Constructor([]int{1, 2, 3, 4, 5})
	defect.Equal(t, s.Pick(2), 1)

	s = Constructor([]int{5, 3, 1})
	defect.Equal(t, s.Pick(1), 2)
}
