// Problem: Spiral Matrix
// Link:    https://neetcode.io/problems/spiral-matrix
// Source:  Data Structures & Algorithms/spiral-matrix/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./spiral-matrix
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: keep four walls: top, bottom, left, right. Walk around the
// outside (right, down, left, up), and after finishing each side move
// that wall inwards. Repeat until the walls cross.
// Time: O(rows * cols).
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		// go right along the top row
		for c := left; c <= right; c++ {
			res = append(res, matrix[top][c])
		}
		top++

		// go down along the right column
		for r := top; r <= bottom; r++ {
			res = append(res, matrix[r][right])
		}
		right--

		// go left along the bottom row (only if there is still a row left)
		if top <= bottom {
			for c := right; c >= left; c-- {
				res = append(res, matrix[bottom][c])
			}
			bottom--
		}

		// go up along the left column (only if there is still a column left)
		if left <= right {
			for r := bottom; r >= top; r-- {
				res = append(res, matrix[r][left])
			}
			left++
		}
	}
	return res
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1: 2 x 2", spiralOrder([][]int{{1, 2}, {3, 4}}), []int{1, 2, 4, 3})

	// Example 2 from the problem.
	check("example 2: 3 x 3", spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), []int{1, 2, 3, 6, 9, 8, 7, 4, 5})

	// Example 3 from the problem: more columns than rows.
	check("example 3: 3 x 4",
		spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}),
		[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7})

	// One cell.
	check("1 x 1", spiralOrder([][]int{{1}}), []int{1})

	// One row: just left to right.
	check("one row", spiralOrder([][]int{{1, 2, 3}}), []int{1, 2, 3})

	// One column: just top to bottom.
	check("one column", spiralOrder([][]int{{1}, {2}, {3}}), []int{1, 2, 3})

	// 4 x 4: two full rings.
	check("4 x 4",
		spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}),
		[]int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10})

	// More rows than columns.
	check("3 x 2", spiralOrder([][]int{{1, 2}, {3, 4}, {5, 6}}), []int{1, 2, 4, 6, 5, 3})

	// Two rows.
	check("2 x 3", spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}}), []int{1, 2, 3, 6, 5, 4})

	report()
}

// ===================== Helpers for the tests =====================

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
