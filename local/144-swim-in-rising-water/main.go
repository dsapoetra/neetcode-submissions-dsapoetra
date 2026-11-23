// Problem:    Swim in Rising Water
// Difficulty: Hard
// Link:       https://neetcode.io/problems/swim-in-rising-water
// Source:     Data Structures & Algorithms/swim-in-rising-water/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// grid[r][c] is the height of a cell in an n x n grid. It rains: at time t
// the water level is t. You can swim between neighboring cells (up, down,
// left, right) when both cells are at most as high as the water level.
// Return the earliest time at which you can get from the top-left cell to
// the bottom-right cell.
//
// Example 1:
//
//	Input:  grid = [[0, 1], [2, 3]]
//	Output: 3
//
// Example 2:
//
//	Input:  grid = [
//	          [ 0,  1,  2, 10],
//	          [ 9, 14,  4, 13],
//	          [12,  3,  8, 15],
//	          [11,  5,  7,  6]
//	        ]
//	Output: 8
//
// How to run (from the "local" folder):
//
//	go run ./144-swim-in-rising-water
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: at time t we can stand on every cell with height <= t.
// If we can reach the end at time t, we can also reach it at any later
// time. So the times look like: no, no, no, YES, YES, YES ...
// and we can binary search the first YES.
// For one time t, a simple DFS tells us if the end is reachable.
// Time: O(n^2 * log(n^2)).
func swimInWater(grid [][]int) int {
	n := len(grid)

	// the answer is somewhere between time 0 and time n*n-1
	left := 0
	right := n*n - 1 // the heights are 0 .. n*n-1, so this time always works

	for left < right {
		// try the time in the middle
		mid := (left + right) / 2

		// a fresh "visited" grid for every try
		visited := make([][]bool, n)
		for i := range visited {
			visited[i] = make([]bool, n)
		}

		if canSwim(grid, visited, 0, 0, mid) {
			right = mid // works: maybe an earlier time works too
		} else {
			left = mid + 1 // too early
		}
	}
	// left == right: the first time that works
	return left
}

// canSwim is true if the bottom-right cell can be reached from (r, c)
// using only cells with height <= time.
func canSwim(grid [][]int, visited [][]bool, r int, c int, time int) bool {
	n := len(grid)
	// stop: outside the grid, already visited, or the water is too low
	// to stand on this cell at this time
	if r < 0 || r >= n || c < 0 || c >= n || visited[r][c] || grid[r][c] > time {
		return false
	}
	// reached the bottom-right cell
	if r == n-1 && c == n-1 {
		return true
	}
	// mark it, so we do not walk in circles
	visited[r][c] = true

	// try all four directions; one working way is enough
	return canSwim(grid, visited, r+1, c, time) ||
		canSwim(grid, visited, r-1, c, time) ||
		canSwim(grid, visited, r, c+1, time) ||
		canSwim(grid, visited, r, c-1, time)
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: the end cell itself has height 3.
	check("example 1", `swimInWater([][]int{{0, 1}, {2, 3}})`,
		swimInWater([][]int{{0, 1}, {2, 3}}), 3)

	// Example 2 from the problem: the path 0,1,2,4,8,7,6 never goes above 8.
	check("example 2", `swimInWater([][]int{{0, 1, 2, 10}, {9, 14, 4, 13}, {12, 3, 8, 15}, {11, 5, 7, 6}})`,
		swimInWater([][]int{
			{0, 1, 2, 10},
			{9, 14, 4, 13},
			{12, 3, 8, 15},
			{11, 5, 7, 6},
		}), 8)

	// One cell: we are already there.
	check("one cell", `swimInWater([][]int{{0}})`,
		swimInWater([][]int{{0}}), 0)

	// The classic example: the only low path is a long spiral.
	check("classic spiral", `swimInWater([][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}})`,
		swimInWater([][]int{
			{0, 1, 2, 3, 4},
			{24, 23, 22, 21, 5},
			{12, 13, 14, 15, 16},
			{11, 17, 18, 19, 20},
			{10, 9, 8, 7, 6},
		}), 16)

	// The START cell is the highest: we must wait for it.
	check("high start cell", `swimInWater([][]int{{3, 2}, {0, 1}})`,
		swimInWater([][]int{{3, 2}, {0, 1}}), 3)

	// Going down first (0 -> 1 -> 2) is better than going right (0 -> 3 -> 2).
	check("choose the lower way", `swimInWater([][]int{{0, 3}, {1, 2}})`,
		swimInWater([][]int{{0, 3}, {1, 2}}), 2)

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
