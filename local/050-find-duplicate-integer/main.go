// Problem:    Find the Duplicate Number
// Difficulty: Medium
// Link:       https://neetcode.io/problems/find-duplicate-integer
// Source:     Data Structures & Algorithms/find-duplicate-integer/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// The list holds n + 1 numbers, each between 1 and n. Exactly one number
// appears more than once. Return that number.
//
// Example 1:
//
//	Input:  nums = [1, 2, 3, 2, 2]
//	Output: 2
//
// Example 2:
//
//	Input:  nums = [1, 2, 3, 4, 4]
//	Output: 4
//
// How to run (from the "local" folder):
//
//	go run ./050-find-duplicate-integer
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: the numbers are between 1 and n, so we can use a list of
// true/false flags: seen[x] = true once we have met the number x.
// The first number we meet twice is the answer.
// Time: O(n), extra space: O(n).
//
// (The follow-up asks for no extra memory. That needs the "slow and fast
// pointer" cycle trick, which is clever but much harder to understand.)
func findDuplicate(nums []int) int {
	// seen[x] is true once we have met the number x
	// (the numbers go up to len(nums)-1, so every number is a valid index)
	seen := make([]bool, len(nums))

	for _, num := range nums {
		// we met this number before: it is the duplicate
		if seen[num] {
			return num
		}
		// first time we meet this number: remember it
		seen[num] = true
	}
	return -1 // never happens: the problem promises a duplicate
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", `findDuplicate([]int{1, 2, 3, 2, 2})`,
		findDuplicate([]int{1, 2, 3, 2, 2}), 2)

	// Example 2 from the problem.
	check("example 2", `findDuplicate([]int{1, 2, 3, 4, 4})`,
		findDuplicate([]int{1, 2, 3, 4, 4}), 4)

	// The smallest possible list.
	check("two numbers", `findDuplicate([]int{1, 1})`,
		findDuplicate([]int{1, 1}), 1)

	// The classic examples.
	check("classic 1", `findDuplicate([]int{1, 3, 4, 2, 2})`,
		findDuplicate([]int{1, 3, 4, 2, 2}), 2)
	check("classic 2", `findDuplicate([]int{3, 1, 3, 4, 2})`,
		findDuplicate([]int{3, 1, 3, 4, 2}), 3)

	// Only one number, many times.
	check("all the same", `findDuplicate([]int{2, 2, 2, 2})`,
		findDuplicate([]int{2, 2, 2, 2}), 2)

	// The duplicate appears three times.
	check("three copies", `findDuplicate([]int{1, 4, 4, 2, 4})`,
		findDuplicate([]int{1, 4, 4, 2, 4}), 4)

	// The duplicate is at the end.
	check("duplicate at the end", `findDuplicate([]int{2, 1, 1})`,
		findDuplicate([]int{2, 1, 1}), 1)

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
