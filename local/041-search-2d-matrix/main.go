// Problem:    Search a 2D Matrix
// Difficulty: Medium
// Link:       https://neetcode.io/problems/search-2d-matrix
// Source:     Data Structures & Algorithms/search-2d-matrix/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Every row of the matrix is sorted, and the first number of a row is
// bigger than the last number of the row above. Return true if target is in
// the matrix. The solution should run in O(log(rows * columns)) time.
//
// Example 1:
//
//	Input:  matrix = [[1, 2, 4, 8], [10, 11, 12, 13], [14, 20, 30, 40]], target = 10
//	Output: true
//
// Example 2:
//
//	Input:  the same matrix, target = 15
//	Output: false
//
// How to run (from the "local" folder):
//
//	go run ./041-search-2d-matrix
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: every row is sorted and each row starts after the previous row
// ends, so the whole matrix is ONE sorted list written row by row.
// Do a normal binary search on positions 0 .. rows*cols-1 and turn a
// position back into (row, column):
//
//	row = position / cols,  column = position % cols
//
// Time: O(log(rows * cols)).
func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	left := 0
	right := rows*cols - 1

	for left <= right {
		mid := (left + right) / 2
		value := matrix[mid/cols][mid%cols]

		if value == target {
			return true
		}
		if value < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1: found", `searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 10)`,
		searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 10), true)

	// Example 2 from the problem.
	check("example 2: not found", `searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 15)`,
		searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 15), false)

	// One cell.
	check("one cell, found", `searchMatrix([][]int{{1}}, 1)`,
		searchMatrix([][]int{{1}}, 1), true)
	check("one cell, not found", `searchMatrix([][]int{{1}}, 2)`,
		searchMatrix([][]int{{1}}, 2), false)

	// The corners.
	check("first cell", `searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 1)`,
		searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 1), true)
	check("last cell", `searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 40)`,
		searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 40), true)

	// Outside the range.
	check("smaller than all", `searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 0)`,
		searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 0), false)
	check("bigger than all", `searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 41)`,
		searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 41), false)

	// One row and one column.
	check("one row", `searchMatrix([][]int{{1, 3, 5}}, 3)`,
		searchMatrix([][]int{{1, 3, 5}}, 3), true)
	check("one column", `searchMatrix([][]int{{1}, {3}, {5}}, 5)`,
		searchMatrix([][]int{{1}, {3}, {5}}, 5), true)

	// The classic example: 13 is between two rows.
	check("between two rows", `searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13)`,
		searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13), false)

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
