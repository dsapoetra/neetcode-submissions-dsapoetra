// Every time we find land, count one island and sink the whole island with DFS.
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

func sinkIsland(grid [][]byte, r int, c int) {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] != '1' {
		return
	}
	grid[r][c] = '0'
	sinkIsland(grid, r+1, c)
	sinkIsland(grid, r-1, c)
	sinkIsland(grid, r, c+1)
	sinkIsland(grid, r, c-1)
}
