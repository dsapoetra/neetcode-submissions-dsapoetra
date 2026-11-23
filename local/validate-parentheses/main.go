// Problem: Valid Parentheses
// Link:    https://neetcode.io/problems/validate-parentheses
// Source:  Data Structures & Algorithms/validate-parentheses/submission-2.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./validate-parentheses
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: use a stack. Push every opening bracket. When we see a closing
// bracket, it must match the opening bracket on TOP of the stack
// (the most recent one that is still open). At the end nothing may be
// left open.
// Time: O(n).
func isValid(s string) bool {
	// closing bracket -> the opening bracket it needs
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		opening, isClosing := pairs[ch]

		if !isClosing {
			// opening bracket: remember it
			stack = append(stack, ch)
			continue
		}

		// closing bracket: the top of the stack must be its partner
		if len(stack) == 0 || stack[len(stack)-1] != opening {
			return false
		}
		stack = stack[:len(stack)-1] // pop
	}

	// valid only if every opening bracket was closed
	return len(stack) == 0
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1: []", isValid("[]"), true)

	// Example 2 from the problem: brackets inside brackets.
	check("example 2: nested", isValid("([{}])"), true)

	// Example 3 from the problem: closed in the wrong order.
	check("example 3: wrong order", isValid("[(])"), false)

	// Nothing to close is fine.
	check("empty string", isValid(""), true)

	// Opened but never closed.
	check("only an opening bracket", isValid("("), false)

	// Closed but never opened.
	check("only a closing bracket", isValid(")"), false)

	// Pairs one after another.
	check("pairs in a row", isValid("()[]{}"), true)

	// Wrong type of closing bracket.
	check("wrong type", isValid("(]"), false)

	// Two opened, none closed.
	check("two opening brackets", isValid("(("), false)

	// One closing bracket too many.
	check("one closing too many", isValid("())"), false)

	// One opening bracket too many.
	check("one opening too many", isValid("(()"), false)

	// A longer valid mix.
	check("long valid mix", isValid("{[()()]}"), true)

	// Right brackets, but closing comes first.
	check("closing before opening", isValid("}{"), false)

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
