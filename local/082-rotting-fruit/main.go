// Problem:    Rotting Oranges
// Difficulty: Medium
// Link:       https://neetcode.io/problems/rotting-fruit
// Source:     Data Structures & Algorithms/rotting-fruit/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// In a grid, 0 is an empty cell, 1 is a fresh fruit and 2 is a rotten
// fruit. Every minute, each fresh fruit next to a rotten one (up, down,
// left, right) becomes rotten too. Return the number of minutes until no
// fresh fruit is left, or -1 if that never happens.
//
// Example 1:
//
//	Input:  grid = [[1, 1, 0], [0, 1, 1], [0, 1, 2]]
//	Output: 4
//
// Example 2:
//
//	Input:  grid = [[1, 0, 1], [0, 2, 0], [1, 0, 1]]
//	Output: -1
//
// How to run (from the "local" folder):
//
//	go run ./082-rotting-fruit
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (BFS, minute by minute): put all rotten fruits in a queue. In each
// minute, every rotten fruit in the queue rots its fresh neighbors, and
// those neighbors become the queue for the next minute.
// Count the fresh fruits at the start; if some are still fresh when the
// queue is empty, they can never rot, so the answer is -1.
// Time: O(rows * cols).
func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	queue := [][2]int{} // the fruits that became rotten most recently
	fresh := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			} else if grid[r][c] == 1 {
				fresh++
			}
		}
	}

	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	minutes := 0

	for len(queue) > 0 && fresh > 0 {
		minutes++
		next := [][2]int{}

		for _, cell := range queue {
			for _, d := range directions {
				r := cell[0] + d[0]
				c := cell[1] + d[1]

				// inside the grid and fresh: it rots now
				if r >= 0 && r < rows && c >= 0 && c < cols && grid[r][c] == 1 {
					grid[r][c] = 2
					fresh--
					next = append(next, [2]int{r, c})
				}
			}
		}
		queue = next
	}

	if fresh > 0 {
		return -1
	}
	return minutes
}

// ===================== Test cases =====================

func main() {
	// 0 = empty, 1 = fresh, 2 = rotten.

	// Example 1 from the problem.
	check("example 1", `orangesRotting([][]int{{1, 1, 0}, {0, 1, 1}, {0, 1, 2}})`,
		orangesRotting([][]int{{1, 1, 0}, {0, 1, 1}, {0, 1, 2}}), 4)

	// Example 2 from the problem: the corners are never reached.
	check("example 2: impossible", `orangesRotting([][]int{{1, 0, 1}, {0, 2, 0}, {1, 0, 1}})`,
		orangesRotting([][]int{{1, 0, 1}, {0, 2, 0}, {1, 0, 1}}), -1)

	// The classic examples.
	check("classic", `orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})`,
		orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}), 4)
	check("classic impossible", `orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}})`,
		orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}), -1)

	// No fresh fruit at the start: 0 minutes.
	check("no fresh fruit", `orangesRotting([][]int{{0, 2}})`,
		orangesRotting([][]int{{0, 2}}), 0)
	check("empty cell only", `orangesRotting([][]int{{0}})`,
		orangesRotting([][]int{{0}}), 0)
	check("everything rotten already", `orangesRotting([][]int{{2, 2}, {2, 2}})`,
		orangesRotting([][]int{{2, 2}, {2, 2}}), 0)

	// A fresh fruit, but nothing rotten to start with.
	check("nothing rotten", `orangesRotting([][]int{{1}})`,
		orangesRotting([][]int{{1}}), -1)

	// Rot spreads from both ends at the same time.
	check("two rotten fruits", `orangesRotting([][]int{{2, 1, 1, 1, 2}})`,
		orangesRotting([][]int{{2, 1, 1, 1, 2}}), 2)

	// One step.
	check("one step", `orangesRotting([][]int{{2, 1}})`,
		orangesRotting([][]int{{2, 1}}), 1)

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
