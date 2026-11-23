// Problem:    Valid Parenthesis String
// Difficulty: Medium
// Link:       https://neetcode.io/problems/valid-parenthesis-string
// Source:     Data Structures & Algorithms/valid-parenthesis-string/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// The string contains '(', ')' and '*'. A '*' can stand for '(' or ')' or
// nothing at all. Return true if the stars can be chosen so that the
// brackets are valid.
//
// Example 1:
//
//	Input:  s = "((**)"
//	Output: true
//
// Example 2:
//
//	Input:  s = "(((*)"
//	Output: false
//
// How to run (from the "local" folder):
//
//	go run ./117-valid-parenthesis-string
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: we do not know what each '*' is, so we track a RANGE for
// "how many open brackets are waiting to be closed":
//
//	low  = the smallest possible number
//	high = the biggest possible number
//	'('  -> both go up
//	')'  -> both go down
//	'*'  -> low goes down (it acts as ')'), high goes up (it acts as '(')
//
// If high drops below 0, there are too many ')' no matter what.
// low can never be below 0 (we just would not use a '*' as ')').
// At the end the string is valid if 0 is possible, that is low == 0.
// Time: O(n).
func checkValidString(s string) bool {
	low := 0
	high := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			low++
			high++
		} else if s[i] == ')' {
			low--
			high--
		} else { // '*'
			low--
			high++
		}

		if high < 0 {
			return false // too many ')' even in the best case
		}
		if low < 0 {
			low = 0
		}
	}
	return low == 0
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: one star is ")", the other is empty.
	check("example 1", `checkValidString("((**)")`,
		checkValidString("((**)"), true)

	// Example 2 from the problem: three "(" but only one ")" and one star.
	check("example 2", `checkValidString("(((*)")`,
		checkValidString("(((*)"), false)

	// The classic examples.
	check("plain pair", `checkValidString("()")`,
		checkValidString("()"), true)
	check("star as nothing", `checkValidString("(*)")`,
		checkValidString("(*)"), true)
	check("star as (", `checkValidString("(*))")`,
		checkValidString("(*))"), true)

	// Empty string.
	check("empty string", `checkValidString("")`,
		checkValidString(""), true)

	// Only stars: use them all as nothing.
	check("one star", `checkValidString("*")`,
		checkValidString("*"), true)
	check("three stars", `checkValidString("***")`,
		checkValidString("***"), true)

	// Wrong order.
	check("closing first", `checkValidString(")(")`,
		checkValidString(")("), false)

	// The star comes BEFORE the "(": it cannot close it.
	check("star before (", `checkValidString("*(")`,
		checkValidString("*("), false)

	// The star comes after the "(": it closes it.
	check("star after (", `checkValidString("(*")`,
		checkValidString("(*"), true)

	// The star opens for the ")".
	check("star before )", `checkValidString("*)")`,
		checkValidString("*)"), true)

	// Too many ")".
	check("too many closing", `checkValidString("())")`,
		checkValidString("())"), false)

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
