package flood

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestFloodFill(t *testing.T) {
	defect.DeepEqual(t, floodFill([][]int{{1}}, 0, 0, 2), [][]int{{2}})
	defect.DeepEqual(t, floodFill([][]int{{1, 0}}, 0, 0, 2), [][]int{{2, 0}})
	defect.DeepEqual(t, floodFill([][]int{{1, 1}}, 0, 0, 2), [][]int{{2, 2}})
	defect.DeepEqual(t, floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2), [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}})
	defect.DeepEqual(t, floodFill([][]int{{0, 0, 0}, {1, 1, 1}}, 1, 1, 1), [][]int{{0, 0, 0}, {1, 1, 1}})
}
