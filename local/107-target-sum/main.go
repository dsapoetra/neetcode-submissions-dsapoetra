// Problem:    Target Sum
// Difficulty: Medium
// Link:       https://neetcode.io/problems/target-sum
// Source:     Data Structures & Algorithms/target-sum/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Put a + or a - in front of every number of the list and add everything
// up. Return the number of different ways to reach exactly the target.
//
// Example 1:
//
//	Input:  nums = [2, 2, 2], target = 2
//	Output: 3             (-2+2+2, +2-2+2, +2+2-2)
//
// How to run (from the "local" folder):
//
//	go run ./107-target-sum
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: go through the numbers one by one and remember every total we
// can have so far, together with HOW MANY ways lead to that total.
// For the next number, each total splits in two: total + num and
// total - num.
// At the end, look up how many ways gave exactly the target.
// Time: O(n * number of different totals).
func findTargetSumWays(nums []int, target int) int {
	ways := map[int]int{0: 1} // before any number: total 0, one way

	for _, num := range nums {
		next := map[int]int{}
		for total, count := range ways {
			next[total+num] += count // put a "+" in front of num
			next[total-num] += count // put a "-" in front of num
		}
		ways = next
	}
	return ways[target]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: -2+2+2, +2-2+2, +2+2-2.
	check("example 1", `findTargetSumWays([]int{2, 2, 2}, 2)`,
		findTargetSumWays([]int{2, 2, 2}, 2), 3)

	// The classic example: one of the five 1s gets the minus.
	check("classic", `findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)`,
		findTargetSumWays([]int{1, 1, 1, 1, 1}, 3), 5)

	// One number.
	check("one number, reachable", `findTargetSumWays([]int{1}, 1)`,
		findTargetSumWays([]int{1}, 1), 1)
	check("one number, not reachable", `findTargetSumWays([]int{1}, 2)`,
		findTargetSumWays([]int{1}, 2), 0)

	// Tricky: +0 and -0 are two different ways.
	check("a zero counts twice", `findTargetSumWays([]int{0}, 0)`,
		findTargetSumWays([]int{0}, 0), 2)
	check("two zeros", `findTargetSumWays([]int{0, 0, 1}, 1)`,
		findTargetSumWays([]int{0, 0, 1}, 1), 4)

	// +1+2-3 and -1-2+3.
	check("target 0", `findTargetSumWays([]int{1, 2, 3}, 0)`,
		findTargetSumWays([]int{1, 2, 3}, 0), 2)

	// Everything plus.
	check("all plus", `findTargetSumWays([]int{1, 2, 3}, 6)`,
		findTargetSumWays([]int{1, 2, 3}, 6), 1)

	// Bigger than the sum of all numbers.
	check("target too big", `findTargetSumWays([]int{1, 2, 3}, 7)`,
		findTargetSumWays([]int{1, 2, 3}, 7), 0)

	// A negative target: everything minus.
	check("negative target", `findTargetSumWays([]int{5, 5}, -10)`,
		findTargetSumWays([]int{5, 5}, -10), 1)

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
