// Idea (BFS from all treasures at once): put every treasure chest (0) in
// a queue. Then spread out step by step: the land cells next to a chest
// get distance 1, their neighbors distance 2, and so on.
// Because BFS reaches closer cells first, the first time we reach a land
// cell is the shortest distance to ANY chest.
// A land cell that was not reached yet still holds the big "INF" number.
// Time: O(rows * cols).
func islandsAndTreasure(grid [][]int) {
	const inf = 2147483647
	rows := len(grid)
	cols := len(grid[0])

	// start with all treasure chests
	queue := [][2]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		for _, d := range directions {
			r := cell[0] + d[0]
			c := cell[1] + d[1]

			// outside the grid, or not an unvisited land cell (water, chest, or already reached)
			if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != inf {
				continue
			}

			grid[r][c] = grid[cell[0]][cell[1]] + 1 // one step further than the cell we came from
			queue = append(queue, [2]int{r, c})
		}
	}
}
