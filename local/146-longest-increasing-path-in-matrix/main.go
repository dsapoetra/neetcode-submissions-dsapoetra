// Problem:    Longest Increasing Path in a Matrix
// Difficulty: Hard
// Link:       https://neetcode.io/problems/longest-increasing-path-in-matrix
// Source:     Data Structures & Algorithms/longest-increasing-path-in-matrix/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a matrix of numbers, return the length of the longest path in which
// every step goes to a neighbor (up, down, left, right) with a strictly
// bigger value.
//
// Example 1:
//
//	Input:  matrix = [[5, 5, 3], [2, 3, 6], [1, 1, 1]]
//	Output: 4             (1 -> 2 -> 3 -> 6)
//
// Example 2:
//
//	Input:  matrix = [[1, 2, 3], [2, 1, 4], [7, 6, 5]]
//	Output: 7             (1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7)
//
// How to run (from the "local" folder):
//
//	go run ./146-longest-increasing-path-in-matrix
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (DFS + remember): the longest path that STARTS at a cell is
// 1 + the longest path starting at a bigger neighbor.
// Many paths pass through the same cells, so we remember the answer of
// every cell (memo) and compute each cell only once.
// We never walk in circles, because we only step to bigger numbers.
// Time: O(rows * cols).
func longestIncreasingPath(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	memo := make([][]int, rows) // 0 = not computed yet
	for i := range memo {
		memo[i] = make([]int, cols)
	}

	// the four neighbors: down, up, right, left
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// longestFrom returns the length of the longest increasing path starting at (r, c).
	var longestFrom func(r int, c int) int
	longestFrom = func(r int, c int) int {
		// already computed: reuse the answer
		if memo[r][c] != 0 {
			return memo[r][c]
		}

		best := 1 // the cell alone
		for _, d := range directions {
			// (nr, nc) is the neighbor in this direction
			nr := r + d[0]
			nc := c + d[1]
			// inside the grid AND bigger than this cell: we may step there
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && matrix[nr][nc] > matrix[r][c] {
				// this cell + the longest path that starts at the neighbor
				length := 1 + longestFrom(nr, nc)
				if length > best {
					best = length
				}
			}
		}

		// remember the answer for this cell
		memo[r][c] = best
		return best
	}

	// the path may start anywhere: try every cell and keep the longest
	answer := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if longestFrom(r, c) > answer {
				answer = longestFrom(r, c)
			}
		}
	}
	return answer
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 1 -> 2 -> 3 -> 6 (or 1 -> 2 -> 3 -> 5).
	check("example 1", `longestIncreasingPath([][]int{{5, 5, 3}, {2, 3, 6}, {1, 1, 1}})`,
		longestIncreasingPath([][]int{{5, 5, 3}, {2, 3, 6}, {1, 1, 1}}), 4)

	// Example 2 from the problem: a spiral from 1 to 7.
	check("example 2", `longestIncreasingPath([][]int{{1, 2, 3}, {2, 1, 4}, {7, 6, 5}})`,
		longestIncreasingPath([][]int{{1, 2, 3}, {2, 1, 4}, {7, 6, 5}}), 7)

	// The classic examples: 1 -> 2 -> 6 -> 9 and 3 -> 4 -> 5 -> 6.
	check("classic 1", `longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}})`,
		longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}), 4)
	check("classic 2", `longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}})`,
		longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}), 4)

	// One cell.
	check("one cell", `longestIncreasingPath([][]int{{1}})`,
		longestIncreasingPath([][]int{{1}}), 1)

	// All equal: the path must be STRICTLY increasing, so no step is possible.
	check("all equal", `longestIncreasingPath([][]int{{7, 7}, {7, 7}})`,
		longestIncreasingPath([][]int{{7, 7}, {7, 7}}), 1)

	// One row, already increasing.
	check("one row", `longestIncreasingPath([][]int{{1, 2, 3, 4}})`,
		longestIncreasingPath([][]int{{1, 2, 3, 4}}), 4)

	// One column, decreasing: walk it from the bottom up.
	check("one column", `longestIncreasingPath([][]int{{4}, {3}, {2}})`,
		longestIncreasingPath([][]int{{4}, {3}, {2}}), 3)

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
