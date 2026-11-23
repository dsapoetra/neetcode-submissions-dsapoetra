// Idea: instead of asking "can water from this cell reach the ocean?",
// go backwards: start AT each ocean and climb uphill (to cells that are
// the same height or higher). Every cell we can climb to can send its
// water down to that ocean.
// Do this once for the Pacific (top + left edges) and once for the
// Atlantic (bottom + right edges). Cells reached by both are the answer.
// Time: O(rows * cols).
func pacificAtlantic(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])

	pacific := make([][]bool, rows)  // pacific[r][c] = can reach the Pacific
	atlantic := make([][]bool, rows) // atlantic[r][c] = can reach the Atlantic
	for r := 0; r < rows; r++ {
		pacific[r] = make([]bool, cols)
		atlantic[r] = make([]bool, cols)
	}

	// left edge touches the Pacific, right edge touches the Atlantic
	for r := 0; r < rows; r++ {
		climb(heights, r, 0, pacific)
		climb(heights, r, cols-1, atlantic)
	}
	// top edge touches the Pacific, bottom edge touches the Atlantic
	for c := 0; c < cols; c++ {
		climb(heights, 0, c, pacific)
		climb(heights, rows-1, c, atlantic)
	}

	// keep the cells that reach both oceans
	res := [][]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pacific[r][c] && atlantic[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}

// climb marks (r, c) as reached, then moves to every neighbor that is
// the same height or higher (DFS).
func climb(heights [][]int, r int, c int, reached [][]bool) {
	reached[r][c] = true

	// the 4 directions: down, up, right, left
	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, move := range moves {
		nr := r + move[0]
		nc := c + move[1]

		// skip neighbors outside the grid
		if nr < 0 || nc < 0 || nr >= len(heights) || nc >= len(heights[0]) {
			continue
		}
		// skip neighbors we already visited or that are lower
		if reached[nr][nc] || heights[nr][nc] < heights[r][c] {
			continue
		}
		climb(heights, nr, nc, reached)
	}
}
