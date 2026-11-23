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
	check("example 1: []", `isValid("[]")`,
		isValid("[]"), true)

	// Example 2 from the problem: brackets inside brackets.
	check("example 2: nested", `isValid("([{}])")`,
		isValid("([{}])"), true)

	// Example 3 from the problem: closed in the wrong order.
	check("example 3: wrong order", `isValid("[(])")`,
		isValid("[(])"), false)

	// Nothing to close is fine.
	check("empty string", `isValid("")`,
		isValid(""), true)

	// Opened but never closed.
	check("only an opening bracket", `isValid("(")`,
		isValid("("), false)

	// Closed but never opened.
	check("only a closing bracket", `isValid(")")`,
		isValid(")"), false)

	// Pairs one after another.
	check("pairs in a row", `isValid("()[]{}")`,
		isValid("()[]{}"), true)

	// Wrong type of closing bracket.
	check("wrong type", `isValid("(]")`,
		isValid("(]"), false)

	// Two opened, none closed.
	check("two opening brackets", `isValid("((")`,
		isValid("(("), false)

	// One closing bracket too many.
	check("one closing too many", `isValid("())")`,
		isValid("())"), false)

	// One opening bracket too many.
	check("one opening too many", `isValid("(()")`,
		isValid("(()"), false)

	// A longer valid mix.
	check("long valid mix", `isValid("{[()()]}")`,
		isValid("{[()()]}"), true)

	// Right brackets, but closing comes first.
	check("closing before opening", `isValid("}{")`,
		isValid("}{"), false)

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
