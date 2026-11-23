// Problem:    Max Area of Island
// Difficulty: Medium
// Link:       https://neetcode.io/problems/max-area-of-island
// Source:     Data Structures & Algorithms/max-area-of-island/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a grid of 1 (land) and 0 (water), return the area (number of cells)
// of the biggest island. An island is land connected up, down, left or
// right. Return 0 if there is no land.
//
// Example 1:
//
//	Input:  grid = [
//	          [0, 1, 1, 0, 1],
//	          [1, 0, 1, 0, 1],
//	          [0, 1, 1, 0, 1],
//	          [0, 1, 0, 0, 1]
//	        ]
//	Output: 6
//
// How to run (from the "local" folder):
//
//	go run ./079-max-area-of-island
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: look at every cell. When we find land, walk over the whole island
// (DFS: up, down, left, right) and count its cells. We set every visited
// cell to 0 ("sink" it) so it is never counted twice.
// Time: O(rows * cols).
func maxAreaOfIsland(grid [][]int) int {
	best := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				area := islandArea(grid, r, c)
				if area > best {
					best = area
				}
			}
		}
	}
	return best
}

// islandArea counts the land cells connected to (r, c) and sinks them.
func islandArea(grid [][]int, r int, c int) int {
	// outside the grid, or water
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == 0 {
		return 0
	}

	grid[r][c] = 0 // sink this cell so we do not count it again

	return 1 +
		islandArea(grid, r+1, c) +
		islandArea(grid, r-1, c) +
		islandArea(grid, r, c+1) +
		islandArea(grid, r, c-1)
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: the island in the middle has 6 cells.
	check("example 1", `maxAreaOfIsland([][]int{{0, 1, 1, 0, 1}, {1, 0, 1, 0, 1}, {0, 1, 1, 0, 1}, {0, 1, 0, 0, 1}})`,
		maxAreaOfIsland([][]int{
			{0, 1, 1, 0, 1},
			{1, 0, 1, 0, 1},
			{0, 1, 1, 0, 1},
			{0, 1, 0, 0, 1},
		}), 6)

	// Only water.
	check("only water", `maxAreaOfIsland([][]int{{0, 0}, {0, 0}})`,
		maxAreaOfIsland([][]int{{0, 0}, {0, 0}}), 0)

	// Only land.
	check("only land", `maxAreaOfIsland([][]int{{1, 1}, {1, 1}})`,
		maxAreaOfIsland([][]int{{1, 1}, {1, 1}}), 4)

	// The smallest grids.
	check("single land cell", `maxAreaOfIsland([][]int{{1}})`,
		maxAreaOfIsland([][]int{{1}}), 1)
	check("single water cell", `maxAreaOfIsland([][]int{{0}})`,
		maxAreaOfIsland([][]int{{0}}), 0)

	// Cells that only touch at the corner are separate islands.
	check("diagonal is not connected", `maxAreaOfIsland([][]int{{1, 0}, {0, 1}})`,
		maxAreaOfIsland([][]int{{1, 0}, {0, 1}}), 1)

	// Two islands: a small one (2 cells) and a bigger one (4 cells).
	check("two islands", `maxAreaOfIsland([][]int{{1, 1, 0, 0}, {0, 0, 0, 1}, {0, 1, 1, 1}})`,
		maxAreaOfIsland([][]int{
			{1, 1, 0, 0},
			{0, 0, 0, 1},
			{0, 1, 1, 1},
		}), 4)

	// One row.
	check("one row", `maxAreaOfIsland([][]int{{1, 1, 0, 1, 1, 1}})`,
		maxAreaOfIsland([][]int{{1, 1, 0, 1, 1, 1}}), 3)

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
