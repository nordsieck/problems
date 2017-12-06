package alt

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestHasAlternatingBits(t *testing.T) {
	defect.Equal(t, hasAlternatingBits(5), true)
	defect.Equal(t, hasAlternatingBits(7), false)
	defect.Equal(t, hasAlternatingBits(11), false)
	defect.Equal(t, hasAlternatingBits(10), true)
	defect.Equal(t, hasAlternatingBits(32), false)
	defect.Equal(t, hasAlternatingBits(0), true)
	defect.Equal(t, hasAlternatingBits(1), true)
}
