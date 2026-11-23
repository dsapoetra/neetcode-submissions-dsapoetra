// Problem:    Rotate Image
// Difficulty: Medium
// Link:       https://neetcode.io/problems/rotate-matrix
// Source:     Data Structures & Algorithms/rotate-matrix/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a square matrix, rotate it by 90 degrees clockwise. The matrix must
// be changed in place (no second matrix).
//
// Example 1:
//
//	Input:  matrix = [[1, 2],
//	                  [3, 4]]
//	Output:          [[3, 1],
//	                  [4, 2]]
//
// Example 2:
//
//	Input:  matrix = [[1, 2, 3],
//	                  [4, 5, 6],
//	                  [7, 8, 9]]
//	Output:          [[7, 4, 1],
//	                  [8, 5, 2],
//	                  [9, 6, 3]]
//
// How to run (from the "local" folder):
//
//	go run ./124-rotate-matrix
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: rotating 90 degrees clockwise is the same as two easy steps:
//  1. transpose: swap matrix[r][c] with matrix[c][r] (flip over the diagonal)
//  2. reverse every row
//
// Example:  1 2 3      1 4 7      7 4 1
//
//	4 5 6  ->  2 5 8  ->  8 5 2
//	7 8 9      3 6 9      9 6 3
//
// Time: O(n^2), no extra matrix needed.
func rotate(matrix [][]int) {
	n := len(matrix)

	// step 1: transpose (only the part above the diagonal, or we swap twice)
	for r := 0; r < n; r++ {
		for c := r + 1; c < n; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}

	// step 2: reverse every row
	for r := 0; r < n; r++ {
		left, right := 0, n-1
		for left < right {
			matrix[r][left], matrix[r][right] = matrix[r][right], matrix[r][left]
			left++
			right--
		}
	}
}

// ===================== Test cases =====================

func main() {
	// rotate changes the matrix in place and returns nothing, so this small
	// shortcut rotates the matrix and then gives it back.
	rotated := func(matrix [][]int) [][]int {
		rotate(matrix)
		return matrix
	}

	// Example 1 from the problem.
	check("example 1: 2 x 2", `rotated([][]int{{1, 2}, {3, 4}})`,
		rotated([][]int{{1, 2}, {3, 4}}), [][]int{{3, 1}, {4, 2}})

	// Example 2 from the problem.
	check("example 2: 3 x 3", `rotated([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})`,
		rotated([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}})

	// One cell: nothing to do.
	check("1 x 1", `rotated([][]int{{1}})`,
		rotated([][]int{{1}}), [][]int{{1}})

	// The classic 4 x 4 example.
	check("4 x 4", `rotated([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}})`,
		rotated([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}), [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}})

	// All the same value: looks the same after rotating.
	check("all the same", `rotated([][]int{{0, 0}, {0, 0}})`,
		rotated([][]int{{0, 0}, {0, 0}}), [][]int{{0, 0}, {0, 0}})

	// Negative values.
	check("negative values", `rotated([][]int{{-1, -2}, {-3, -4}})`,
		rotated([][]int{{-1, -2}, {-3, -4}}), [][]int{{-3, -1}, {-4, -2}})

	// Rotating four times brings the matrix back to the start.
	check("four rotations", `rotated(rotated(rotated(rotated([][]int{{1, 2}, {3, 4}}))))`,
		rotated(rotated(rotated(rotated([][]int{{1, 2}, {3, 4}})))), [][]int{{1, 2}, {3, 4}})

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
