// Problem:    Permutation in String
// Difficulty: Medium
// Link:       https://neetcode.io/problems/permutation-string
// Source:     Data Structures & Algorithms/permutation-string/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given two strings s1 and s2, return true if s2 contains a permutation of
// s1 as a piece: some letters next to each other in s2 that are exactly the
// letters of s1 in any order.
//
// Example 1:
//
//	Input:  s1 = "abc", s2 = "lecabee"
//	Output: true          ("cab")
//
// Example 2:
//
//	Input:  s1 = "abc", s2 = "lecaabee"
//	Output: false
//
// How to run (from the "local" folder):
//
//	go run ./040-permutation-string
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (sliding window): a permutation of s1 has exactly the same letter
// counts as s1. So slide a window of length len(s1) over s2 and compare
// the letter counts of the window with the letter counts of s1.
// When the window moves one step, one letter enters and one letter leaves.
// Time: O(26 * len(s2)).
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var need [26]int   // letter counts of s1
	var window [26]int // letter counts of the current window in s2

	// the first window
	for i := 0; i < len(s1); i++ {
		need[s1[i]-'a']++
		window[s2[i]-'a']++
	}
	if need == window { // Go can compare two arrays directly
		return true
	}

	// slide the window: add the letter on the right, remove the one on the left
	for right := len(s1); right < len(s2); right++ {
		window[s2[right]-'a']++
		window[s2[right-len(s1)]-'a']--
		if need == window {
			return true
		}
	}
	return false
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: "cab" is a permutation of "abc".
	check("example 1", `checkInclusion("abc", "lecabee")`,
		checkInclusion("abc", "lecabee"), true)

	// Example 2 from the problem: the letters are there, but never next to each other.
	check("example 2", `checkInclusion("abc", "lecaabee")`,
		checkInclusion("abc", "lecaabee"), false)

	// The classic examples.
	check("classic true", `checkInclusion("ab", "eidbaooo")`,
		checkInclusion("ab", "eidbaooo"), true)
	check("classic false", `checkInclusion("ab", "eidboaoo")`,
		checkInclusion("ab", "eidboaoo"), false)

	// One letter.
	check("one letter", `checkInclusion("a", "a")`,
		checkInclusion("a", "a"), true)

	// s1 is longer than s2.
	check("s1 longer than s2", `checkInclusion("abc", "ab")`,
		checkInclusion("abc", "ab"), false)

	// The window is in the middle: "dcd" no, "cda" yes.
	check("window in the middle", `checkInclusion("adc", "dcda")`,
		checkInclusion("adc", "dcda"), true)

	// Repeated letters: the counts must match exactly.
	check("repeated letters, match", `checkInclusion("aab", "aba")`,
		checkInclusion("aab", "aba"), true)
	check("repeated letters, no match", `checkInclusion("aab", "abb")`,
		checkInclusion("aab", "abb"), false)

	// The window is at the very end.
	check("window at the end", `checkInclusion("zyx", "abcxyz")`,
		checkInclusion("zyx", "abcxyz"), true)

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
