package greater

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	result, _ := convert(root, 0)
	return result
}

func convert(root *TreeNode, add int) (t *TreeNode, sum int) {
	if root == nil {
		return nil, add
	}

	var rightSum, leftSum int
	result := &TreeNode{}
	result.Right, rightSum = convert(root.Right, add)
	result.Val = root.Val + rightSum
	result.Left, leftSum = convert(root.Left, result.Val)

	return result, leftSum
}
