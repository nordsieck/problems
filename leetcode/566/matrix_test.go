package matrix

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestMatrixReshape(t *testing.T) {
	defect.DeepEqual(t, matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4), [][]int{{1, 2, 3, 4}})
	defect.DeepEqual(t, matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4), [][]int{{1, 2}, {3, 4}})
}
