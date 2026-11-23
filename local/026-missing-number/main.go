// Problem:    Missing Number
// Difficulty: Easy
// Link:       https://neetcode.io/problems/missing-number
// Source:     Data Structures & Algorithms/missing-number/submission-1.go (the solution below is a copy of it)
//
// THE PROBLEM
// The list holds n different numbers from the range 0..n, so exactly one
// number of that range is missing. Return the missing number.
//
// Example 1:
//
//	Input:  nums = [1, 2, 3]
//	Output: 0
//
// Example 2:
//
//	Input:  nums = [0, 2]
//	Output: 1
//
// How to run (from the "local" folder):
//
//	go run ./026-missing-number
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
	check("example 1: 0 is missing", `missingNumber([]int{1, 2, 3})`,
		missingNumber([]int{1, 2, 3}), 0)

	// Example 2 from the problem: 1 is missing.
	check("example 2: 1 is missing", `missingNumber([]int{0, 2})`,
		missingNumber([]int{0, 2}), 1)

	// One number.
	check("only 0", `missingNumber([]int{0})`,
		missingNumber([]int{0}), 1)
	check("only 1", `missingNumber([]int{1})`,
		missingNumber([]int{1}), 0)

	// The classic example.
	check("classic", `missingNumber([]int{3, 0, 1})`,
		missingNumber([]int{3, 0, 1}), 2)

	// The biggest number (n) is missing.
	check("biggest is missing", `missingNumber([]int{0, 1})`,
		missingNumber([]int{0, 1}), 2)
	check("biggest is missing (longer)", `missingNumber([]int{0, 1, 2, 3})`,
		missingNumber([]int{0, 1, 2, 3}), 4)

	// A longer unsorted list.
	check("longer unsorted list", `missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})`,
		missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}), 8)

	// A number in the middle is missing.
	check("middle is missing", `missingNumber([]int{1, 0, 3})`,
		missingNumber([]int{1, 0, 3}), 2)

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
