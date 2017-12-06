package avg

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type count struct {
	sum, num int
}

func sum(root *TreeNode) []count {
	if root == nil {
		return nil
	}

	long, short := sum(root.Left), sum(root.Right)
	if len(long) < len(short) {
		long, short = short, long
	}

	result := make([]count, len(long)+1)
	result[0] = count{sum: root.Val, num: 1}
	for i := 0; i < len(short); i++ {
		result[i+1] = count{sum: long[i].sum + short[i].sum, num: long[i].num + short[i].num}
	}
	for i := len(short); i < len(long); i++ {
		result[i+1] = long[i]
	}
	return result
}

func averageOfLevels(root *TreeNode) []float64 {
	s := sum(root)
	result := make([]float64, len(s))
	for i := 0; i < len(s); i++ {
		result[i] = float64(s[i].sum) / float64(s[i].num)
	}
	return result
}
