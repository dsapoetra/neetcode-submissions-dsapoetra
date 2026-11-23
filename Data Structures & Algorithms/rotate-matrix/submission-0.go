// Idea: rotating 90 degrees clockwise is the same as two easy steps:
//   1) transpose: swap matrix[r][c] with matrix[c][r] (flip over the diagonal)
//   2) reverse every row
// Example:  1 2 3      1 4 7      7 4 1
//           4 5 6  ->  2 5 8  ->  8 5 2
//           7 8 9      3 6 9      9 6 3
// Time: O(n^2), no extra matrix needed.
func rotate(matrix [][]int) {
	n := len(matrix)

	// step 1: transpose (only the part above the diagonal, or we swap twice)
	for r := 0; r < n; r++ {
		for c := r + 1; c < n; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}

	// step 2: reverse every row
	for r := 0; r < n; r++ {
		left, right := 0, n-1
		for left < right {
			matrix[r][left], matrix[r][right] = matrix[r][right], matrix[r][left]
			left++
			right--
		}
	}
}
