package pivot

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	var idx, bot, top int
	for i := 1; i < len(nums); i++ {
		top += nums[i]
	}
	for bot != top {
		if idx == len(nums)-1 {
			return -1
		}

		bot += nums[idx]
		idx++
		top -= nums[idx]
	}
	return idx
}
