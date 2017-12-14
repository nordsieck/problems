package rang

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestMaxCount(t *testing.T) {
	defect.Equal(t, maxCount(3, 3, [][]int{{2, 2}, {3, 3}}), 4)
}
