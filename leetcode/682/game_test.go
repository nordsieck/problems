package game

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestCalPoints(t *testing.T) {
	defect.Equal(t, calPoints([]string{`5`, `2`, `C`, `D`, `+`}), 30)
	defect.Equal(t, calPoints([]string{`5`, `-2`, `4`, `C`, `D`, `9`, `+`, `+`}), 27)
}
