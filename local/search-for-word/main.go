// Problem: Word Search
// Link:    https://neetcode.io/problems/search-for-word
// Source:  Data Structures & Algorithms/search-for-word/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./search-for-word
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

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

// ===================== Test cases =====================

func main() {
	// byteGrid("AB", "CD") builds the board: each string is one row.

	// Example 1 from the problem: C -> A -> T.
	check("example 1: CAT", exist(byteGrid("ABCD", "SAAT", "ACAE"), "CAT"), true)

	// Example 2 from the problem: there is no T next to the "BA" path.
	check("example 2: BAT", exist(byteGrid("ABCD", "SAAT", "ACAE"), "BAT"), false)

	// The classic examples.
	check("classic: ABCCED", exist(byteGrid("ABCE", "SFCS", "ADEE"), "ABCCED"), true)
	check("classic: SEE", exist(byteGrid("ABCE", "SFCS", "ADEE"), "SEE"), true)

	// "ABCB" would need to use the same B twice.
	check("classic: ABCB (cell used twice)", exist(byteGrid("ABCE", "SFCS", "ADEE"), "ABCB"), false)

	// A board with one cell.
	check("one cell, found", exist(byteGrid("A"), "A"), true)
	check("one cell, wrong letter", exist(byteGrid("A"), "B"), false)
	check("one cell, word too long", exist(byteGrid("A"), "AA"), false)

	// A -> B -> D -> C goes around the square.
	check("around the square", exist(byteGrid("AB", "CD"), "ABDC"), true)

	// B and C are only diagonal neighbors: not allowed.
	check("diagonal is not allowed", exist(byteGrid("AB", "CD"), "ABCD"), false)

	// All 6 cells are needed.
	check("uses every cell", exist(byteGrid("AAA", "AAA"), "AAAAAA"), true)

	// 7 letters but only 6 cells.
	check("more letters than cells", exist(byteGrid("AAA", "AAA"), "AAAAAAA"), false)

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

// check compares the answer we got with the answer we want, and prints
// PASS or FAIL for that test case.
func check(name string, got interface{}, want interface{}) {
	if fmt.Sprint(got) == fmt.Sprint(want) {
		passed++
		fmt.Printf("PASS  %s\n", name)
	} else {
		failed++
		fmt.Printf("FAIL  %s\n      got:  %v\n      want: %v\n", name, got, want)
	}
}

// report prints the total. The program exits with an error if a test failed.
func report() {
	fmt.Printf("\n%d passed, %d failed\n", passed, failed)
	if failed > 0 {
		os.Exit(1)
	}
}
