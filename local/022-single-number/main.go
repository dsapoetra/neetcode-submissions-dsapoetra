// Problem:    Single Number
// Difficulty: Easy
// Link:       https://neetcode.io/problems/single-number
// Source:     Data Structures & Algorithms/single-number/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Every number in the list appears exactly twice, except one number that
// appears only once. Return that number. The solution should run in O(n)
// time and use O(1) extra memory.
//
// Example 1:
//
//	Input:  nums = [3, 2, 3]
//	Output: 2
//
// Example 2:
//
//	Input:  nums = [7, 6, 6, 7, 8]
//	Output: 8
//
// How to run (from the "local" folder):
//
//	go run ./022-single-number
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (XOR): the ^ operator has two handy rules:
//
//	x ^ x = 0   (a number cancels itself)
//	x ^ 0 = x
//
// XOR all the numbers together: every pair cancels out and only the
// number without a partner is left.
// Time: O(n), no extra memory.
func singleNumber(nums []int) int {
	// result starts at 0, because 0 ^ x = x
	result := 0
	for _, num := range nums {
		// mix this number in; the second time a number comes, it cancels out
		result ^= num
	}
	// all the pairs are gone, only the single number is left
	return result
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", `singleNumber([]int{3, 2, 3})`,
		singleNumber([]int{3, 2, 3}), 2)

	// Example 2 from the problem.
	check("example 2", `singleNumber([]int{7, 6, 6, 7, 8})`,
		singleNumber([]int{7, 6, 6, 7, 8}), 8)

	// The classic examples.
	check("classic 1", `singleNumber([]int{2, 2, 1})`,
		singleNumber([]int{2, 2, 1}), 1)
	check("classic 2", `singleNumber([]int{4, 1, 2, 1, 2})`,
		singleNumber([]int{4, 1, 2, 1, 2}), 4)

	// Only one number.
	check("one number", `singleNumber([]int{1})`,
		singleNumber([]int{1}), 1)

	// Negative numbers.
	check("negative numbers", `singleNumber([]int{-1, -1, -2})`,
		singleNumber([]int{-1, -1, -2}), -2)

	// The single number is 0.
	check("zero is the answer", `singleNumber([]int{0, 5, 5})`,
		singleNumber([]int{0, 5, 5}), 0)

	// The pair is split by the single number.
	check("single in the middle", `singleNumber([]int{9, 3, 9})`,
		singleNumber([]int{9, 3, 9}), 3)

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
