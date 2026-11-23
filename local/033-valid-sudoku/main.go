// Problem:    Valid Sudoku
// Difficulty: Medium
// Link:       https://neetcode.io/problems/valid-sudoku
// Source:     Data Structures & Algorithms/valid-sudoku/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a 9 x 9 Sudoku board that is partly filled ('.' is an empty cell),
// return true if it breaks no rule so far: no digit twice in a row, in a
// column, or in one of the nine 3 x 3 boxes. The board does not have to be
// solvable.
//
// Example 1:
//
//	Input:  board =
//	          1 2 . | . 3 . | . . .
//	          4 . . | 5 . . | . . .
//	          . 9 8 | . . . | . . 3
//	          ------+-------+------
//	          5 . . | . 6 . | . . 4
//	          . . . | 8 . 3 | . . 5
//	          7 . . | . 2 . | . . 6
//	          ------+-------+------
//	          . . . | . . . | 2 . .
//	          . . . | 4 1 9 | . . 8
//	          . . . | . 8 . | . 7 9
//	Output: true
//
// Example 2:
//
//	Input:  the same board, but the first cell of row 3 is 1 instead of empty
//	Output: false         (two 1s in the top-left box)
//
// How to run (from the "local" folder):
//
//	go run ./033-valid-sudoku
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

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

// ===================== Test cases =====================

func main() {
	// A valid board (example 1 from the problem). '.' is an empty cell.
	valid := []string{
		"12..3....",
		"4..5.....",
		".98.....3",
		"5...6...4",
		"...8.3..5",
		"7...2...6",
		"......2..",
		"...419..8",
		"....8..79",
	}

	// withRow returns the valid board with ONE row replaced. That lets us
	// break exactly one rule at a time.
	withRow := func(row int, text string) [][]byte {
		rows := append([]string{}, valid...)
		rows[row] = text
		return byteGrid(rows...)
	}

	// Example 1 from the problem.
	check("example 1: valid board", `isValidSudoku(byteGrid(valid...))`,
		isValidSudoku(byteGrid(valid...)), true)

	// Example 2 from the problem: a second 1 in the top-left box (and in column 0).
	check("example 2: 1 twice in a box", `isValidSudoku(withRow(2, "198.....3"))`,
		isValidSudoku(withRow(2, "198.....3")), false)

	// The same digit twice in one ROW.
	check("same digit twice in a row", `isValidSudoku(withRow(0, "12..3...1"))`,
		isValidSudoku(withRow(0, "12..3...1")), false)

	// The same digit twice in one COLUMN (a 1 at the bottom of column 0).
	check("same digit twice in a column", `isValidSudoku(withRow(8, "1...8..79"))`,
		isValidSudoku(withRow(8, "1...8..79")), false)

	// Only the BOX rule is broken: the new 9 shares no row and no column
	// with the other 9, but they are in the same 3x3 box.
	check("same digit twice in a box only", `isValidSudoku(withRow(0, "129.3...."))`,
		isValidSudoku(withRow(0, "129.3....")), false)

	// A board with no numbers at all is valid.
	check("empty board", `isValidSudoku(byteGrid(".........", ".........", ".........", ".........", ".........", ".........", ".........", ".........", "........."))`,
		isValidSudoku(byteGrid(
			".........", ".........", ".........",
			".........", ".........", ".........",
			".........", ".........", ".........")), true)

	// Changing an empty row to an empty row keeps it valid.
	check("one row emptied", `isValidSudoku(withRow(7, "........."))`,
		isValidSudoku(withRow(7, ".........")), true)

	report()
}

// ===================== Helpers for the tests =====================

// byteGrid turns rows written as strings into a [][]byte grid.
// Example: byteGrid("110", "001") is a grid with 2 rows and 3 columns.
func byteGrid(rows ...string) [][]byte {
	grid := [][]byte{}
	for _, row := range rows {
		grid = append(grid, []byte(row))
	}
	return grid
}

// passed and failed count the test results.
var passed, failed int

// check runs one test case. It prints the input, the result we got and the
// result we expected, with PASS or FAIL in front.
//
//	name:  a short description of the case
//	input: the call we are testing, written as text so it can be printed
//	got:   the answer the solution returned
//	want:  the answer we expect
func check(name string, input string, got interface{}, want interface{}) {
	status := "PASS"
	if fmt.Sprint(got) == fmt.Sprint(want) {
		passed++
	} else {
		failed++
		status = "FAIL"
	}
	fmt.Printf("%s  %s\n", status, name)
	fmt.Printf("      input:    %s\n", input)
	fmt.Printf("      result:   %s\n", show(got))
	fmt.Printf("      expected: %s\n\n", show(want))
}

// show turns a value into text. Strings get quotes, so an empty string is visible.
func show(value interface{}) string {
	if text, ok := value.(string); ok {
		return fmt.Sprintf("%q", text)
	}
	return fmt.Sprint(value)
}

// report prints the total. The program exits with an error if a test failed.
func report() {
	fmt.Printf("%d passed, %d failed\n", passed, failed)
	if failed > 0 {
		os.Exit(1)
	}
}
