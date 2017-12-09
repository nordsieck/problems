package paren

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestIsValid(t *testing.T) {
	defect.Equal(t, isValid(`{`), false)
	defect.Equal(t, isValid(`{}`), true)
	defect.Equal(t, isValid(`(]`), false)
	defect.Equal(t, isValid(`([)]`), false)
	defect.Equal(t, isValid(`{{}}`), true)
	defect.Equal(t, isValid(`({[]})`), true)
	defect.Equal(t, isValid(`]`), false)
}
