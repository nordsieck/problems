package trim

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < L {
		return trimBST(root.Right, L, R)
	} else if root.Val > R {
		return trimBST(root.Left, L, R)
	}
	return &TreeNode{Val: root.Val, Left: trimBST(root.Left, L, R), Right: trimBST(root.Right, L, R)}
}
