package char

import (
	"testing"

	"github.com/nordsieck/defect"
)

func is(i []int) bool { return isOneBitCharacter(i) }

func TestIsOneBitCharacter(t *testing.T) {
	defect.Equal(t, is([]int{1, 0, 0}), true)
	defect.Equal(t, is([]int{1, 1, 1, 0}), false)
}
