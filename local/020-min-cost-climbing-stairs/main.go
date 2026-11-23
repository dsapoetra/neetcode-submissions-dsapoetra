// Problem:    Min Cost Climbing Stairs
// Difficulty: Easy
// Link:       https://neetcode.io/problems/min-cost-climbing-stairs
// Source:     Data Structures & Algorithms/min-cost-climbing-stairs/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// cost[i] is the price you pay when you step on stair i. After paying you
// can go up one or two stairs. You may start on stair 0 or stair 1. Return
// the cheapest way to get past the last stair.
//
// Example 1:
//
//	Input:  cost = [1, 2, 3]
//	Output: 2             (start on stair 1, pay 2, jump two stairs)
//
// Example 2:
//
//	Input:  cost = [1, 2, 1, 2, 1, 1, 1]
//	Output: 4
//
// How to run (from the "local" folder):
//
//	go run ./020-min-cost-climbing-stairs
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (dynamic programming): reach[i] = the cheapest total cost to stand
// on step i (before paying for step i itself).
// We may start on step 0 or step 1 for free.
// To arrive at step i we come from step i-1 or step i-2, and we pay the
// cost of the step we leave.
// The "top" is one step past the last one, so the answer is reach[n].
// Time: O(n).
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	reach := make([]int, n+1) // reach[0] = reach[1] = 0: free starting points

	for i := 2; i <= n; i++ {
		fromOneBelow := reach[i-1] + cost[i-1]
		fromTwoBelow := reach[i-2] + cost[i-2]

		if fromOneBelow < fromTwoBelow {
			reach[i] = fromOneBelow
		} else {
			reach[i] = fromTwoBelow
		}
	}
	return reach[n]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: start on step 1 (cost 2) and jump over step 2.
	check("example 1", `minCostClimbingStairs([]int{1, 2, 3})`,
		minCostClimbingStairs([]int{1, 2, 3}), 2)

	// Example 2 from the problem: steps 0, 2, 4, 6 cost 1 each.
	check("example 2", `minCostClimbingStairs([]int{1, 2, 1, 2, 1, 1, 1})`,
		minCostClimbingStairs([]int{1, 2, 1, 2, 1, 1, 1}), 4)

	// The classic examples.
	check("classic 1", `minCostClimbingStairs([]int{10, 15, 20})`,
		minCostClimbingStairs([]int{10, 15, 20}), 15)
	check("classic 2", `minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})`,
		minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}), 6)

	// Two steps: pay for the cheaper one.
	check("two equal steps", `minCostClimbingStairs([]int{5, 5})`,
		minCostClimbingStairs([]int{5, 5}), 5)
	check("first is cheaper", `minCostClimbingStairs([]int{1, 100})`,
		minCostClimbingStairs([]int{1, 100}), 1)
	check("second is cheaper", `minCostClimbingStairs([]int{100, 1})`,
		minCostClimbingStairs([]int{100, 1}), 1)

	// Free steps.
	check("free steps", `minCostClimbingStairs([]int{0, 0})`,
		minCostClimbingStairs([]int{0, 0}), 0)

	// Step 0 (free) and then step 2.
	check("use the free step", `minCostClimbingStairs([]int{0, 1, 2, 2})`,
		minCostClimbingStairs([]int{0, 1, 2, 2}), 2)

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
