// Idea: think backwards. An 'O' region is NOT surrounded only if it
// touches the border. So:
//   1) start at every 'O' on the border and mark everything connected to
//      it as safe ('S')
//   2) every 'O' that is left is surrounded -> turn it into 'X'
//   3) turn the safe cells back into 'O'
// Time: O(rows * cols).
func solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	// 1) mark the regions that touch the border
	for r := 0; r < rows; r++ {
		markSafe(board, r, 0)
		markSafe(board, r, cols-1)
	}
	for c := 0; c < cols; c++ {
		markSafe(board, 0, c)
		markSafe(board, rows-1, c)
	}

	// 2) and 3)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X' // surrounded
			} else if board[r][c] == 'S' {
				board[r][c] = 'O' // safe: back to normal
			}
		}
	}
}

// markSafe marks the 'O' at (r, c) and all 'O's connected to it with 'S'.
func markSafe(board [][]byte, r int, c int) {
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) || board[r][c] != 'O' {
		return
	}
	board[r][c] = 'S'
	markSafe(board, r+1, c)
	markSafe(board, r-1, c)
	markSafe(board, r, c+1)
	markSafe(board, r, c-1)
}
