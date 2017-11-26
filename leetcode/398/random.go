package random

import "math/rand"

type Solution struct{ nums []int }

func Constructor(nums []int) Solution { return Solution{nums} }

func (this *Solution) Pick(target int) int {
	idx, num := -1, 0
	for i, v := range this.nums {
		if v == target {
			if rand.Intn(num+1) == num {
				idx = i
			}
			num++
		}
	}
	return idx
}
