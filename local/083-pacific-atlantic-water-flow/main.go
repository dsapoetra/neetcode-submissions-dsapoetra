// Problem:    Pacific Atlantic Water Flow
// Difficulty: Medium
// Link:       https://neetcode.io/problems/pacific-atlantic-water-flow
// Source:     Data Structures & Algorithms/pacific-atlantic-water-flow/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// heights[r][c] is the height of a cell on an island. The Pacific ocean
// touches the top and left edge, the Atlantic the bottom and right edge.
// Water flows from a cell to a neighbor (up, down, left, right) that is
// lower or equally high. Return all cells from which water can reach BOTH oceans.
//
// Example 1:
//
//	Input:  heights = [
//	          [4, 2, 7, 3, 4],
//	          [7, 4, 6, 4, 7],
//	          [6, 3, 5, 3, 6]
//	        ]
//	Output: [[0,2], [0,4], [1,0], [1,1], [1,2], [1,3], [1,4], [2,0]]
//
// Example 2:
//
//	Input:  heights = [[1], [1]]
//	Output: [[0,0], [1,0]]
//
// How to run (from the "local" folder):
//
//	go run ./083-pacific-atlantic-water-flow
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: instead of asking "can water from this cell reach the ocean?",
// go backwards: start AT each ocean and climb uphill (to cells that are
// the same height or higher). Every cell we can climb to can send its
// water down to that ocean.
// Do this once for the Pacific (top + left edges) and once for the
// Atlantic (bottom + right edges). Cells reached by both are the answer.
// Time: O(rows * cols).
func pacificAtlantic(heights [][]int) [][]int {
	rows := len(heights)
	cols := len(heights[0])

	pacific := make([][]bool, rows)  // pacific[r][c] = can reach the Pacific
	atlantic := make([][]bool, rows) // atlantic[r][c] = can reach the Atlantic
	for r := 0; r < rows; r++ {
		pacific[r] = make([]bool, cols)
		atlantic[r] = make([]bool, cols)
	}

	// left edge touches the Pacific, right edge touches the Atlantic
	for r := 0; r < rows; r++ {
		climb(heights, r, 0, pacific)
		climb(heights, r, cols-1, atlantic)
	}
	// top edge touches the Pacific, bottom edge touches the Atlantic
	for c := 0; c < cols; c++ {
		climb(heights, 0, c, pacific)
		climb(heights, rows-1, c, atlantic)
	}

	// keep the cells that reach both oceans
	res := [][]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pacific[r][c] && atlantic[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}

// climb marks (r, c) as reached, then moves to every neighbor that is
// the same height or higher (DFS).
func climb(heights [][]int, r int, c int, reached [][]bool) {
	reached[r][c] = true

	// the 4 directions: down, up, right, left
	moves := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, move := range moves {
		nr := r + move[0]
		nc := c + move[1]

		// skip neighbors outside the grid
		if nr < 0 || nc < 0 || nr >= len(heights) || nc >= len(heights[0]) {
			continue
		}
		// skip neighbors we already visited or that are lower
		if reached[nr][nc] || heights[nr][nc] < heights[r][c] {
			continue
		}
		climb(heights, nr, nc, reached)
	}
}

// ===================== Test cases =====================

func main() {
	// The Pacific touches the top and left edge, the Atlantic the bottom and
	// right edge. The answer lists [row, column] from top-left to bottom-right.

	// Example 1 from the problem.
	check("example 1", `pacificAtlantic([][]int{{4, 2, 7, 3, 4}, {7, 4, 6, 4, 7}, {6, 3, 5, 3, 6}})`,
		pacificAtlantic([][]int{
			{4, 2, 7, 3, 4},
			{7, 4, 6, 4, 7},
			{6, 3, 5, 3, 6},
		}), [][]int{{0, 2}, {0, 4}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {2, 0}})

	// Example 2 from the problem: one column, every cell touches both oceans' sides.
	check("example 2: one column", `pacificAtlantic([][]int{{1}, {1}})`,
		pacificAtlantic([][]int{{1}, {1}}), [][]int{{0, 0}, {1, 0}})

	// A single cell touches every edge.
	check("single cell", `pacificAtlantic([][]int{{5}})`,
		pacificAtlantic([][]int{{5}}), [][]int{{0, 0}})

	// The classic example.
	check("classic", `pacificAtlantic([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}})`,
		pacificAtlantic([][]int{
			{1, 2, 2, 3, 5},
			{3, 2, 3, 4, 4},
			{2, 4, 5, 3, 1},
			{6, 7, 1, 4, 5},
			{5, 1, 1, 2, 4},
		}), [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}})

	// Everything is flat: water can flow everywhere.
	check("flat", `pacificAtlantic([][]int{{1, 1}, {1, 1}})`,
		pacificAtlantic([][]int{{1, 1}, {1, 1}}), [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}})

	// One row: every cell touches the top (Pacific) and the bottom (Atlantic).
	check("one row", `pacificAtlantic([][]int{{1, 2, 3}})`,
		pacificAtlantic([][]int{{1, 2, 3}}), [][]int{{0, 0}, {0, 1}, {0, 2}})

	// The top-left cell (1) is the lowest: its water cannot reach the Atlantic.
	check("low corner", `pacificAtlantic([][]int{{1, 2}, {4, 3}})`,
		pacificAtlantic([][]int{{1, 2}, {4, 3}}), [][]int{{0, 1}, {1, 0}, {1, 1}})

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
