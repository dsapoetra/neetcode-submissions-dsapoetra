// Idea (DFS + remember): the longest path that STARTS at a cell is
// 1 + the longest path starting at a bigger neighbor.
// Many paths pass through the same cells, so we remember the answer of
// every cell (memo) and compute each cell only once.
// We never walk in circles, because we only step to bigger numbers.
// Time: O(rows * cols).
func longestIncreasingPath(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	memo := make([][]int, rows) // 0 = not computed yet
	for i := range memo {
		memo[i] = make([]int, cols)
	}

	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// longestFrom returns the length of the longest increasing path starting at (r, c).
	var longestFrom func(r int, c int) int
	longestFrom = func(r int, c int) int {
		if memo[r][c] != 0 {
			return memo[r][c]
		}

		best := 1 // the cell alone
		for _, d := range directions {
			nr := r + d[0]
			nc := c + d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && matrix[nr][nc] > matrix[r][c] {
				length := 1 + longestFrom(nr, nc)
				if length > best {
					best = length
				}
			}
		}

		memo[r][c] = best
		return best
	}

	answer := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if longestFrom(r, c) > answer {
				answer = longestFrom(r, c)
			}
		}
	}
	return answer
}
