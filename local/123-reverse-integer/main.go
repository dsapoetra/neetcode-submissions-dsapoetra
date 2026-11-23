// Problem:    Reverse Integer
// Difficulty: Medium
// Link:       https://neetcode.io/problems/reverse-integer
// Source:     Data Structures & Algorithms/reverse-integer/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a signed 32-bit number, return it with its digits reversed (the
// sign stays). If the reversed number does not fit into a signed 32-bit
// integer, return 0.
//
// Example 1:
//
//	Input:  x = 1234
//	Output: 4321
//
// Example 2:
//
//	Input:  x = -1234
//	Output: -4321
//
// Example 3:
//
//	Input:  x = 1234236467
//	Output: 0
//
// How to run (from the "local" folder):
//
//	go run ./123-reverse-integer
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: take the digits off the end of x one by one (x % 10) and push
// them onto the end of the result (result * 10 + digit).
// In Go, % keeps the sign, so negative numbers work without extra code.
// Before pushing a digit, check that the result would still fit in a
// 32-bit integer; if not, the answer is 0.
// Time: O(number of digits).
func reverse(x int) int {
	const maxInt32 = 2147483647
	const minInt32 = -2147483648

	result := 0
	for x != 0 {
		// the last digit of x (negative if x is negative)
		digit := x % 10
		// remove the last digit from x
		x /= 10

		// would result*10 + digit leave the 32-bit range?
		// too big: 2147483647 ends in 7, so a last digit above 7 does not fit
		if result > maxInt32/10 || (result == maxInt32/10 && digit > 7) {
			return 0
		}
		// too small: -2147483648 ends in 8, so a last digit below -8 does not fit
		if result < minInt32/10 || (result == minInt32/10 && digit < -8) {
			return 0
		}

		// safe: push the digit onto the end of the result
		result = result*10 + digit
	}
	return result
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", `reverse(1234)`,
		reverse(1234), 4321)

	// Example 2 from the problem: the sign stays.
	check("example 2: negative", `reverse(-1234)`,
		reverse(-1234), -4321)

	// Example 3 from the problem: 7646324321 does not fit in 32 bits.
	check("example 3: too big", `reverse(1234236467)`,
		reverse(1234236467), 0)

	// Zero.
	check("zero", `reverse(0)`,
		reverse(0), 0)

	// The zero at the end disappears.
	check("ends with zero", `reverse(120)`,
		reverse(120), 21)
	check("ends with two zeros", `reverse(100)`,
		reverse(100), 1)

	// One digit.
	check("one digit", `reverse(5)`,
		reverse(5), 5)
	check("one negative digit", `reverse(-7)`,
		reverse(-7), -7)

	// The biggest and smallest 32-bit numbers do not fit when reversed.
	check("biggest 32-bit number", `reverse(2147483647)`,
		reverse(2147483647), 0)
	check("smallest 32-bit number", `reverse(-2147483648)`,
		reverse(-2147483648), 0)

	// This one just fits: 2147483641 is below the limit 2147483647.
	check("just fits", `reverse(1463847412)`,
		reverse(1463847412), 2147483641)

	// This one is just too big: 9646324351.
	check("just too big", `reverse(1534236469)`,
		reverse(1534236469), 0)

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
