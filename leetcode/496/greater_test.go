package greater

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestNextGreaterElement(t *testing.T) {
	defect.DeepEqual(t, nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}), []int{-1, 3, -1})
	defect.DeepEqual(t, nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}), []int{3, -1})
}
