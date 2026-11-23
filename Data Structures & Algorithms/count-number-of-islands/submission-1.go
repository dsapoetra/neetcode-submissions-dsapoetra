// Idea: scan every cell. When we find land ('1') it is a new island:
// count it, then use DFS to turn the whole island into water ('0')
// so we never count it again.
// Time: O(rows * cols).
func numIslands(grid [][]byte) int {
	count := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				count++
				sinkIsland(grid, r, c)
			}
		}
	}
	return count
}

// sinkIsland turns the land at (r, c) and all land connected to it into water.
func sinkIsland(grid [][]byte, r int, c int) {
	// stop if we are outside the grid or this cell is not land
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] != '1' {
		return
	}

	grid[r][c] = '0' // sink this cell

	// continue in the 4 directions
	sinkIsland(grid, r+1, c)
	sinkIsland(grid, r-1, c)
	sinkIsland(grid, r, c+1)
	sinkIsland(grid, r, c-1)
}
