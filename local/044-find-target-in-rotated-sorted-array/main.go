// Problem:    Search in Rotated Sorted Array
// Difficulty: Medium
// Link:       https://neetcode.io/problems/find-target-in-rotated-sorted-array
// Source:     Data Structures & Algorithms/find-target-in-rotated-sorted-array/submission-1.go (the solution below is a copy of it)
//
// THE PROBLEM
// A sorted list of different numbers was rotated (see the problem above).
// Return the index of target, or -1 if it is not in the list.
// The solution should run in O(log n) time.
//
// Example 1:
//
//	Input:  nums = [3, 4, 5, 6, 1, 2], target = 1
//	Output: 4
//
// Example 2:
//
//	Input:  nums = [3, 5, 6, 0, 1, 2], target = 4
//	Output: -1
//
// How to run (from the "local" folder):
//
//	go run ./044-find-target-in-rotated-sorted-array
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
	check("example 1", `search([]int{3, 4, 5, 6, 1, 2}, 1)`,
		search([]int{3, 4, 5, 6, 1, 2}, 1), 4)

	// Example 2 from the problem: 4 is not in the list.
	check("example 2: not found", `search([]int{3, 5, 6, 0, 1, 2}, 4)`,
		search([]int{3, 5, 6, 0, 1, 2}, 4), -1)

	// One number, and it is the target.
	check("one number, found", `search([]int{1}, 1)`,
		search([]int{1}, 1), 0)

	// One number, and it is not the target.
	check("one number, not found", `search([]int{1}, 0)`,
		search([]int{1}, 0), -1)

	// The classic example: target is the smallest number.
	check("target is the minimum", `search([]int{4, 5, 6, 7, 0, 1, 2}, 0)`,
		search([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4)

	// Target is the first number.
	check("target at the start", `search([]int{4, 5, 6, 7, 0, 1, 2}, 4)`,
		search([]int{4, 5, 6, 7, 0, 1, 2}, 4), 0)

	// Target is the last number.
	check("target at the end", `search([]int{4, 5, 6, 7, 0, 1, 2}, 2)`,
		search([]int{4, 5, 6, 7, 0, 1, 2}, 2), 6)

	// Target would be in the gap between the two sorted parts.
	check("target in the gap", `search([]int{4, 5, 6, 7, 0, 1, 2}, 3)`,
		search([]int{4, 5, 6, 7, 0, 1, 2}, 3), -1)

	// Not rotated: a normal binary search.
	check("not rotated", `search([]int{1, 2, 3, 4, 5}, 4)`,
		search([]int{1, 2, 3, 4, 5}, 4), 3)

	// Two numbers, rotated.
	check("two numbers, second", `search([]int{3, 1}, 1)`,
		search([]int{3, 1}, 1), 1)
	check("two numbers, first", `search([]int{3, 1}, 3)`,
		search([]int{3, 1}, 3), 0)

	// Three numbers, target is the biggest and comes first.
	check("three numbers", `search([]int{5, 1, 3}, 5)`,
		search([]int{5, 1, 3}, 5), 0)

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
