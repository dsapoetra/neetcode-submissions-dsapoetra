// Problem:    N-Queens
// Difficulty: Hard
// Link:       https://neetcode.io/problems/n-queens
// Source:     Data Structures & Algorithms/n-queens/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Place n queens on an n x n chess board so that no two queens attack each
// other (no two in the same row, column or diagonal). Return all solutions.
// A board is a list of strings: 'Q' is a queen, '.' is an empty square.
//
// Example 1:
//
//	Input:  n = 4
//	Output: [[".Q..", "...Q", "Q...", "..Q."],
//	         ["..Q.", "Q...", "...Q", ".Q.."]]
//
// Example 2:
//
//	Input:  n = 1
//	Output: [["Q"]]
//
// How to run (from the "local" folder):
//
//	go run ./140-n-queens
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (backtracking): put one queen in every row, from top to bottom.
// In each row try every column that is not attacked. A square is attacked
// if a queen is already in the same column or on the same diagonal:
//   - cells on one "\" diagonal all have the same (row - col)
//   - cells on one "/" diagonal all have the same (row + col)
//
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

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: the two solutions for a 4 x 4 board.
	check("example 1: n = 4", `solveNQueens(4)`,
		solveNQueens(4), [][]string{
			{".Q..", "...Q", "Q...", "..Q."},
			{"..Q.", "Q...", "...Q", ".Q.."},
		})

	// Example 2 from the problem: one queen on a 1 x 1 board.
	check("example 2: n = 1", `solveNQueens(1)`,
		solveNQueens(1), [][]string{{"Q"}})

	// 2 x 2 and 3 x 3 have no solution.
	check("n = 2: impossible", `len(solveNQueens(2))`,
		len(solveNQueens(2)), 0)
	check("n = 3: impossible", `len(solveNQueens(3))`,
		len(solveNQueens(3)), 0)

	// Known numbers of solutions.
	check("n = 5 gives 10", `len(solveNQueens(5))`,
		len(solveNQueens(5)), 10)
	check("n = 6 gives 4", `len(solveNQueens(6))`,
		len(solveNQueens(6)), 4)
	check("n = 8 gives 92", `len(solveNQueens(8))`,
		len(solveNQueens(8)), 92)

	report()
}

// ===================== Helpers for the tests =====================

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
