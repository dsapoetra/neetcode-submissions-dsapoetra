// Problem: Unique Paths
// Link:    https://neetcode.io/problems/count-paths
// Source:  Data Structures & Algorithms/count-paths/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./count-paths
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (dynamic programming): we can only move right or down, so
// the number of ways to reach a cell = ways from above + ways from the left.
// The first row and first column have only 1 way each.
// Time: O(m * n).
func uniquePaths(m int, n int) int {
	// ways[r][c] = number of paths to reach cell (r, c)
	ways := make([][]int, m)
	for r := 0; r < m; r++ {
		ways[r] = make([]int, n)
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if r == 0 || c == 0 {
				ways[r][c] = 1 // top row or left column: only one way
			} else {
				ways[r][c] = ways[r-1][c] + ways[r][c-1]
			}
		}
	}
	return ways[m-1][n-1]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1: 3 x 6", `uniquePaths(3, 6)`,
		uniquePaths(3, 6), 21)

	// Example 2 from the problem.
	check("example 2: 3 x 3", `uniquePaths(3, 3)`,
		uniquePaths(3, 3), 6)

	// A 1 x 1 grid: we are already there, that counts as one path.
	check("1 x 1", `uniquePaths(1, 1)`,
		uniquePaths(1, 1), 1)

	// One row: we can only go right.
	check("one row", `uniquePaths(1, 10)`,
		uniquePaths(1, 10), 1)

	// One column: we can only go down.
	check("one column", `uniquePaths(10, 1)`,
		uniquePaths(10, 1), 1)

	// 2 x 2: right-down or down-right.
	check("2 x 2", `uniquePaths(2, 2)`,
		uniquePaths(2, 2), 2)

	// The classic example.
	check("3 x 7", `uniquePaths(3, 7)`,
		uniquePaths(3, 7), 28)

	// Swapping rows and columns gives the same answer.
	check("7 x 3", `uniquePaths(7, 3)`,
		uniquePaths(7, 3), 28)

	// A bigger grid.
	check("10 x 10", `uniquePaths(10, 10)`,
		uniquePaths(10, 10), 48620)

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
