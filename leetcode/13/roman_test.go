package roman

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestRomanToInt(t *testing.T) {
	defect.Equal(t, romanToInt(`I`), 1)
	defect.Equal(t, romanToInt(`M`), 1000)
	defect.Equal(t, romanToInt(`CMXCIX`), 999)
	defect.Equal(t, romanToInt(`DCXXI`), 621)
}
