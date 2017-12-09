package dividing

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestDigits(t *testing.T) {
	defect.DeepEqual(t, digits(1), map[int]struct{}{1: {}})
	defect.DeepEqual(t, digits(128), map[int]struct{}{1: {}, 2: {}, 8: {}})
}

func TestSelfDividing(t *testing.T) {
	defect.Equal(t, selfDividing(128), true)
	defect.Equal(t, selfDividing(1), true)
	defect.Equal(t, selfDividing(13), false)
	defect.Equal(t, selfDividing(20), false)
}

func TestSelfDividingNumbers(t *testing.T) {
	defect.DeepEqual(t, selfDividingNumbers(128, 128), []int{128})
	defect.DeepEqual(t, selfDividingNumbers(1, 22), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22})
}
