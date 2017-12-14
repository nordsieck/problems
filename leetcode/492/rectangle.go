package rectangle

import "math"

func constructRectangle(area int) []int {
	sqrt := int(math.Sqrt(float64(area)))
	for i := sqrt; i > 1; i-- {
		if area%i == 0 {
			return []int{area / i, i}
		}
	}
	return []int{area, 1}
}
