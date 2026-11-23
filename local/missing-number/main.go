// Problem: Missing Number
// Link:    https://neetcode.io/problems/missing-number
// Source:  Data Structures & Algorithms/missing-number/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./missing-number
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: the numbers should be 0, 1, 2, ..., n. Add up 0..n and subtract
// every number we actually have. What is left is the missing number.
// Time: O(n).
func missingNumber(nums []int) int {
	missing := 0
	for i := 0; i < len(nums); i++ {
		missing += (i + 1) // add the expected number (1..n)
		missing -= nums[i] // remove the number we have
	}
	return missing
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 0 is missing.
	check("example 1: 0 is missing", missingNumber([]int{1, 2, 3}), 0)

	// Example 2 from the problem: 1 is missing.
	check("example 2: 1 is missing", missingNumber([]int{0, 2}), 1)

	// One number.
	check("only 0", missingNumber([]int{0}), 1)
	check("only 1", missingNumber([]int{1}), 0)

	// The classic example.
	check("classic", missingNumber([]int{3, 0, 1}), 2)

	// The biggest number (n) is missing.
	check("biggest is missing", missingNumber([]int{0, 1}), 2)
	check("biggest is missing (longer)", missingNumber([]int{0, 1, 2, 3}), 4)

	// A longer unsorted list.
	check("longer unsorted list", missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}), 8)

	// A number in the middle is missing.
	check("middle is missing", missingNumber([]int{1, 0, 3}), 2)

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
