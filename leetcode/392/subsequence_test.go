package subsequence

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestIsSubsequence(t *testing.T) {
	defect.Equal(t, isSubsequence("abc", "ahbgdc"), true)
	defect.Equal(t, isSubsequence("axc", "ahbgdc"), false)
}
