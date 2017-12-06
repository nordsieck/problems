package island

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestDFS(t *testing.T) {
	seen := map[[2]int]struct{}{}
	grid := [][]int{{0, 1, 0}, {0, 1, 0}, {0, 1, 0}}
	defect.DeepEqual(t, dfs(&seen, grid, 0, 0), 0)
	defect.DeepEqual(t, seen, map[[2]int]struct{}{})

	defect.DeepEqual(t, dfs(&seen, grid, 0, 1), 3)
	defect.DeepEqual(t, seen, map[[2]int]struct{}{{0, 1}: {}, {1, 1}: {}, {2, 1}: {}})

	defect.DeepEqual(t, dfs(&seen, grid, 1, 1), 0)
	defect.DeepEqual(t, seen, map[[2]int]struct{}{{0, 1}: {}, {1, 1}: {}, {2, 1}: {}})

	seen = map[[2]int]struct{}{}
	grid = [][]int{{1, 1}, {1, 1}}
	defect.DeepEqual(t, dfs(&seen, grid, 0, 0), 4)
	defect.DeepEqual(t, seen, map[[2]int]struct{}{{0, 0}: {}, {0, 1}: {}, {1, 0}: {}, {1, 1}: {}})
}

func TestMaxAreaOfIsland(t *testing.T) {
	defect.Equal(t, maxAreaOfIsland([][]int{{0, 1, 0}, {0, 1, 0}, {0, 1, 0}}), 3)
	defect.Equal(t, maxAreaOfIsland([][]int{{1, 1}, {1, 1}}), 4)
	defect.Equal(t, maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}), 6)
}
