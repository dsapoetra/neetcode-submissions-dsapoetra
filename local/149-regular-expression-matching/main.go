// Problem:    Regular Expression Matching
// Difficulty: Hard
// Link:       https://neetcode.io/problems/regular-expression-matching
// Source:     Data Structures & Algorithms/regular-expression-matching/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a string s and a pattern p, return true if the pattern matches the
// WHOLE string. In the pattern:
//
//	'.'  matches any one character
//	'*'  means "zero or more of the character before it"
//
// Example 1:
//
//	Input:  s = "aa", p = ".b"
//	Output: false
//
// Example 2:
//
//	Input:  s = "nnn", p = "n*"
//	Output: true
//
// Example 3:
//
//	Input:  s = "xyz", p = ".*z"
//	Output: true
//
// How to run (from the "local" folder):
//
//	go run ./149-regular-expression-matching
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (dynamic programming):
// match[i][j] = true if the first i letters of s match the first j
// characters of the pattern p.
// Look at the last pattern character p[j-1]:
//   - a letter or '.':  the last letter of s must fit, and the rest must
//     match:  match[i-1][j-1]
//   - a '*' (it belongs to the character before it, p[j-2]):
//     use it ZERO times: drop "x*" from the pattern:  match[i][j-2]
//     use it ONE MORE time: the last letter of s must fit p[j-2], and
//     the rest of s must still match the same pattern: match[i-1][j]
//
// Time: O(len(s) * len(p)).
func isMatch(s string, p string) bool {
	match := make([][]bool, len(s)+1)
	for i := range match {
		match[i] = make([]bool, len(p)+1)
	}
	match[0][0] = true // empty text matches empty pattern

	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				// zero times
				if match[i][j-2] {
					match[i][j] = true
				}
				// one more time
				if i > 0 && (p[j-2] == '.' || p[j-2] == s[i-1]) && match[i-1][j] {
					match[i][j] = true
				}
			} else {
				// a normal letter or '.'
				if i > 0 && (p[j-1] == '.' || p[j-1] == s[i-1]) && match[i-1][j-1] {
					match[i][j] = true
				}
			}
		}
	}
	return match[len(s)][len(p)]
}

// ===================== Test cases =====================

func main() {
	// '.' matches any one letter. '*' means "zero or more of the character before it".

	// Example 1 from the problem: the second letter must be "b".
	check("example 1", `isMatch("aa", ".b")`,
		isMatch("aa", ".b"), false)

	// Example 2 from the problem: "n*" can be "nnn".
	check("example 2", `isMatch("nnn", "n*")`,
		isMatch("nnn", "n*"), true)

	// Example 3 from the problem: ".*" can be anything.
	check("example 3", `isMatch("xyz", ".*z")`,
		isMatch("xyz", ".*z"), true)

	// The classic examples.
	check("pattern too short", `isMatch("aa", "a")`,
		isMatch("aa", "a"), false)
	check("a* matches aa", `isMatch("aa", "a*")`,
		isMatch("aa", "a*"), true)
	check(".* matches everything", `isMatch("ab", ".*")`,
		isMatch("ab", ".*"), true)
	check("c* used zero times", `isMatch("aab", "c*a*b")`,
		isMatch("aab", "c*a*b"), true)
	check("mississippi", `isMatch("mississippi", "mis*is*p*.")`,
		isMatch("mississippi", "mis*is*p*."), false)

	// Empty strings.
	check("both empty", `isMatch("", "")`,
		isMatch("", ""), true)
	check("a* can be empty", `isMatch("", "a*")`,
		isMatch("", "a*"), true)
	check("empty pattern", `isMatch("a", "")`,
		isMatch("a", ""), false)

	// ".*" matches "ab", but then the "c" is missing.
	check("letter missing at the end", `isMatch("ab", ".*c")`,
		isMatch("ab", ".*c"), false)

	// "a*" must leave one "a" for the last pattern letter.
	check("star must not take everything", `isMatch("aaa", "a*a")`,
		isMatch("aaa", "a*a"), true)

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
