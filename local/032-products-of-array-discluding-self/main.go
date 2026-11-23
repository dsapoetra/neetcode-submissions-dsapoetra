// Problem:    Product of Array Except Self
// Difficulty: Medium
// Link:       https://neetcode.io/problems/products-of-array-discluding-self
// Source:     Data Structures & Algorithms/products-of-array-discluding-self/submission-2.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a list of numbers, return a list where entry i is the product of
// all numbers EXCEPT nums[i]. The solution should run in O(n) time and
// should not use division.
//
// Example 1:
//
//	Input:  nums = [1, 2, 4, 6]
//	Output: [48, 24, 12, 8]
//
// Example 2:
//
//	Input:  nums = [-1, 0, 1, 2, 3]
//	Output: [0, -6, 0, 0, 0]
//
// How to run (from the "local" folder):
//
//	go run ./032-products-of-array-discluding-self
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: the answer for position i is
// (product of everything left of i) * (product of everything right of i).
// Build the left products in one pass and the right products in a second
// pass. No division is needed, so zeros in the list are not a problem.
// Time: O(n).
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	// Pass 1: res[i] = product of everything LEFT of i
	// left starts at 1 because nothing is left of the first number
	left := 1
	for i := 0; i < len(nums); i++ {
		res[i] = left
		// include nums[i] for the next position
		left *= nums[i]
	}

	// Pass 2: multiply in the product of everything RIGHT of i
	// right starts at 1 because nothing is right of the last number
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= right
		// include nums[i] for the next position (one step to the left)
		right *= nums[i]
	}

	return res
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 2*4*6, 1*4*6, 1*2*6, 1*2*4.
	check("example 1", `productExceptSelf([]int{1, 2, 4, 6})`,
		productExceptSelf([]int{1, 2, 4, 6}), []int{48, 24, 12, 8})

	// Example 2 from the problem: one zero. Only the position of the zero is not 0.
	check("example 2: one zero", `productExceptSelf([]int{-1, 0, 1, 2, 3})`,
		productExceptSelf([]int{-1, 0, 1, 2, 3}), []int{0, -6, 0, 0, 0})

	// Two numbers: they just swap.
	check("two numbers", `productExceptSelf([]int{1, 2})`,
		productExceptSelf([]int{1, 2}), []int{2, 1})

	// Two zeros: every product contains a zero.
	check("two zeros", `productExceptSelf([]int{0, 0})`,
		productExceptSelf([]int{0, 0}), []int{0, 0})

	// The classic example.
	check("classic", `productExceptSelf([]int{1, 2, 3, 4})`,
		productExceptSelf([]int{1, 2, 3, 4}), []int{24, 12, 8, 6})

	// All negative.
	check("all negative", `productExceptSelf([]int{-1, -2, -3})`,
		productExceptSelf([]int{-1, -2, -3}), []int{6, 3, 2})

	// Ones do not change the product.
	check("with ones", `productExceptSelf([]int{5, 1, 1})`,
		productExceptSelf([]int{5, 1, 1}), []int{1, 5, 5})

	// A zero in the middle.
	check("zero in the middle", `productExceptSelf([]int{2, 0, 4})`,
		productExceptSelf([]int{2, 0, 4}), []int{0, 8, 0})

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
