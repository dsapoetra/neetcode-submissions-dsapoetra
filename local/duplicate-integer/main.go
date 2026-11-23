// Problem: Contains Duplicate
// Link:    https://neetcode.io/problems/duplicate-integer
// Source:  Data Structures & Algorithms/duplicate-integer/submission-6.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./duplicate-integer
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

func hasDuplicate(nums []int) bool {
	slices.Sort(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 3 appears twice.
	check("example 1: 3 appears twice", hasDuplicate([]int{1, 2, 3, 3}), true)

	// Example 2 from the problem: every number is different.
	check("example 2: all different", hasDuplicate([]int{1, 2, 3, 4}), false)

	// An empty list has nothing to repeat.
	check("empty list", hasDuplicate([]int{}), false)

	// One number cannot be a duplicate of itself.
	check("one number", hasDuplicate([]int{7}), false)

	// The smallest possible duplicate: two equal numbers.
	check("two equal numbers", hasDuplicate([]int{1, 1}), true)

	// The two equal numbers are far apart (first and last position).
	check("duplicate at both ends", hasDuplicate([]int{3, 1, 2, 3}), true)

	// Negative numbers must work too.
	check("negative duplicate", hasDuplicate([]int{-1, 2, -1}), true)

	// Numbers in descending order, no duplicates.
	check("descending, all different", hasDuplicate([]int{5, 4, 3, 2, 1}), false)

	// The same number many times.
	check("all the same", hasDuplicate([]int{9, 9, 9, 9}), true)

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
