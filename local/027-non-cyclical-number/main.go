// Problem:    Happy Number
// Difficulty: Easy
// Link:       https://neetcode.io/problems/non-cyclical-number
// Source:     Data Structures & Algorithms/non-cyclical-number/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Take a positive number and replace it by the sum of the squares of its
// digits. Repeat. If this ends at 1, the number is "happy" (non-cyclical)
// and the answer is true. If it runs in a loop forever, the answer is false.
//
// Example 1:
//
//	Input:  n = 100
//	Output: true          (1*1 + 0*0 + 0*0 = 1)
//
// Example 2:
//
//	Input:  n = 101
//	Output: false         (101 -> 2 -> 4 -> 16 -> 37 -> ... -> 4, a loop)
//
// How to run (from the "local" folder):
//
//	go run ./027-non-cyclical-number
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: keep replacing the number with the sum of the squares of its
// digits. Either we reach 1 (happy), or we come back to a number we have
// seen before. Then we are in a loop and will never reach 1.
// A set of the numbers we have seen detects the loop.
func isHappy(n int) bool {
	// seen[x] is true if the number x came up before
	seen := map[int]bool{}

	for n != 1 {
		if seen[n] {
			return false // we were here before: a loop
		}
		// remember this number, then go to the next one
		seen[n] = true
		n = sumOfDigitSquares(n)
	}
	// the loop ended, so n is 1: the number is happy
	return true
}

// sumOfDigitSquares returns, for example, 1*1 + 9*9 = 82 for 19.
func sumOfDigitSquares(n int) int {
	sum := 0
	// take the digits off one by one, starting from the right
	for n > 0 {
		// the last digit (for 19 this is 9)
		digit := n % 10
		sum += digit * digit
		// drop the last digit (19 becomes 1)
		n /= 10
	}
	return sum
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 1*1 + 0 + 0 = 1.
	check("example 1", `isHappy(100)`,
		isHappy(100), true)

	// Example 2 from the problem: 101 -> 2 -> 4 -> 16 -> ... -> 4 (a loop).
	check("example 2", `isHappy(101)`,
		isHappy(101), false)

	// The classic examples: 19 -> 82 -> 68 -> 100 -> 1.
	check("classic true", `isHappy(19)`,
		isHappy(19), true)
	check("classic false", `isHappy(2)`,
		isHappy(2), false)

	// 1 is happy right away.
	check("1", `isHappy(1)`,
		isHappy(1), true)

	// 7 -> 49 -> 97 -> 130 -> 10 -> 1.
	check("7", `isHappy(7)`,
		isHappy(7), true)

	// 4 is in the loop that all unhappy numbers end in.
	check("4", `isHappy(4)`,
		isHappy(4), false)

	// 10 -> 1.
	check("10", `isHappy(10)`,
		isHappy(10), true)

	// 13 -> 10 -> 1.
	check("13", `isHappy(13)`,
		isHappy(13), true)

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
