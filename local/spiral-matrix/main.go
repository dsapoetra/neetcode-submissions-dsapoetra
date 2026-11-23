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
	check("example 1: 2 x 2", `spiralOrder([][]int{{1, 2}, {3, 4}})`,
		spiralOrder([][]int{{1, 2}, {3, 4}}), []int{1, 2, 4, 3})

	// Example 2 from the problem.
	check("example 2: 3 x 3", `spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})`,
		spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), []int{1, 2, 3, 6, 9, 8, 7, 4, 5})

	// Example 3 from the problem: more columns than rows.
	check("example 3: 3 x 4", `spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})`,
		spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}), []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7})

	// One cell.
	check("1 x 1", `spiralOrder([][]int{{1}})`,
		spiralOrder([][]int{{1}}), []int{1})

	// One row: just left to right.
	check("one row", `spiralOrder([][]int{{1, 2, 3}})`,
		spiralOrder([][]int{{1, 2, 3}}), []int{1, 2, 3})

	// One column: just top to bottom.
	check("one column", `spiralOrder([][]int{{1}, {2}, {3}})`,
		spiralOrder([][]int{{1}, {2}, {3}}), []int{1, 2, 3})

	// 4 x 4: two full rings.
	check("4 x 4", `spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}})`,
		spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}), []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10})

	// More rows than columns.
	check("3 x 2", `spiralOrder([][]int{{1, 2}, {3, 4}, {5, 6}})`,
		spiralOrder([][]int{{1, 2}, {3, 4}, {5, 6}}), []int{1, 2, 4, 6, 5, 3})

	// Two rows.
	check("2 x 3", `spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}})`,
		spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}}), []int{1, 2, 3, 6, 5, 4})

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
