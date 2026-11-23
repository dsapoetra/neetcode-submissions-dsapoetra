// Problem:    Generate Parentheses
// Difficulty: Medium
// Link:       https://neetcode.io/problems/generate-parentheses
// Source:     Data Structures & Algorithms/generate-parentheses/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a number n, return every valid string that can be made from n pairs
// of round brackets. The order of the answer does not matter.
//
// Example 1:
//
//	Input:  n = 1
//	Output: ["()"]
//
// Example 2:
//
//	Input:  n = 3
//	Output: ["((()))", "(()())", "(())()", "()(())", "()()()"]
//
// How to run (from the "local" folder):
//
//	go run ./070-generate-parentheses
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"sort"
)

// ===================== Solution =====================

// Idea (backtracking): build the string one bracket at a time.
//   - we may add "(" while we have used fewer than n of them
//   - we may add ")" only if there is an open bracket waiting for it
//     (fewer ")" than "(" so far)
//
// Following these two rules, every finished string is valid.
func generateParenthesis(n int) []string {
	res := []string{}

	// build adds one bracket at a time to current.
	// open / closed = how many "(" and ")" are already in current
	var build func(current string, open int, closed int)
	build = func(current string, open int, closed int) {
		// n pairs means 2*n characters: the string is finished
		if len(current) == 2*n {
			res = append(res, current)
			return
		}
		// option 1: add "(" if we still have some left
		if open < n {
			build(current+"(", open+1, closed)
		}
		// option 2: add ")" if there is an open bracket to close
		if closed < open {
			build(current+")", open, closed+1)
		}
	}

	// start with an empty string and no brackets used
	build("", 0, 0)
	return res
}

// ===================== Test cases =====================

func main() {
	// The strings may come back in any order, so we sort them before comparing.

	// Example 1 from the problem.
	check("example 1: n = 1", `sortedStrings(generateParenthesis(1))`,
		sortedStrings(generateParenthesis(1)), []string{"()"})

	// Example 2 from the problem.
	check("example 2: n = 3", `sortedStrings(generateParenthesis(3))`,
		sortedStrings(generateParenthesis(3)), []string{"((()))", "(()())", "(())()", "()(())", "()()()"})

	// Two pairs.
	check("n = 2", `sortedStrings(generateParenthesis(2))`,
		sortedStrings(generateParenthesis(2)), []string{"(())", "()()"})

	// The number of answers follows the Catalan numbers: 1, 2, 5, 14, 42, 132 ...
	check("n = 4 gives 14", `len(generateParenthesis(4))`,
		len(generateParenthesis(4)), 14)
	check("n = 5 gives 42", `len(generateParenthesis(5))`,
		len(generateParenthesis(5)), 42)
	check("n = 6 gives 132", `len(generateParenthesis(6))`,
		len(generateParenthesis(6)), 132)

	report()
}

// ===================== Helpers for the tests =====================

// sortedStrings sorts the words. We use it when the answer may be in any order.
func sortedStrings(words []string) []string {
	sort.Strings(words)
	return words
}

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
