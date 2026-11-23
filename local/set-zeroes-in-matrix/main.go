// Problem: Set Matrix Zeroes
// Link:    https://neetcode.io/problems/set-zeroes-in-matrix
// Source:  Data Structures & Algorithms/set-zeroes-in-matrix/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./set-zeroes-in-matrix
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: do it in two passes so new zeros do not confuse us.
//
//	pass 1: remember which rows and which columns contain a zero
//	pass 2: set a cell to zero if its row or its column was remembered
//
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

// ===================== Test cases =====================

func main() {
	// setZeroes changes the matrix in place and returns nothing, so this
	// small shortcut runs it and then gives the matrix back.
	zeroed := func(matrix [][]int) [][]int {
		setZeroes(matrix)
		return matrix
	}

	// Example 1 from the problem: two zeros wipe out everything.
	check("example 1", `zeroed([][]int{{0, 1}, {1, 0}})`,
		zeroed([][]int{{0, 1}, {1, 0}}), [][]int{{0, 0}, {0, 0}})

	// Example 2 from the problem: one zero in the middle makes a cross.
	check("example 2", `zeroed([][]int{{1, 2, 3}, {4, 0, 5}, {6, 7, 8}})`,
		zeroed([][]int{{1, 2, 3}, {4, 0, 5}, {6, 7, 8}}), [][]int{{1, 0, 3}, {0, 0, 0}, {6, 0, 8}})

	// No zeros: nothing changes.
	check("no zeros", `zeroed([][]int{{1, 2}, {3, 4}})`,
		zeroed([][]int{{1, 2}, {3, 4}}), [][]int{{1, 2}, {3, 4}})

	// One cell.
	check("single zero", `zeroed([][]int{{0}})`,
		zeroed([][]int{{0}}), [][]int{{0}})
	check("single one", `zeroed([][]int{{1}})`,
		zeroed([][]int{{1}}), [][]int{{1}})

	// The classic example: zeros in two corners of the first row.
	check("classic", `zeroed([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}})`,
		zeroed([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}), [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}})

	// One row.
	check("one row", `zeroed([][]int{{1, 0, 3}})`,
		zeroed([][]int{{1, 0, 3}}), [][]int{{0, 0, 0}})

	// One column.
	check("one column", `zeroed([][]int{{1}, {0}, {3}})`,
		zeroed([][]int{{1}, {0}, {3}}), [][]int{{0}, {0}, {0}})

	// Tricky: the new zeros must NOT cause more zeros. Only the last row
	// and last column change here; the 1 in the top-left stays.
	check("zero in a corner", `zeroed([][]int{{1, 2}, {3, 0}})`,
		zeroed([][]int{{1, 2}, {3, 0}}), [][]int{{1, 0}, {0, 0}})

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
