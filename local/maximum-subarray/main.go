// Problem: Maximum Subarray
// Link:    https://neetcode.io/problems/maximum-subarray
// Source:  Data Structures & Algorithms/maximum-subarray/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./maximum-subarray
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (Kadane's algorithm): keep a running sum. If the running sum
// becomes negative it can only hurt us, so throw it away and start again.
// Time: O(n).
func maxSubArray(nums []int) int {
	best := nums[0] // best sum found so far
	cur := 0        // sum of the subarray we are building now

	for _, n := range nums {
		// a negative running sum is useless, restart from here
		if cur < 0 {
			cur = 0
		}
		cur += n

		if cur > best {
			best = cur
		}
	}
	return best
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: [4,-2,2,1,-1,4] = 8.
	check("example 1", maxSubArray([]int{2, -3, 4, -2, 2, 1, -1, 4}), 8)

	// Example 2 from the problem: one negative number (we must pick at least one).
	check("example 2: one negative", maxSubArray([]int{-1}), -1)

	// One positive number.
	check("one positive", maxSubArray([]int{5}), 5)

	// The classic example: [4,-1,2,1] = 6.
	check("classic", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6)

	// Taking everything is best, even with the -1 inside.
	check("take everything", maxSubArray([]int{5, 4, -1, 7, 8}), 23)

	// All negative: the answer is the biggest (least negative) single number.
	check("all negative", maxSubArray([]int{-3, -2, -5}), -2)

	// All positive: take everything.
	check("all positive", maxSubArray([]int{1, 2, 3}), 6)

	// Only zeros.
	check("zeros", maxSubArray([]int{0, 0}), 0)

	// [2,-1,2] = 3 is better than a single 2.
	check("small dip in the middle", maxSubArray([]int{-1, 2, -1, 2, -1}), 3)

	// A big negative splits the list: [5,-4,20] = 21 on the right side.
	check("big negative splits the list", maxSubArray([]int{8, -19, 5, -4, 20}), 21)

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
