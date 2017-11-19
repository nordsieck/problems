package validation

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestValidUtf8(t *testing.T) {
	defect.Equal(t, validUtf8([]int{5}), true)
	defect.Equal(t, validUtf8([]int{255}), false)
	defect.Equal(t, validUtf8([]int{197, 130, 1}), true)
	defect.Equal(t, validUtf8([]int{235, 140, 4}), false)
}
