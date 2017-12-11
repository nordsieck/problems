package tree

import (
	"testing"

	"github.com/nordsieck/defect"
)

func t2(t *TreeNode) string { return tree2str(t) }

func TestTree2Str(t *testing.T) {
	defect.Equal(t, t2(nil), ``)
	defect.Equal(t, t2(&TreeNode{Val: 1}), `1`)
	defect.Equal(t, t2(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}), `1(2)`)
	defect.Equal(t, t2(&TreeNode{Val: 1, Right: &TreeNode{Val: 3}}), `1()(3)`)
	defect.Equal(t, t2(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}), `1(2(4))(3)`)
	defect.Equal(t, t2(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}), `1(2()(4))(3)`)
}
