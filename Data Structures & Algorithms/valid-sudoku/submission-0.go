// Idea: walk over every cell once. For each digit remember that we have
// seen it in its row, in its column and in its 3x3 box. If we see the same
// digit again in one of those three places, the board is not valid.
// Time: O(81), the board always has 81 cells.
func isValidSudoku(board [][]byte) bool {
	// rows[r][d] = true means digit d+1 was already seen in row r
	var rows [9][9]bool
	var cols [9][9]bool
	var boxes [9][9]bool

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			// empty cells are always fine
			if board[r][c] == '.' {
				continue
			}

			digit := int(board[r][c] - '1') // '1'..'9' becomes 0..8
			box := (r/3)*3 + c/3            // boxes are numbered 0..8, left to right, top to bottom

			if rows[r][digit] || cols[c][digit] || boxes[box][digit] {
				return false // seen before
			}

			rows[r][digit] = true
			cols[c][digit] = true
			boxes[box][digit] = true
		}
	}
	return true
}
