// Idea (backtracking): try to start the word from every cell.
// From a cell that matches the current letter, move to the 4 neighbors
// to match the next letter. Mark the cell as used while we are on this
// path so it is not used twice, and un-mark it when we come back.
func exist(board [][]byte, word string) bool {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if searchFrom(board, word, r, c, 0) {
				return true
			}
		}
	}
	return false
}

// searchFrom returns true if word[index:] can be found starting at (r, c).
func searchFrom(board [][]byte, word string, r int, c int, index int) bool {
	// every letter was matched
	if index == len(word) {
		return true
	}
	// outside the board
	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
		return false
	}
	// wrong letter (a used cell holds '#', so it never matches)
	if board[r][c] != word[index] {
		return false
	}

	letter := board[r][c]
	board[r][c] = '#' // mark as used

	found := searchFrom(board, word, r+1, c, index+1) ||
		searchFrom(board, word, r-1, c, index+1) ||
		searchFrom(board, word, r, c+1, index+1) ||
		searchFrom(board, word, r, c-1, index+1)

	board[r][c] = letter // un-mark (backtrack)
	return found
}
