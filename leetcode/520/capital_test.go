package capital

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestDetectCapitalUse(t *testing.T) {
	defect.Equal(t, detectCapitalUse("USA"), true)
	defect.Equal(t, detectCapitalUse("mL"), false)
	defect.Equal(t, detectCapitalUse("FlaG"), false)
}
