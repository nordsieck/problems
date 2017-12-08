package reverse

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestReverse(t *testing.T) {
	defect.Equal(t, reverse(2), 2)
	defect.Equal(t, reverse(123), 321)
	defect.Equal(t, reverse(-123), -321)
	defect.Equal(t, reverse(120), 21)
	defect.Equal(t, reverse(1), 1)
	defect.Equal(t, reverse(0), 0)
	defect.Equal(t, reverse(-0), -0)
	defect.Equal(t, reverse(-1), -1)
	defect.Equal(t, reverse(10), 1)
	defect.Equal(t, reverse(100), 1)
	defect.Equal(t, reverse(1534236469), 0)
}
