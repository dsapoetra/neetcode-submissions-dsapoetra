// Problem:    Contains Duplicate
// Difficulty: Easy
// Link:       https://neetcode.io/problems/duplicate-integer
// Source:     Data Structures & Algorithms/duplicate-integer/submission-6.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a list of numbers, return true if any number appears more than once.
// Return false if every number is different.
//
// Example 1:
//
//	Input:  nums = [1, 2, 3, 3]
//	Output: true
//
// Example 2:
//
//	Input:  nums = [1, 2, 3, 4]
//	Output: false
//
// How to run (from the "local" folder):
//
//	go run ./001-duplicate-integer
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"slices"
)

// ===================== Solution =====================

// Idea: sort the numbers. After sorting, equal numbers sit next to each
// other, so we only have to compare each number with the one before it.
// Time: O(n log n). Note: sorting changes the order of the input list.
func hasDuplicate(nums []int) bool {
	slices.Sort(nums)

	// start at 1 so there is always a previous number to compare with
	for i := 1; i < len(nums); i++ {
		// same as the neighbour before it: found a duplicate
		if nums[i] == nums[i-1] {
			return true
		}
	}

	// no two neighbours were equal, so every number is different
	return false
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 3 appears twice.
	check("example 1: 3 appears twice", `hasDuplicate([]int{1, 2, 3, 3})`,
		hasDuplicate([]int{1, 2, 3, 3}), true)

	// Example 2 from the problem: every number is different.
	check("example 2: all different", `hasDuplicate([]int{1, 2, 3, 4})`,
		hasDuplicate([]int{1, 2, 3, 4}), false)

	// An empty list has nothing to repeat.
	check("empty list", `hasDuplicate([]int{})`,
		hasDuplicate([]int{}), false)

	// One number cannot be a duplicate of itself.
	check("one number", `hasDuplicate([]int{7})`,
		hasDuplicate([]int{7}), false)

	// The smallest possible duplicate: two equal numbers.
	check("two equal numbers", `hasDuplicate([]int{1, 1})`,
		hasDuplicate([]int{1, 1}), true)

	// The two equal numbers are far apart (first and last position).
	check("duplicate at both ends", `hasDuplicate([]int{3, 1, 2, 3})`,
		hasDuplicate([]int{3, 1, 2, 3}), true)

	// Negative numbers must work too.
	check("negative duplicate", `hasDuplicate([]int{-1, 2, -1})`,
		hasDuplicate([]int{-1, 2, -1}), true)

	// Numbers in descending order, no duplicates.
	check("descending, all different", `hasDuplicate([]int{5, 4, 3, 2, 1})`,
		hasDuplicate([]int{5, 4, 3, 2, 1}), false)

	// The same number many times.
	check("all the same", `hasDuplicate([]int{9, 9, 9, 9})`,
		hasDuplicate([]int{9, 9, 9, 9}), true)

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
