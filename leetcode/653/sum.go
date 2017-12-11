package sum

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	m := map[int]struct{}{}
	return find(root, k, &m)
}

func find(root *TreeNode, k int, m *map[int]struct{}) bool {
	if root == nil {
		return false
	}

	if _, ok := (*m)[k-root.Val]; ok {
		return true
	}
	(*m)[root.Val] = struct{}{}

	if find(root.Left, k, m) {
		return true
	}
	return find(root.Right, k, m)
}
