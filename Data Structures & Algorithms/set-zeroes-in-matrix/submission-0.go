// Idea: do it in two passes so new zeros do not confuse us.
//   pass 1: remember which rows and which columns contain a zero
//   pass 2: set a cell to zero if its row or its column was remembered
// Time: O(rows * cols), extra space: O(rows + cols).
func setZeroes(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	zeroRow := make([]bool, rows) // zeroRow[r] = row r has a zero
	zeroCol := make([]bool, cols) // zeroCol[c] = column c has a zero

	// pass 1: find the zeros
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 0 {
				zeroRow[r] = true
				zeroCol[c] = true
			}
		}
	}

	// pass 2: write the zeros
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if zeroRow[r] || zeroCol[c] {
				matrix[r][c] = 0
			}
		}
	}
}
