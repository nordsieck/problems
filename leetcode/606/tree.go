package tree

import "strconv"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ``
	}
	left, right := tree2str(t.Left), tree2str(t.Right)
	if right != `` {
		return strconv.Itoa(t.Val) + `(` + left + `)(` + right + `)`
	} else if left != `` {
		return strconv.Itoa(t.Val) + `(` + left + `)`
	}
	return strconv.Itoa(t.Val)
}
