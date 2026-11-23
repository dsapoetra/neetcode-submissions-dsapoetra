// Idea (BFS, minute by minute): put all rotten fruits in a queue. In each
// minute, every rotten fruit in the queue rots its fresh neighbors, and
// those neighbors become the queue for the next minute.
// Count the fresh fruits at the start; if some are still fresh when the
// queue is empty, they can never rot, so the answer is -1.
// Time: O(rows * cols).
func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	queue := [][2]int{} // the fruits that became rotten most recently
	fresh := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			} else if grid[r][c] == 1 {
				fresh++
			}
		}
	}

	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	minutes := 0

	for len(queue) > 0 && fresh > 0 {
		minutes++
		next := [][2]int{}

		for _, cell := range queue {
			for _, d := range directions {
				r := cell[0] + d[0]
				c := cell[1] + d[1]

				// inside the grid and fresh: it rots now
				if r >= 0 && r < rows && c >= 0 && c < cols && grid[r][c] == 1 {
					grid[r][c] = 2
					fresh--
					next = append(next, [2]int{r, c})
				}
			}
		}
		queue = next
	}

	if fresh > 0 {
		return -1
	}
	return minutes
}
