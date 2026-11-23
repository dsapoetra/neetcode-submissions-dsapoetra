// Problem:    Jump Game II
// Difficulty: Medium
// Link:       https://neetcode.io/problems/jump-game-ii
// Source:     Data Structures & Algorithms/jump-game-ii/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// You start at index 0. nums[i] is the longest jump you can make from
// index i. Return the smallest number of jumps needed to reach the last
// index (it is always reachable).
//
// Example 1:
//
//	Input:  nums = [2, 4, 1, 1, 1, 1]
//	Output: 2             (index 0 -> 1 -> 5)
//
// Example 2:
//
//	Input:  nums = [2, 1, 2, 1, 0]
//	Output: 2
//
// How to run (from the "local" folder):
//
//	go run ./112-jump-game-ii
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (greedy): think in "rounds". With 1 jump we can reach some range of
// positions, with 2 jumps a further range, and so on.
//
//	currentEnd = the last position we can reach with the jumps made so far
//	farthest   = the farthest position we could reach with one more jump
//
// Walk through the list. When we arrive at currentEnd we MUST jump again,
// and the new range ends at farthest.
// Time: O(n).
func jump(nums []int) int {
	jumps := 0
	currentEnd := 0
	farthest := 0

	// we stop before the last index: standing there needs no more jumps
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		if i == currentEnd {
			jumps++ // one more jump
			currentEnd = farthest
		}
	}
	return jumps
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 0 -> 1 -> 5.
	check("example 1", `jump([]int{2, 4, 1, 1, 1, 1})`,
		jump([]int{2, 4, 1, 1, 1, 1}), 2)

	// Example 2 from the problem.
	check("example 2", `jump([]int{2, 1, 2, 1, 0})`,
		jump([]int{2, 1, 2, 1, 0}), 2)

	// The classic examples.
	check("classic 1", `jump([]int{2, 3, 1, 1, 4})`,
		jump([]int{2, 3, 1, 1, 4}), 2)
	check("classic 2", `jump([]int{2, 3, 0, 1, 4})`,
		jump([]int{2, 3, 0, 1, 4}), 2)

	// One element: we are already at the end.
	check("one element", `jump([]int{0})`,
		jump([]int{0}), 0)

	// Two elements.
	check("two elements", `jump([]int{1, 2})`,
		jump([]int{1, 2}), 1)

	// Only steps of 1.
	check("all ones", `jump([]int{1, 1, 1, 1})`,
		jump([]int{1, 1, 1, 1}), 3)

	// One big jump.
	check("one big jump", `jump([]int{5, 1, 1, 1, 1})`,
		jump([]int{5, 1, 1, 1, 1}), 1)

	// A jump of exactly the right length.
	check("exact jump", `jump([]int{3, 2, 1, 0})`,
		jump([]int{3, 2, 1, 0}), 1)

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
