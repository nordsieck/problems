package island

func islandPerimeter(grid [][]int) int {
	sum := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 1 {
				if c-1 < 0 || grid[r][c-1] == 0 {
					sum++
				}
				if c+1 == len(grid[r]) || grid[r][c+1] == 0 {
					sum++
				}
				if r-1 < 0 || grid[r-1][c] == 0 {
					sum++
				}
				if r+1 == len(grid) || grid[r+1][c] == 0 {
					sum++
				}
			}
		}
	}
	return sum
}
