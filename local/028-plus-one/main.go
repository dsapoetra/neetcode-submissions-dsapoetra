// Problem:    Plus One
// Difficulty: Easy
// Link:       https://neetcode.io/problems/plus-one
// Source:     Data Structures & Algorithms/plus-one/submission-1.go (the solution below is a copy of it)
//
// THE PROBLEM
// A number is given as a list of its digits (most important digit first).
// Add one to the number and return the result as a list of digits.
//
// Example 1:
//
//	Input:  digits = [1, 2, 3, 4]
//	Output: [1, 2, 3, 5]
//
// Example 2:
//
//	Input:  digits = [9, 9, 9]
//	Output: [1, 0, 0, 0]
//
// How to run (from the "local" folder):
//
//	go run ./028-plus-one
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: add 1 like on paper, starting at the last digit.
//   - a digit smaller than 9 just goes up by one, and we are done
//   - a 9 becomes 0 and the "carry" moves on to the digit on its left
//
// If every digit was a 9 (like 999), we fall out of the loop and need a
// new 1 in front: 1000.
// Time: O(n).
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0 // 9 + 1 = 10: write 0, carry on
	}

	// all digits were 9
	return append([]int{1}, digits...)
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", `plusOne([]int{1, 2, 3, 4})`,
		plusOne([]int{1, 2, 3, 4}), []int{1, 2, 3, 5})

	// Example 2 from the problem: 999 + 1 = 1000, one digit more.
	check("example 2: all nines", `plusOne([]int{9, 9, 9})`,
		plusOne([]int{9, 9, 9}), []int{1, 0, 0, 0})

	// Zero.
	check("zero", `plusOne([]int{0})`,
		plusOne([]int{0}), []int{1})

	// A single 9.
	check("single nine", `plusOne([]int{9})`,
		plusOne([]int{9}), []int{1, 0})

	// The carry stops at the 2.
	check("one carry", `plusOne([]int{1, 2, 9})`,
		plusOne([]int{1, 2, 9}), []int{1, 3, 0})

	// The carry runs over two nines.
	check("two carries", `plusOne([]int{1, 9, 9})`,
		plusOne([]int{1, 9, 9}), []int{2, 0, 0})

	// The classic example.
	check("classic", `plusOne([]int{4, 3, 2, 1})`,
		plusOne([]int{4, 3, 2, 1}), []int{4, 3, 2, 2})

	// 89 + 1 = 90.
	check("89", `plusOne([]int{8, 9})`,
		plusOne([]int{8, 9}), []int{9, 0})

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
