// Problem: Jump Game
// Link:    https://neetcode.io/problems/jump-game
// Source:  Data Structures & Algorithms/jump-game/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./jump-game
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (greedy): keep track of the farthest index we can reach so far.
// Walk through the array; if we are standing on an index that is beyond
// the farthest reach, we are stuck.
// Time: O(n).
func canJump(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums); i++ {
		// we cannot even get to index i
		if i > farthest {
			return false
		}
		// from i we can reach up to i + nums[i]
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
	}
	return true
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 0 -> 1 -> 3 -> 4.
	check("example 1: reachable", canJump([]int{1, 2, 0, 1, 0}), true)

	// Example 2 from the problem: we always get stuck on the 0 at index 3.
	check("example 2: stuck on a zero", canJump([]int{1, 2, 1, 0, 1}), false)

	// One element: we are already at the end.
	check("one element", canJump([]int{0}), true)

	// One big jump over the zeros.
	check("jump over zeros", canJump([]int{2, 0, 0}), true)

	// We cannot even leave the start.
	check("zero at the start", canJump([]int{0, 1}), false)

	// The classic "false" example: every path ends on the 0.
	check("classic false", canJump([]int{3, 2, 1, 0, 4}), false)

	// The classic "true" example.
	check("classic true", canJump([]int{2, 3, 1, 1, 4}), true)

	// One step at a time.
	check("all ones", canJump([]int{1, 1, 1, 1}), true)

	// The first jump reaches the end directly.
	check("one huge jump", canJump([]int{5, 0, 0, 0, 0, 0}), true)

	// A zero right in the way.
	check("zero in the way", canJump([]int{1, 0, 1}), false)

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
