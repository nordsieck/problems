package island

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestIslandPerimeter(t *testing.T) {
	defect.Equal(t, islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}), 16)
}
