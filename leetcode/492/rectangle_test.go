package rectangle

import (
	"testing"

	"github.com/nordsieck/defect"
)

func c(area int) []int { return constructRectangle(area) }

func TestConstructRectangle(t *testing.T) {
	defect.DeepEqual(t, c(3), []int{3, 1})
	defect.DeepEqual(t, c(4), []int{2, 2})
}
