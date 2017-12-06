package ones

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	defect.Equal(t, findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}), 3)
}
