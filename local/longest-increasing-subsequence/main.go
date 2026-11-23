// Problem: Longest Increasing Subsequence
// Link:    https://neetcode.io/problems/longest-increasing-subsequence
// Source:  Data Structures & Algorithms/longest-increasing-subsequence/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./longest-increasing-subsequence
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (dynamic programming): dp[i] = length of the longest increasing
// subsequence that ENDS at index i.
// For each i, look at every earlier j with nums[j] < nums[i]:
// we can put nums[i] after it, so dp[i] = dp[j] + 1.
// Time: O(n^2).
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	best := 0

	for i := 0; i < len(nums); i++ {
		dp[i] = 1 // the number alone is a subsequence of length 1

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}

		if dp[i] > best {
			best = dp[i]
		}
	}
	return best
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 1, 2, 3, 7.
	check("example 1", lengthOfLIS([]int{9, 1, 4, 2, 3, 3, 7}), 4)

	// Example 2 from the problem: 0, 1, 2, 3.
	check("example 2", lengthOfLIS([]int{0, 3, 1, 3, 2, 3}), 4)

	// One number.
	check("one number", lengthOfLIS([]int{5}), 1)

	// Going down only: any single number.
	check("descending", lengthOfLIS([]int{5, 4, 3, 2, 1}), 1)

	// Already increasing: everything.
	check("ascending", lengthOfLIS([]int{1, 2, 3, 4, 5}), 5)

	// Equal numbers are not "increasing".
	check("all the same", lengthOfLIS([]int{7, 7, 7}), 1)

	// The classic example: 2, 3, 7, 101.
	check("classic", lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}), 4)

	// 1, 3, 6, 7, 9, 10.
	check("longer case", lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}), 6)

	// No numbers.
	check("empty list", lengthOfLIS([]int{}), 0)

	// Negative numbers: -3, -2, -1.
	check("negative numbers", lengthOfLIS([]int{-3, 5, -2, -1}), 3)

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
