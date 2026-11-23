// Problem: Number of Islands
// Link:    https://neetcode.io/problems/count-number-of-islands
// Source:  Data Structures & Algorithms/count-number-of-islands/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./count-number-of-islands
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: scan every cell. When we find land ('1') it is a new island:
// count it, then use DFS to turn the whole island into water ('0')
// so we never count it again.
// Time: O(rows * cols).
func numIslands(grid [][]byte) int {
	count := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				count++
				sinkIsland(grid, r, c)
			}
		}
	}
	return count
}

// sinkIsland turns the land at (r, c) and all land connected to it into water.
func sinkIsland(grid [][]byte, r int, c int) {
	// stop if we are outside the grid or this cell is not land
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] != '1' {
		return
	}

	grid[r][c] = '0' // sink this cell

	// continue in the 4 directions
	sinkIsland(grid, r+1, c)
	sinkIsland(grid, r-1, c)
	sinkIsland(grid, r, c+1)
	sinkIsland(grid, r, c-1)
}

// ===================== Test cases =====================

func main() {
	// byteGrid("011", "000") builds the grid: each string is one row.

	// Example 1 from the problem: all the land is connected.
	check("example 1", `numIslands(byteGrid("01110", "01010", "11010", "00000"))`,
		numIslands(byteGrid("01110", "01010", "11010", "00000")), 1)

	// Example 2 from the problem.
	check("example 2", `numIslands(byteGrid("11001", "11001", "00100", "00011"))`,
		numIslands(byteGrid("11001", "11001", "00100", "00011")), 4)

	// Only water.
	check("only water", `numIslands(byteGrid("000", "000"))`,
		numIslands(byteGrid("000", "000")), 0)

	// Only land: one big island.
	check("only land", `numIslands(byteGrid("11", "11"))`,
		numIslands(byteGrid("11", "11")), 1)

	// The smallest grids.
	check("single land cell", `numIslands(byteGrid("1"))`,
		numIslands(byteGrid("1")), 1)
	check("single water cell", `numIslands(byteGrid("0"))`,
		numIslands(byteGrid("0")), 0)

	// Land that only touches at the corners is NOT connected.
	check("diagonal is not connected", `numIslands(byteGrid("100", "010", "001"))`,
		numIslands(byteGrid("100", "010", "001")), 3)

	// A grid with one row.
	check("one row", `numIslands(byteGrid("10101"))`,
		numIslands(byteGrid("10101")), 3)

	// A snake shape is still one island.
	check("snake shape", `numIslands(byteGrid("111", "001", "111"))`,
		numIslands(byteGrid("111", "001", "111")), 1)

	// A ring of land around water is one island.
	check("ring", `numIslands(byteGrid("111", "101", "111"))`,
		numIslands(byteGrid("111", "101", "111")), 1)

	report()
}

// ===================== Helpers for the tests =====================

// byteGrid turns rows written as strings into a [][]byte grid.
// Example: byteGrid("110", "001") is a grid with 2 rows and 3 columns.
func byteGrid(rows ...string) [][]byte {
	grid := [][]byte{}
	for _, row := range rows {
		grid = append(grid, []byte(row))
	}
	return grid
}

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
