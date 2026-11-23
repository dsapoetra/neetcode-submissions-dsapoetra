// Problem: Search in Rotated Sorted Array
// Link:    https://neetcode.io/problems/find-target-in-rotated-sorted-array
// Source:  Data Structures & Algorithms/find-target-in-rotated-sorted-array/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./find-target-in-rotated-sorted-array
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (binary search): when we cut a rotated sorted array in the
// middle, at least ONE half is normally sorted. For the sorted half it
// is easy to check whether the target is inside it (between its ends).
// If yes, search that half; if no, search the other half.
// Time: O(log n).
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			// the left half is sorted
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1 // target is inside the left half
			} else {
				left = mid + 1
			}
		} else {
			// the right half is sorted
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1 // target is inside the right half
			} else {
				right = mid - 1
			}
		}
	}
	return -1 // not found
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", search([]int{3, 4, 5, 6, 1, 2}, 1), 4)

	// Example 2 from the problem: 4 is not in the list.
	check("example 2: not found", search([]int{3, 5, 6, 0, 1, 2}, 4), -1)

	// One number, and it is the target.
	check("one number, found", search([]int{1}, 1), 0)

	// One number, and it is not the target.
	check("one number, not found", search([]int{1}, 0), -1)

	// The classic example: target is the smallest number.
	check("target is the minimum", search([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4)

	// Target is the first number.
	check("target at the start", search([]int{4, 5, 6, 7, 0, 1, 2}, 4), 0)

	// Target is the last number.
	check("target at the end", search([]int{4, 5, 6, 7, 0, 1, 2}, 2), 6)

	// Target would be in the gap between the two sorted parts.
	check("target in the gap", search([]int{4, 5, 6, 7, 0, 1, 2}, 3), -1)

	// Not rotated: a normal binary search.
	check("not rotated", search([]int{1, 2, 3, 4, 5}, 4), 3)

	// Two numbers, rotated.
	check("two numbers, second", search([]int{3, 1}, 1), 1)
	check("two numbers, first", search([]int{3, 1}, 3), 0)

	// Three numbers, target is the biggest and comes first.
	check("three numbers", search([]int{5, 1, 3}, 5), 0)

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
