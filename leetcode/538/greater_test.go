package greater

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestConverBST(t *testing.T) {
	defect.DeepEqual(t, convertBST(&TreeNode{Val: 5, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 13}}),
		&TreeNode{Val: 18, Left: &TreeNode{Val: 20}, Right: &TreeNode{Val: 13}})
	defect.DeepEqual(t, convertBST(&TreeNode{Val: 2, Left: &TreeNode{Val: 0, Left: &TreeNode{Val: -4}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 3}}),
		&TreeNode{Val: 5, Left: &TreeNode{Val: 6, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 3}})
}
