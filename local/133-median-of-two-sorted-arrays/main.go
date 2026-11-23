// Problem:    Median of Two Sorted Arrays
// Difficulty: Hard
// Link:       https://neetcode.io/problems/median-of-two-sorted-arrays
// Source:     Data Structures & Algorithms/median-of-two-sorted-arrays/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given two sorted lists, return the median of all their numbers together:
// the middle value, or the average of the two middle values if the total
// count is even.
//
// Example 1:
//
//	Input:  nums1 = [1, 2], nums2 = [3]
//	Output: 2.0
//
// Example 2:
//
//	Input:  nums1 = [1, 3], nums2 = [2, 4]
//	Output: 2.5
//
// How to run (from the "local" folder):
//
//	go run ./133-median-of-two-sorted-arrays
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: both lists are sorted, so we can walk through them together in
// sorted order (like the merge step of merge sort) WITHOUT building the
// merged list. We only need to walk until the middle and remember the
// last two numbers we passed.
//   - odd total:  the median is the middle number
//   - even total: the median is the average of the two middle numbers
//
// This is the simple version: O(m + n).
// The faster version (binary search, O(log(m + n))) is much harder to get right.
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	middle := total / 2

	i, j := 0, 0 // positions in nums1 and nums2
	previous := 0
	current := 0

	// take the next smallest number, middle+1 times
	for step := 0; step <= middle; step++ {
		previous = current

		if j >= len(nums2) || (i < len(nums1) && nums1[i] <= nums2[j]) {
			// nums2 is used up, or the next number of nums1 is smaller
			current = nums1[i]
			i++
		} else {
			current = nums2[j]
			j++
		}
	}

	if total%2 == 1 {
		return float64(current)
	}
	return float64(previous+current) / 2.0
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: merged [1,2,3], middle is 2.
	check("example 1", `findMedianSortedArrays([]int{1, 2}, []int{3})`,
		findMedianSortedArrays([]int{1, 2}, []int{3}), 2.0)

	// Example 2 from the problem: merged [1,2,3,4], average of 2 and 3.
	check("example 2", `findMedianSortedArrays([]int{1, 3}, []int{2, 4})`,
		findMedianSortedArrays([]int{1, 3}, []int{2, 4}), 2.5)

	// One list is empty.
	check("first empty", `findMedianSortedArrays([]int{}, []int{1})`,
		findMedianSortedArrays([]int{}, []int{1}), 1.0)
	check("second empty", `findMedianSortedArrays([]int{2}, []int{})`,
		findMedianSortedArrays([]int{2}, []int{}), 2.0)

	// The lists do not overlap.
	check("no overlap", `findMedianSortedArrays([]int{1, 2}, []int{3, 4})`,
		findMedianSortedArrays([]int{1, 2}, []int{3, 4}), 2.5)
	check("far apart", `findMedianSortedArrays([]int{1, 2, 3}, []int{10, 20, 30})`,
		findMedianSortedArrays([]int{1, 2, 3}, []int{10, 20, 30}), 6.5)

	// Equal values.
	check("all the same", `findMedianSortedArrays([]int{1, 1}, []int{1, 1})`,
		findMedianSortedArrays([]int{1, 1}, []int{1, 1}), 1.0)

	// Negative numbers.
	check("negative numbers", `findMedianSortedArrays([]int{-5, -1}, []int{0})`,
		findMedianSortedArrays([]int{-5, -1}, []int{0}), -1.0)

	// Very different lengths.
	check("different lengths", `findMedianSortedArrays([]int{1}, []int{2, 3, 4, 5})`,
		findMedianSortedArrays([]int{1}, []int{2, 3, 4, 5}), 3.0)

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
