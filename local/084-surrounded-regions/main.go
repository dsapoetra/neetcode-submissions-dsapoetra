// Problem:    Surrounded Regions
// Difficulty: Medium
// Link:       https://neetcode.io/problems/surrounded-regions
// Source:     Data Structures & Algorithms/surrounded-regions/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a board of 'X' and 'O', find every region of connected 'O's (up,
// down, left, right) that is completely surrounded by 'X', and turn it into
// 'X'. A region that touches the border of the board is NOT surrounded.
// Change the board in place.
//
// Example 1:
//
//	Input:  board = [
//	          ["X","X","X","X"],
//	          ["X","O","O","X"],
//	          ["X","O","O","X"],
//	          ["X","X","X","O"]
//	        ]
//	Output:        [
//	          ["X","X","X","X"],
//	          ["X","X","X","X"],
//	          ["X","X","X","X"],
//	          ["X","X","X","O"]
//	        ]
//
// How to run (from the "local" folder):
//
//	go run ./084-surrounded-regions
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: think backwards. An 'O' region is NOT surrounded only if it
// touches the border. So:
//  1. start at every 'O' on the border and mark everything connected to
//     it as safe ('S')
//  2. every 'O' that is left is surrounded -> turn it into 'X'
//  3. turn the safe cells back into 'O'
//
// Time: O(rows * cols).
func solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	// 1) mark the regions that touch the border
	for r := 0; r < rows; r++ {
		// left and right border
		markSafe(board, r, 0)
		markSafe(board, r, cols-1)
	}
	for c := 0; c < cols; c++ {
		// top and bottom border
		markSafe(board, 0, c)
		markSafe(board, rows-1, c)
	}

	// 2) and 3)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// still 'O' means it was not reached from the border
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
	// stop: outside the board, or not an 'O' (an 'X', or already marked 'S')
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) || board[r][c] != 'O' {
		return
	}
	// mark it; this also stops us from visiting it again
	board[r][c] = 'S'
	// continue in the 4 directions
	markSafe(board, r+1, c)
	markSafe(board, r-1, c)
	markSafe(board, r, c+1)
	markSafe(board, r, c-1)
}

// ===================== Test cases =====================

func main() {
	// solve changes the board in place and returns nothing, so this small
	// shortcut builds the board from strings, runs solve, and gives the
	// rows back as strings.
	solved := func(rows ...string) []string {
		board := byteGrid(rows...)
		solve(board)
		result := []string{}
		for _, row := range board {
			result = append(result, string(row))
		}
		return result
	}

	// Example 1 from the problem: the O in the corner touches the border and stays.
	check("example 1", `solved("XXXX", "XOOX", "XOOX", "XXXO")`,
		solved("XXXX", "XOOX", "XOOX", "XXXO"), []string{"XXXX", "XXXX", "XXXX", "XXXO"})

	// The classic example: the O at the bottom touches the border and stays.
	check("classic", `solved("XXXX", "XOOX", "XXOX", "XOXX")`,
		solved("XXXX", "XOOX", "XXOX", "XOXX"), []string{"XXXX", "XXXX", "XXXX", "XOXX"})

	// One cell is always on the border.
	check("single X", `solved("X")`,
		solved("X"), []string{"X"})
	check("single O", `solved("O")`,
		solved("O"), []string{"O"})

	// One O completely surrounded.
	check("one surrounded O", `solved("XXX", "XOX", "XXX")`,
		solved("XXX", "XOX", "XXX"), []string{"XXX", "XXX", "XXX"})

	// The inner O is connected to an O on the border: both stay.
	check("connected to the border", `solved("XOX", "XOX", "XXX")`,
		solved("XOX", "XOX", "XXX"), []string{"XOX", "XOX", "XXX"})

	// Every O is on the border.
	check("ring of O", `solved("OOO", "OXO", "OOO")`,
		solved("OOO", "OXO", "OOO"), []string{"OOO", "OXO", "OOO"})

	// Two separate surrounded regions.
	check("two surrounded regions", `solved("XXXXX", "XOXOX", "XXXXX")`,
		solved("XXXXX", "XOXOX", "XXXXX"), []string{"XXXXX", "XXXXX", "XXXXX"})

	// Diagonal does not count as connected: the inner O is surrounded.
	check("diagonal is not connected", `solved("OXX", "XOX", "XXX")`,
		solved("OXX", "XOX", "XXX"), []string{"OXX", "XXX", "XXX"})

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
