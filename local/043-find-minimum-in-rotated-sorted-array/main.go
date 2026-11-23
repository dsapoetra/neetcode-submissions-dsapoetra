// Problem:    Find Minimum in Rotated Sorted Array
// Difficulty: Medium
// Link:       https://neetcode.io/problems/find-minimum-in-rotated-sorted-array
// Source:     Data Structures & Algorithms/find-minimum-in-rotated-sorted-array/submission-1.go (the solution below is a copy of it)
//
// THE PROBLEM
// A sorted list of different numbers was "rotated": some numbers were moved
// from the front to the back, e.g. [1,2,3,4,5,6] -> [3,4,5,6,1,2].
// Return the smallest number. The solution should run in O(log n) time.
//
// Example 1:
//
//	Input:  nums = [3, 4, 5, 6, 1, 2]
//	Output: 1
//
// Example 2:
//
//	Input:  nums = [4, 5, 0, 1, 2, 3]
//	Output: 0
//
// How to run (from the "local" folder):
//
//	go run ./043-find-minimum-in-rotated-sorted-array
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
	check("example 1", `findMin([]int{3, 4, 5, 6, 1, 2})`,
		findMin([]int{3, 4, 5, 6, 1, 2}), 1)

	// Example 2 from the problem.
	check("example 2", `findMin([]int{4, 5, 0, 1, 2, 3})`,
		findMin([]int{4, 5, 0, 1, 2, 3}), 0)

	// Example 3 from the problem: not rotated at all, the first number is the minimum.
	check("example 3: not rotated", `findMin([]int{4, 5, 6, 7})`,
		findMin([]int{4, 5, 6, 7}), 4)

	// One number.
	check("one number", `findMin([]int{1})`,
		findMin([]int{1}), 1)

	// Two numbers, rotated.
	check("two numbers, rotated", `findMin([]int{2, 1})`,
		findMin([]int{2, 1}), 1)

	// Two numbers, not rotated.
	check("two numbers, sorted", `findMin([]int{1, 2})`,
		findMin([]int{1, 2}), 1)

	// Rotated by one: the minimum is at index 1.
	check("minimum at index 1", `findMin([]int{5, 1, 2, 3, 4})`,
		findMin([]int{5, 1, 2, 3, 4}), 1)

	// The minimum is the very last number.
	check("minimum at the end", `findMin([]int{2, 3, 4, 5, 1})`,
		findMin([]int{2, 3, 4, 5, 1}), 1)

	// Negative numbers.
	check("negative numbers", `findMin([]int{-3, -2, -1, -10, -5})`,
		findMin([]int{-3, -2, -1, -10, -5}), -10)

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
