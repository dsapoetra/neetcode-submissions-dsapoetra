// Problem:    Distinct Subsequences
// Difficulty: Hard
// Link:       https://neetcode.io/problems/count-subsequences
// Source:     Data Structures & Algorithms/count-subsequences/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given two strings s and t, return in how many different ways t can be
// made from s by deleting letters (the remaining letters keep their order).
//
// Example 1:
//
//	Input:  s = "caaat", t = "cat"
//	Output: 3             (use any one of the three a's)
//
// Example 2:
//
//	Input:  s = "xxyxy", t = "xy"
//	Output: 5
//
// How to run (from the "local" folder):
//
//	go run ./147-count-subsequences
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
// ways[i][j] = number of ways to build the first j letters of t from the
// first i letters of s.
// Look at letter s[i-1]:
//   - we can always SKIP it:                       ways[i-1][j]
//   - if it equals t[j-1] we can also USE it:      ways[i-1][j-1]
//
// Building the empty string (j = 0) always has exactly 1 way: take nothing.
// Time: O(len(s) * len(t)).
func numDistinct(s string, t string) int {
	ways := make([][]int, len(s)+1)
	for i := range ways {
		ways[i] = make([]int, len(t)+1)
		ways[i][0] = 1 // the empty t: one way
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			ways[i][j] = ways[i-1][j] // skip s[i-1]
			if s[i-1] == t[j-1] {
				ways[i][j] += ways[i-1][j-1] // use s[i-1] for t[j-1]
			}
		}
	}
	return ways[len(s)][len(t)]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: "c", one of the three "a", and "t".
	check("example 1", `numDistinct("caaat", "cat")`,
		numDistinct("caaat", "cat"), 3)

	// Example 2 from the problem.
	check("example 2", `numDistinct("xxyxy", "xy")`,
		numDistinct("xxyxy", "xy"), 5)

	// The classic examples.
	check("classic 1", `numDistinct("rabbbit", "rabbit")`,
		numDistinct("rabbbit", "rabbit"), 3)
	check("classic 2", `numDistinct("babgbag", "bag")`,
		numDistinct("babgbag", "bag"), 5)

	// Single letters.
	check("same letter", `numDistinct("a", "a")`,
		numDistinct("a", "a"), 1)
	check("different letter", `numDistinct("a", "b")`,
		numDistinct("a", "b"), 0)

	// The empty string can be built in exactly one way: take nothing.
	check("empty t", `numDistinct("abc", "")`,
		numDistinct("abc", ""), 1)

	// Nothing can be built from an empty string.
	check("empty s", `numDistinct("", "a")`,
		numDistinct("", "a"), 0)

	// Pick 1 of 3, or 2 of 3.
	check("one of three", `numDistinct("aaa", "a")`,
		numDistinct("aaa", "a"), 3)
	check("two of three", `numDistinct("aaa", "aa")`,
		numDistinct("aaa", "aa"), 3)

	// t is longer than s.
	check("t longer than s", `numDistinct("ab", "abc")`,
		numDistinct("ab", "abc"), 0)

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
