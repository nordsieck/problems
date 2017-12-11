package flood

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	stackQ := [][2]int{{sr, sc}}
	stackM := map[[2]int]struct{}{{sr, sc}: {}}

	orig := image[sr][sc]

	if orig == newColor {
		return image
	}

	for len(stackQ) > 0 {
		r, c := stackQ[0][0], stackQ[0][1]
		delete(stackM, stackQ[0])
		stackQ = stackQ[1:]
		image[r][c] = newColor

		loc := [2]int{r - 1, c}
		if _, ok := stackM[loc]; r > 0 && image[r-1][c] == orig && !ok {
			stackM[loc] = struct{}{}
			stackQ = append(stackQ, loc)
		}
		loc = [2]int{r + 1, c}
		if _, ok := stackM[loc]; r < len(image)-1 && image[r+1][c] == orig && !ok {
			stackM[loc] = struct{}{}
			stackQ = append(stackQ, loc)
		}
		loc = [2]int{r, c - 1}
		if _, ok := stackM[loc]; c > 0 && image[r][c-1] == orig && !ok {
			stackM[loc] = struct{}{}
			stackQ = append(stackQ, loc)
		}
		loc = [2]int{r, c + 1}
		if _, ok := stackM[loc]; c < len(image[0])-1 && image[r][c+1] == orig && !ok {
			stackM[loc] = struct{}{}
			stackQ = append(stackQ, loc)
		}
	}
	return image
}
