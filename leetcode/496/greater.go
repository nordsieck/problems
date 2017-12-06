package greater

func nextGreaterElement(findNums []int, nums []int) []int {
	nmap := map[int]int{}
	stack := []int{}
	for i := len(nums) - 1; i >= 0; i-- {

		// set correspondence
		j := len(stack) - 1
		for ; j >= 0 && stack[j] < nums[i]; j-- {
		}
		if j == -1 {
			nmap[nums[i]] = -1
		} else {
			nmap[nums[i]] = stack[j]
		}

		// update stack
		for j := len(stack) - 1; j >= 0 && nums[i] > stack[j]; j, stack = j-1, stack[:j] {
		}

		for {
			j := len(stack) - 1
			if j < 0 {
				break
			}

			if nums[i] < stack[j] {
				break
			}

			stack = stack[:len(stack)-1]

		}
		stack = append(stack, nums[i])
	}

	for i := 0; i < len(findNums); i++ {
		findNums[i] = nmap[findNums[i]]
	}
	return findNums
}
