package matrix

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if r*c != len(nums)*len(nums[0]) {
		return nums
	}

	out := make([][]int, r)
	for i := 0; i < len(out); i++ {
		out[i] = make([]int, c)
	}

	var numsR, numsC, outR, outC int
	for numsR < len(nums) && outR < r {
		out[outR][outC] = nums[numsR][numsC]

		outC++
		if outC >= c {
			outC = 0
			outR++
		}

		numsC++
		if numsC >= len(nums[0]) {
			numsC = 0
			numsR++
		}
	}

	return out
}
