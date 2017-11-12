package pivot

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestPivotIndex(t *testing.T) {
	defect.Equal(t, pivotIndex([]int{1, 7, 3, 6, 5, 6}), 3)
	defect.Equal(t, pivotIndex([]int{0}), 0)
	defect.Equal(t, pivotIndex([]int{1, 1}), -1)
	defect.Equal(t, pivotIndex([]int{}), -1)
}
