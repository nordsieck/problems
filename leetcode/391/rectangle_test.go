package rectangle

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestIsRectangleCover(t *testing.T) {
	defect.Equal(t, isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{3, 2, 4, 4},
		{1, 3, 2, 4},
		{2, 3, 3, 4},
	}), true)
	defect.Equal(t, isRectangleCover([][]int{
		{1, 1, 2, 3},
		{3, 1, 4, 2},
		{1, 3, 2, 4},
		{2, 2, 4, 4},
	}), false)
	defect.Equal(t, isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{1, 3, 2, 4},
		{3, 2, 4, 4},
	}), false)
	defect.Equal(t, isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{1, 3, 2, 4},
		{2, 2, 4, 4},
	}), false)
	defect.Equal(t, isRectangleCover([][]int{
		{0, 0, 4, 1},
		{0, 0, 4, 1},
	}), false)
	defect.Equal(t, isRectangleCover([][]int{
		{0, 0, 3, 3},
		{1, 1, 2, 2},
		{1, 1, 2, 2},
	}), false)
}
