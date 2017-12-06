package longest

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestFindLUSlength(t *testing.T) {
	defect.Equal(t, findLUSlength("abc", "cdc"), 3)
}
