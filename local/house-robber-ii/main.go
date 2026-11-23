// Problem: House Robber II
// Link:    https://neetcode.io/problems/house-robber-ii
// Source:  Data Structures & Algorithms/house-robber-ii/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./house-robber-ii
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: the houses are in a circle, so the first and last house are
// neighbors and we cannot rob both. So solve the normal (straight line)
// House Robber twice:
//  1. without the last house
//  2. without the first house
//
// and take the better result.
// Time: O(n).
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	withoutLast := robLine(nums[:len(nums)-1])
	withoutFirst := robLine(nums[1:])
	if withoutLast > withoutFirst {
		return withoutLast
	}
	return withoutFirst
}

// robLine solves House Robber for houses in a straight line.
func robLine(nums []int) int {
	twoBack := 0 // best money up to two houses back
	oneBack := 0 // best money up to the previous house
	for _, money := range nums {
		best := oneBack // skip this house
		if twoBack+money > best {
			best = twoBack + money // rob this house
		}
		twoBack = oneBack
		oneBack = best
	}
	return oneBack
}

// ===================== Test cases =====================

func main() {
	// The houses are in a circle: the first and the last house are neighbors.

	// Example 1 from the problem: 3 + 3 is not allowed (first and last), so take 4.
	check("example 1", `rob([]int{3, 4, 3})`,
		rob([]int{3, 4, 3}), 4)

	// Example 2 from the problem: 9 + 6.
	check("example 2", `rob([]int{2, 9, 8, 3, 6})`,
		rob([]int{2, 9, 8, 3, 6}), 15)

	// One house: rob it.
	check("one house", `rob([]int{5})`,
		rob([]int{5}), 5)

	// Two houses are neighbors: rob the richer one.
	check("two houses", `rob([]int{1, 2})`,
		rob([]int{1, 2}), 2)

	// The classic example.
	check("classic 2,3,2", `rob([]int{2, 3, 2})`,
		rob([]int{2, 3, 2}), 3)

	// 1 + 3.
	check("classic 1,2,3,1", `rob([]int{1, 2, 3, 1})`,
		rob([]int{1, 2, 3, 1}), 4)

	// Nothing to steal.
	check("all zeros", `rob([]int{0, 0, 0})`,
		rob([]int{0, 0, 0}), 0)

	// Three houses: every pair is neighbors, so only one can be robbed.
	check("three houses", `rob([]int{1, 2, 3})`,
		rob([]int{1, 2, 3}), 3)

	// 200 + 140. The 10 at the end cannot be added (neighbor of the 200).
	check("bigger case", `rob([]int{200, 3, 140, 20, 10})`,
		rob([]int{200, 3, 140, 20, 10}), 340)

	// Tricky: in a straight line 9 + 9 = 18 would work, in a circle it does not.
	check("rich first and last house", `rob([]int{9, 1, 1, 9})`,
		rob([]int{9, 1, 1, 9}), 10)

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
