// Problem:    Walls and Gates (Islands and Treasure)
// Difficulty: Medium
// Link:       https://neetcode.io/problems/islands-and-treasure
// Source:     Data Structures & Algorithms/islands-and-treasure/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// A grid has three kinds of cells: -1 is water (cannot be crossed), 0 is a
// treasure chest, and 2147483647 ("INF") is land. Replace every land cell
// with its distance (steps up, down, left, right) to the nearest treasure
// chest. Land that cannot reach a chest stays INF. Change the grid in place.
//
// Example 1:
//
//	Input:  grid = [
//	          [INF,  -1,   0, INF],
//	          [INF, INF, INF,  -1],
//	          [INF,  -1, INF,  -1],
//	          [  0,  -1, INF, INF]
//	        ]
//	Output:        [
//	          [3, -1, 0,  1],
//	          [2,  2, 1, -1],
//	          [1, -1, 2, -1],
//	          [0, -1, 3,  4]
//	        ]
//
// Example 2:
//
//	Input:  grid = [[0, -1], [INF, INF]]
//	Output: [[0, -1], [1, 2]]
//
// How to run (from the "local" folder):
//
//	go run ./081-islands-and-treasure
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (BFS from all treasures at once): put every treasure chest (0) in
// a queue. Then spread out step by step: the land cells next to a chest
// get distance 1, their neighbors distance 2, and so on.
// Because BFS reaches closer cells first, the first time we reach a land
// cell is the shortest distance to ANY chest.
// A land cell that was not reached yet still holds the big "INF" number.
// Time: O(rows * cols).
func islandsAndTreasure(grid [][]int) {
	const inf = 2147483647
	rows := len(grid)
	cols := len(grid[0])

	// start with all treasure chests
	queue := [][2]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		for _, d := range directions {
			r := cell[0] + d[0]
			c := cell[1] + d[1]

			// outside the grid, or not an unvisited land cell (water, chest, or already reached)
			if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != inf {
				continue
			}

			grid[r][c] = grid[cell[0]][cell[1]] + 1 // one step further than the cell we came from
			queue = append(queue, [2]int{r, c})
		}
	}
}

// ===================== Test cases =====================

func main() {
	// inf is the big number that marks a land cell (distance not known yet).
	// -1 is water, 0 is a treasure chest.
	const inf = 2147483647

	// islandsAndTreasure changes the grid in place and returns nothing, so
	// this small shortcut runs it and then gives the grid back.
	filled := func(grid [][]int) [][]int {
		islandsAndTreasure(grid)
		return grid
	}

	// Example 1 from the problem.
	check("example 1", `filled([][]int{{inf, -1, 0, inf}, {inf, inf, inf, -1}, {inf, -1, inf, -1}, {0, -1, inf, inf}})`,
		filled([][]int{
			{inf, -1, 0, inf},
			{inf, inf, inf, -1},
			{inf, -1, inf, -1},
			{0, -1, inf, inf},
		}), [][]int{
			{3, -1, 0, 1},
			{2, 2, 1, -1},
			{1, -1, 2, -1},
			{0, -1, 3, 4},
		})

	// Example 2 from the problem.
	check("example 2", `filled([][]int{{0, -1}, {inf, inf}})`,
		filled([][]int{{0, -1}, {inf, inf}}), [][]int{{0, -1}, {1, 2}})

	// The land cell is cut off by water: it keeps the big number.
	check("land that cannot be reached", `filled([][]int{{inf, -1, 0}})`,
		filled([][]int{{inf, -1, 0}}), [][]int{{inf, -1, 0}})

	// No treasure at all: nothing changes.
	check("no treasure", `filled([][]int{{inf, inf}})`,
		filled([][]int{{inf, inf}}), [][]int{{inf, inf}})

	// Only a treasure chest.
	check("only a chest", `filled([][]int{{0}})`,
		filled([][]int{{0}}), [][]int{{0}})

	// Two chests: every cell gets the distance to the NEAREST one.
	check("two chests", `filled([][]int{{0, inf, inf, inf, 0}})`,
		filled([][]int{{0, inf, inf, inf, 0}}), [][]int{{0, 1, 2, 1, 0}})

	// The way around the water is longer than the straight line.
	check("around the water", `filled([][]int{{0, -1, inf}, {inf, inf, inf}})`,
		filled([][]int{
			{0, -1, inf},
			{inf, inf, inf},
		}), [][]int{
			{0, -1, 4},
			{1, 2, 3},
		})

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
