package sum

import (
	"testing"

	"github.com/nordsieck/defect"
)

func f(root *TreeNode, k int) bool { return findTarget(root, k) }

func TestFindTarget(t *testing.T) {
	tree := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}
	defect.Equal(t, f(tree, 9), true)
	defect.Equal(t, f(tree, 28), false)
}
