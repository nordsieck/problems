package trim

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestTrimBST(t *testing.T) {
	defect.DeepEqual(t, trimBST(&TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}}, 1, 2),
		&TreeNode{Val: 1, Right: &TreeNode{Val: 2}})
	defect.DeepEqual(t, trimBST(
		&TreeNode{Val: 3,
			Left: &TreeNode{Val: 0,
				Right: &TreeNode{Val: 2,
					Left: &TreeNode{Val: 1}}},
			Right: &TreeNode{Val: 4}}, 1, 3),
		&TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}})

}
