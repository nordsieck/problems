package numbers

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		var idx int
		if nums[i] < 0 {
			idx = -nums[i] - 1
		} else {
			idx = nums[i] - 1
		}
		if nums[idx] > 0 {
			nums[idx] *= -1
		}
	}
	result := []int{}
	for i, v := range nums {
		if v > 0 {
			result = append(result, i+1)
		}
	}
	return result
}
