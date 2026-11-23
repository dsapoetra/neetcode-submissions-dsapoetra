// Problem:    Kth Largest Element in an Array
// Difficulty: Medium
// Link:       https://neetcode.io/problems/kth-largest-element-in-an-array
// Source:     Data Structures & Algorithms/kth-largest-element-in-an-array/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given an unsorted list of numbers and a number k, return the k-th largest
// number in sorted order (duplicates count separately).
//
// Example 1:
//
//	Input:  nums = [2, 3, 1, 5, 4], k = 2
//	Output: 4
//
// Example 2:
//
//	Input:  nums = [2, 3, 1, 1, 5, 5, 4], k = 3
//	Output: 4
//
// How to run (from the "local" folder):
//
//	go run ./075-kth-largest-element-in-an-array
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"sort"
)

// ===================== Solution =====================

// Idea: sort the numbers. The k-th largest is then the k-th number
// counted from the end.
// Time: O(n log n).
func findKthLargest(nums []int, k int) int {
	// sort from small to big (this changes the input list)
	sort.Ints(nums)
	// the last index is len-1 (1st largest), so the k-th largest is at len-k
	return nums[len(nums)-k]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: sorted 1,2,3,4,5 -> second largest is 4.
	check("example 1", `findKthLargest([]int{2, 3, 1, 5, 4}, 2)`,
		findKthLargest([]int{2, 3, 1, 5, 4}, 2), 4)

	// Example 2 from the problem: duplicates count separately (5, 5, 4).
	check("example 2: duplicates", `findKthLargest([]int{2, 3, 1, 1, 5, 5, 4}, 3)`,
		findKthLargest([]int{2, 3, 1, 1, 5, 5, 4}, 3), 4)

	// The classic examples.
	check("classic 1", `findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)`,
		findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2), 5)
	check("classic 2", `findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)`,
		findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4), 4)

	// One number.
	check("one number", `findKthLargest([]int{1}, 1)`,
		findKthLargest([]int{1}, 1), 1)

	// k = 1 is the biggest, k = length is the smallest.
	check("k = 1: the biggest", `findKthLargest([]int{7, 9, 8}, 1)`,
		findKthLargest([]int{7, 9, 8}, 1), 9)
	check("k = length: the smallest", `findKthLargest([]int{7, 9, 8}, 3)`,
		findKthLargest([]int{7, 9, 8}, 3), 7)

	// Negative numbers.
	check("negative numbers", `findKthLargest([]int{-1, -5, -3}, 2)`,
		findKthLargest([]int{-1, -5, -3}, 2), -3)

	// All the same.
	check("all the same", `findKthLargest([]int{2, 2, 2}, 2)`,
		findKthLargest([]int{2, 2, 2}, 2), 2)

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
