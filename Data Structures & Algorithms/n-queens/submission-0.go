// Idea (backtracking): put one queen in every row, from top to bottom.
// In each row try every column that is not attacked. A square is attacked
// if a queen is already in the same column or on the same diagonal:
//   - cells on one "\" diagonal all have the same (row - col)
//   - cells on one "/" diagonal all have the same (row + col)
// After trying a column, take the queen back off and try the next one.
func solveNQueens(n int) [][]string {
	res := [][]string{}

	colUsed := make([]bool, n)
	diagDown := make([]bool, 2*n) // "\" diagonals, index row - col + n
	diagUp := make([]bool, 2*n)   // "/" diagonals, index row + col
	queenCol := make([]int, n)    // queenCol[row] = column of the queen in that row

	var place func(row int)
	place = func(row int) {
		if row == n {
			// all queens are placed: draw the board
			board := []string{}
			for r := 0; r < n; r++ {
				line := ""
				for c := 0; c < n; c++ {
					if queenCol[r] == c {
						line += "Q"
					} else {
						line += "."
					}
				}
				board = append(board, line)
			}
			res = append(res, board)
			return
		}

		for col := 0; col < n; col++ {
			if colUsed[col] || diagDown[row-col+n] || diagUp[row+col] {
				continue // attacked
			}

			// put the queen here
			colUsed[col], diagDown[row-col+n], diagUp[row+col] = true, true, true
			queenCol[row] = col

			place(row + 1)

			// take the queen off again
			colUsed[col], diagDown[row-col+n], diagUp[row+col] = false, false, false
		}
	}

	place(0)
	return res
}
