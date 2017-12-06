package island

func dfs(seen *map[[2]int]struct{}, grid [][]int, r, c int) int {
	if r < 0 || r >= len(grid) {
		return 0
	} else if c < 0 || c >= len(grid[r]) {
		return 0
	}
	if grid[r][c] == 0 {
		return 0
	}
	if _, ok := (*seen)[[2]int{r, c}]; ok {
		return 0
	}

	(*seen)[[2]int{r, c}] = struct{}{}
	return 1 + dfs(seen, grid, r-1, c) + dfs(seen, grid, r+1, c) + dfs(seen, grid, r, c-1) + dfs(seen, grid, r, c+1)
}

func maxAreaOfIsland(grid [][]int) int {
	seen := map[[2]int]struct{}{}
	var max int

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			current := dfs(&seen, grid, r, c)
			if current > max {
				max = current
			}
		}
	}
	return max
}
