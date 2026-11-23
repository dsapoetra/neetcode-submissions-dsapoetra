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
	check("example 1: reachable", `canJump([]int{1, 2, 0, 1, 0})`,
		canJump([]int{1, 2, 0, 1, 0}), true)

	// Example 2 from the problem: we always get stuck on the 0 at index 3.
	check("example 2: stuck on a zero", `canJump([]int{1, 2, 1, 0, 1})`,
		canJump([]int{1, 2, 1, 0, 1}), false)

	// One element: we are already at the end.
	check("one element", `canJump([]int{0})`,
		canJump([]int{0}), true)

	// One big jump over the zeros.
	check("jump over zeros", `canJump([]int{2, 0, 0})`,
		canJump([]int{2, 0, 0}), true)

	// We cannot even leave the start.
	check("zero at the start", `canJump([]int{0, 1})`,
		canJump([]int{0, 1}), false)

	// The classic "false" example: every path ends on the 0.
	check("classic false", `canJump([]int{3, 2, 1, 0, 4})`,
		canJump([]int{3, 2, 1, 0, 4}), false)

	// The classic "true" example.
	check("classic true", `canJump([]int{2, 3, 1, 1, 4})`,
		canJump([]int{2, 3, 1, 1, 4}), true)

	// One step at a time.
	check("all ones", `canJump([]int{1, 1, 1, 1})`,
		canJump([]int{1, 1, 1, 1}), true)

	// The first jump reaches the end directly.
	check("one huge jump", `canJump([]int{5, 0, 0, 0, 0, 0})`,
		canJump([]int{5, 0, 0, 0, 0, 0}), true)

	// A zero right in the way.
	check("zero in the way", `canJump([]int{1, 0, 1})`,
		canJump([]int{1, 0, 1}), false)

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
