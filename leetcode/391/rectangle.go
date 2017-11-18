package rectangle

import (
	"math"
)

func isRectangleCover(rectangles [][]int) bool {
	const X, Y = 0, 1

	lower, upper := [2]int{math.MaxInt32, math.MaxInt32}, [2]int{math.MinInt32, math.MinInt32}
	points := map[[2]int]int{}
	var sum int

	for _, r := range rectangles {
		sum += (r[X+2] - r[X]) * (r[Y+2] - r[Y])

		if r[X] <= lower[X] && r[Y] <= lower[Y] {
			lower[X], lower[Y] = r[X], r[Y]
		}
		if r[X+2] >= upper[X] && r[Y+2] >= upper[Y] {
			upper[X], upper[Y] = r[X+2], r[Y+2]
		}
		points[[2]int{r[X], r[Y]}]++
		points[[2]int{r[X], r[Y+2]}]++
		points[[2]int{r[X+2], r[Y]}]++
		points[[2]int{r[X+2], r[Y+2]}]++
	}

	if sum != (upper[X]-lower[X])*(upper[Y]-lower[Y]) {
		return false
	}

	for point, count := range points {
		if point == lower || point == upper ||
			point[X] == lower[X] && point[Y] == upper[Y] ||
			point[X] == upper[X] && point[Y] == lower[Y] {
			if count == 1 {
				continue
			} else {
				return false
			}
		} else if count%2 == 0 {
			continue
		}

		return false
	}

	return true
}
