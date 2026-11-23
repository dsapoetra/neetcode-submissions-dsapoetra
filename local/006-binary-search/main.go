// Problem:    Binary Search
// Difficulty: Easy
// Link:       https://neetcode.io/problems/binary-search
// Source:     Data Structures & Algorithms/binary-search/submission-2.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a sorted list of different numbers and a target, return the index
// of the target, or -1 if it is not in the list. The solution must run in
// O(log n) time.
//
// Example 1:
//
//	Input:  nums = [-1, 0, 2, 4, 6, 8], target = 4
//	Output: 3
//
// Example 2:
//
//	Input:  nums = [-1, 0, 2, 4, 6, 8], target = 3
//	Output: -1
//
// How to run (from the "local" folder):
//
//	go run ./006-binary-search
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: the list is sorted, so look at the middle number.
//   - it is the target      -> found
//   - it is too small       -> the target can only be on the right half
//   - it is too big         -> the target can only be on the left half
//
// Every step throws away half of the list.
// Time: O(log n).
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1 // search the right half
		} else {
			right = mid - 1 // search the left half
		}
	}
	return -1 // not found
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1: found", `search([]int{-1, 0, 2, 4, 6, 8}, 4)`,
		search([]int{-1, 0, 2, 4, 6, 8}, 4), 3)

	// Example 2 from the problem.
	check("example 2: not found", `search([]int{-1, 0, 2, 4, 6, 8}, 3)`,
		search([]int{-1, 0, 2, 4, 6, 8}, 3), -1)

	// One number.
	check("one number, found", `search([]int{5}, 5)`,
		search([]int{5}, 5), 0)
	check("one number, not found", `search([]int{5}, 1)`,
		search([]int{5}, 1), -1)

	// Two numbers.
	check("two numbers, first", `search([]int{1, 2}, 1)`,
		search([]int{1, 2}, 1), 0)
	check("two numbers, second", `search([]int{1, 2}, 2)`,
		search([]int{1, 2}, 2), 1)

	// The classic examples.
	check("classic found", `search([]int{-1, 0, 3, 5, 9, 12}, 9)`,
		search([]int{-1, 0, 3, 5, 9, 12}, 9), 4)
	check("classic not found", `search([]int{-1, 0, 3, 5, 9, 12}, 2)`,
		search([]int{-1, 0, 3, 5, 9, 12}, 2), -1)

	// Smaller / bigger than everything.
	check("smaller than all", `search([]int{1, 3, 5}, 0)`,
		search([]int{1, 3, 5}, 0), -1)
	check("bigger than all", `search([]int{1, 3, 5}, 6)`,
		search([]int{1, 3, 5}, 6), -1)

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
