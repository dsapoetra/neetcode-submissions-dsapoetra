// Problem:    Interleaving String
// Difficulty: Medium
// Link:       https://neetcode.io/problems/interleaving-string
// Source:     Data Structures & Algorithms/interleaving-string/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given three strings, return true if s3 can be built by mixing s1 and s2:
// all letters of both are used, and the letters of s1 keep their order, and
// so do the letters of s2.
//
// Example 1:
//
//	Input:  s1 = "aaaa", s2 = "bbbb", s3 = "aabbbbaa"
//	Output: true
//
// Example 2:
//
//	Input:  s1 = "", s2 = "", s3 = ""
//	Output: true
//
// Example 3:
//
//	Input:  s1 = "abc", s2 = "xyz", s3 = "abxzcy"
//	Output: false
//
// How to run (from the "local" folder):
//
//	go run ./108-interleaving-string
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
// ok[i][j] = true if the first i letters of s1 and the first j letters of
// s2 can be mixed to build the first i+j letters of s3.
// The last letter of that part of s3 comes either
//   - from s1: then s1[i-1] must equal it and ok[i-1][j] must be true
//   - from s2: then s2[j-1] must equal it and ok[i][j-1] must be true
//
// Time: O(len(s1) * len(s2)).
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	ok := make([][]bool, len(s1)+1)
	for i := range ok {
		ok[i] = make([]bool, len(s2)+1)
	}
	ok[0][0] = true // two empty strings build the empty string

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			// take the last letter from s1
			if i > 0 && s1[i-1] == s3[i+j-1] && ok[i-1][j] {
				ok[i][j] = true
			}
			// take the last letter from s2
			if j > 0 && s2[j-1] == s3[i+j-1] && ok[i][j-1] {
				ok[i][j] = true
			}
		}
	}
	return ok[len(s1)][len(s2)]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: aa + bbbb + aa.
	check("example 1", `isInterleave("aaaa", "bbbb", "aabbbbaa")`,
		isInterleave("aaaa", "bbbb", "aabbbbaa"), true)

	// Example 2 from the problem: three empty strings.
	check("example 2: all empty", `isInterleave("", "", "")`,
		isInterleave("", "", ""), true)

	// Example 3 from the problem: "z" would have to come before "y".
	check("example 3: wrong order", `isInterleave("abc", "xyz", "abxzcy")`,
		isInterleave("abc", "xyz", "abxzcy"), false)

	// The classic examples.
	check("classic true", `isInterleave("aabcc", "dbbca", "aadbbcbcac")`,
		isInterleave("aabcc", "dbbca", "aadbbcbcac"), true)
	check("classic false", `isInterleave("aabcc", "dbbca", "aadbbbaccc")`,
		isInterleave("aabcc", "dbbca", "aadbbbaccc"), false)

	// One string is empty.
	check("s2 empty", `isInterleave("a", "", "a")`,
		isInterleave("a", "", "a"), true)
	check("s1 empty", `isInterleave("", "b", "b")`,
		isInterleave("", "b", "b"), true)

	// Both orders work.
	check("a then b", `isInterleave("a", "b", "ab")`,
		isInterleave("a", "b", "ab"), true)
	check("b then a", `isInterleave("a", "b", "ba")`,
		isInterleave("a", "b", "ba"), true)

	// s3 has the wrong length.
	check("wrong length", `isInterleave("a", "b", "abc")`,
		isInterleave("a", "b", "abc"), false)

	// "cd" sits in the middle of "ab".
	check("s2 inside s1", `isInterleave("ab", "cd", "acdb")`,
		isInterleave("ab", "cd", "acdb"), true)

	// "d" before "c" breaks the order of s2.
	check("order of s2 broken", `isInterleave("ab", "cd", "adcb")`,
		isInterleave("ab", "cd", "adcb"), false)

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
