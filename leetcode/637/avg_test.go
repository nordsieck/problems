package avg

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestAverageOfLevels(t *testing.T) {
	defect.DeepEqual(t, averageOfLevels(&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}), []float64{3, 14.5, 11})
}
