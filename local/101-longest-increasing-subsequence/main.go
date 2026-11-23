// Problem:    Longest Increasing Subsequence
// Difficulty: Medium
// Link:       https://neetcode.io/problems/longest-increasing-subsequence
// Source:     Data Structures & Algorithms/longest-increasing-subsequence/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a list of numbers, return the length of the longest strictly
// increasing subsequence. A subsequence keeps the order of the list but may
// skip numbers.
//
// Example 1:
//
//	Input:  nums = [9, 1, 4, 2, 3, 3, 7]
//	Output: 4             (1, 2, 3, 7)
//
// Example 2:
//
//	Input:  nums = [0, 3, 1, 3, 2, 3]
//	Output: 4             (0, 1, 2, 3)
//
// How to run (from the "local" folder):
//
//	go run ./101-longest-increasing-subsequence
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
	// the longest subsequence found so far, over all end positions
	best := 0

	for i := 0; i < len(nums); i++ {
		dp[i] = 1 // the number alone is a subsequence of length 1

		// try to put nums[i] after every earlier, smaller number nums[j]
		for j := 0; j < i; j++ {
			// only update if that gives something longer than what we have
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}

		// the answer can end at any index, so keep the best dp[i]
		if dp[i] > best {
			best = dp[i]
		}
	}
	return best
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 1, 2, 3, 7.
	check("example 1", `lengthOfLIS([]int{9, 1, 4, 2, 3, 3, 7})`,
		lengthOfLIS([]int{9, 1, 4, 2, 3, 3, 7}), 4)

	// Example 2 from the problem: 0, 1, 2, 3.
	check("example 2", `lengthOfLIS([]int{0, 3, 1, 3, 2, 3})`,
		lengthOfLIS([]int{0, 3, 1, 3, 2, 3}), 4)

	// One number.
	check("one number", `lengthOfLIS([]int{5})`,
		lengthOfLIS([]int{5}), 1)

	// Going down only: any single number.
	check("descending", `lengthOfLIS([]int{5, 4, 3, 2, 1})`,
		lengthOfLIS([]int{5, 4, 3, 2, 1}), 1)

	// Already increasing: everything.
	check("ascending", `lengthOfLIS([]int{1, 2, 3, 4, 5})`,
		lengthOfLIS([]int{1, 2, 3, 4, 5}), 5)

	// Equal numbers are not "increasing".
	check("all the same", `lengthOfLIS([]int{7, 7, 7})`,
		lengthOfLIS([]int{7, 7, 7}), 1)

	// The classic example: 2, 3, 7, 101.
	check("classic", `lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})`,
		lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}), 4)

	// 1, 3, 6, 7, 9, 10.
	check("longer case", `lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})`,
		lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}), 6)

	// No numbers.
	check("empty list", `lengthOfLIS([]int{})`,
		lengthOfLIS([]int{}), 0)

	// Negative numbers: -3, -2, -1.
	check("negative numbers", `lengthOfLIS([]int{-3, 5, -2, -1})`,
		lengthOfLIS([]int{-3, 5, -2, -1}), 3)

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
