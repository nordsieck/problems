package numbers

import (
	"testing"

	"github.com/nordsieck/defect"
)

func f(i []int) []int { return findDisappearedNumbers(i) }

func TestFindDisappearedNumbers(t *testing.T) {
	defect.DeepEqual(t, f([]int{1, 1}), []int{2})
	defect.DeepEqual(t, f([]int{4, 3, 2, 7, 8, 2, 3, 1}), []int{5, 6})
}
