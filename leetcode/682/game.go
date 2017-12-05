package game

import (
	"strconv"
)

func calPoints(ops []string) int {
	stack := []string{}

	// filter the c's
	for _, o := range ops {
		if o == "C" {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, o)
		}
	}

	points := []int{}
	for _, s := range stack {
		switch s {
		case "+":
			points = append(points, points[len(points)-2]+points[len(points)-1])
		case "D":
			points = append(points, points[len(points)-1]*2)
		default: // int
			val, _ := strconv.Atoi(s)
			points = append(points, val)
		}
	}

	sum := 0
	for _, p := range points {
		sum += p
	}
	return sum
}
