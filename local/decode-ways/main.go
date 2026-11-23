// Problem: Decode Ways
// Link:    https://neetcode.io/problems/decode-ways
// Source:  Data Structures & Algorithms/decode-ways/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./decode-ways
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (dynamic programming): ways[i] = number of ways to decode the
// first i digits. The last letter uses either:
//   - one digit  (must be 1..9)   -> add ways[i-1]
//   - two digits (must be 10..26) -> add ways[i-2]
//
// Time: O(n).
func numDecodings(s string) int {
	ways := make([]int, len(s)+1)
	ways[0] = 1 // empty string: one way (decode nothing)

	for i := 1; i <= len(s); i++ {
		// last ONE digit as a letter
		oneDigit := int(s[i-1] - '0')
		if oneDigit >= 1 {
			ways[i] += ways[i-1]
		}

		// last TWO digits as a letter
		if i >= 2 {
			twoDigits := int(s[i-2]-'0')*10 + oneDigit
			if twoDigits >= 10 && twoDigits <= 26 {
				ways[i] += ways[i-2]
			}
		}
	}
	return ways[len(s)]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: "1 2" (AB) or "12" (L).
	check("example 1", numDecodings("12"), 2)

	// Example 2 from the problem: a leading zero cannot be decoded.
	check("example 2: leading zero", numDecodings("01"), 0)

	// The classic example: "2 2 6", "22 6", "2 26".
	check("classic 226", numDecodings("226"), 3)

	// A single zero.
	check("only zero", numDecodings("0"), 0)

	// "10" can only be J (1 and 0 separately does not work).
	check("10", numDecodings("10"), 1)

	// 27 is bigger than 26, so only "2 7".
	check("27 is too big", numDecodings("27"), 1)

	// "06" is not a valid code.
	check("06", numDecodings("06"), 0)

	// Two zeros in a row can never be decoded.
	check("100", numDecodings("100"), 0)

	// "1 1 10 6" or "11 10 6".
	check("11106", numDecodings("11106"), 2)

	// One digit.
	check("one digit", numDecodings("1"), 1)

	// Six 1s: every split works, the count follows Fibonacci.
	check("111111", numDecodings("111111"), 13)

	// Only "2 10 1" works.
	check("2101", numDecodings("2101"), 1)

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
