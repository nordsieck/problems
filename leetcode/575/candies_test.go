package candies

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestDistributeCandies(t *testing.T) {
	defect.Equal(t, distributeCandies([]int{1, 1, 2, 2, 3, 3}), 3)
	defect.Equal(t, distributeCandies([]int{1, 1, 2, 3}), 2)
}
