// Problem: Find Minimum in Rotated Sorted Array
// Link:    https://neetcode.io/problems/find-minimum-in-rotated-sorted-array
// Source:  Data Structures & Algorithms/find-minimum-in-rotated-sorted-array/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./find-minimum-in-rotated-sorted-array
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (binary search): compare the middle number with the rightmost one.
//   - middle > right: the "drop" is to the right of middle, so the
//     minimum is in the right half (and it is not middle)
//   - otherwise: the right half is sorted, so the minimum is middle
//     or something to its left
//
// Time: O(log n).
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1 // minimum is to the right of mid
		} else {
			right = mid // minimum is mid or to its left
		}
	}

	// left == right: only one candidate remains
	return nums[left]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", findMin([]int{3, 4, 5, 6, 1, 2}), 1)

	// Example 2 from the problem.
	check("example 2", findMin([]int{4, 5, 0, 1, 2, 3}), 0)

	// Example 3 from the problem: not rotated at all, the first number is the minimum.
	check("example 3: not rotated", findMin([]int{4, 5, 6, 7}), 4)

	// One number.
	check("one number", findMin([]int{1}), 1)

	// Two numbers, rotated.
	check("two numbers, rotated", findMin([]int{2, 1}), 1)

	// Two numbers, not rotated.
	check("two numbers, sorted", findMin([]int{1, 2}), 1)

	// Rotated by one: the minimum is at index 1.
	check("minimum at index 1", findMin([]int{5, 1, 2, 3, 4}), 1)

	// The minimum is the very last number.
	check("minimum at the end", findMin([]int{2, 3, 4, 5, 1}), 1)

	// Negative numbers.
	check("negative numbers", findMin([]int{-3, -2, -1, -10, -5}), -10)

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
