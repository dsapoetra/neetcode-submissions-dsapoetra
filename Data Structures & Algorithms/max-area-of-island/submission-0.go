// Idea: look at every cell. When we find land, walk over the whole island
// (DFS: up, down, left, right) and count its cells. We set every visited
// cell to 0 ("sink" it) so it is never counted twice.
// Time: O(rows * cols).
func maxAreaOfIsland(grid [][]int) int {
	best := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				area := islandArea(grid, r, c)
				if area > best {
					best = area
				}
			}
		}
	}
	return best
}

// islandArea counts the land cells connected to (r, c) and sinks them.
func islandArea(grid [][]int, r int, c int) int {
	// outside the grid, or water
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == 0 {
		return 0
	}

	grid[r][c] = 0 // sink this cell so we do not count it again

	return 1 +
		islandArea(grid, r+1, c) +
		islandArea(grid, r-1, c) +
		islandArea(grid, r, c+1) +
		islandArea(grid, r, c-1)
}
