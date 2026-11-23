// Idea (dynamic programming): we can only move right or down, so
// the number of ways to reach a cell = ways from above + ways from the left.
// The first row and first column have only 1 way each.
// Time: O(m * n).
func uniquePaths(m int, n int) int {
	// ways[r][c] = number of paths to reach cell (r, c)
	ways := make([][]int, m)
	for r := 0; r < m; r++ {
		ways[r] = make([]int, n)
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if r == 0 || c == 0 {
				ways[r][c] = 1 // top row or left column: only one way
			} else {
				ways[r][c] = ways[r-1][c] + ways[r][c-1]
			}
		}
	}
	return ways[m-1][n-1]
}
