// Problem: House Robber
// Link:    https://neetcode.io/problems/house-robber
// Source:  Data Structures & Algorithms/house-robber/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./house-robber
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: at each house we have two choices:
//   - skip it  -> money stays the same as the previous house
//   - rob it   -> money = this house + best up to two houses back
//
// Take the bigger one. We only need the last two answers.
// Time: O(n).
func rob(nums []int) int {
	twoBack := 0 // best money up to two houses back
	oneBack := 0 // best money up to the previous house

	for _, money := range nums {
		best := oneBack // choice 1: skip this house
		if twoBack+money > best {
			best = twoBack + money // choice 2: rob this house
		}
		// slide forward by one house
		twoBack = oneBack
		oneBack = best
	}
	return oneBack
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: houses 0 and 2 (1 + 3).
	check("example 1", rob([]int{1, 1, 3, 3}), 4)

	// Example 2 from the problem: 2 + 8 + 6.
	check("example 2", rob([]int{2, 9, 8, 3, 6}), 16)

	// One house: rob it.
	check("one house", rob([]int{5}), 5)

	// Two houses: rob the richer one.
	check("two houses", rob([]int{1, 2}), 2)

	// The classic example: 2 + 9 + 1.
	check("classic", rob([]int{2, 7, 9, 3, 1}), 12)

	// 1 + 3.
	check("1,2,3,1", rob([]int{1, 2, 3, 1}), 4)

	// Nothing to steal.
	check("all zeros", rob([]int{0, 0, 0}), 0)

	// Tricky: the best plan skips TWO houses in a row (first and last).
	check("skip two in a row", rob([]int{2, 1, 1, 2}), 4)

	// The two rich houses are 3 apart.
	check("rich houses far apart", rob([]int{10, 1, 1, 10, 1}), 20)

	// No houses at all.
	check("no houses", rob([]int{}), 0)

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
