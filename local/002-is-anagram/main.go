// Problem:    Valid Anagram
// Difficulty: Easy
// Link:       https://neetcode.io/problems/is-anagram
// Source:     Data Structures & Algorithms/is-anagram/submission-5.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given two strings s and t, return true if t is an anagram of s: it uses
// exactly the same letters, the same number of times, in any order.
//
// Example 1:
//
//	Input:  s = "racecar", t = "carrace"
//	Output: true
//
// Example 2:
//
//	Input:  s = "jar", t = "jam"
//	Output: false
//
// How to run (from the "local" folder):
//
//	go run ./002-is-anagram
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"slices"
)

// ===================== Solution =====================

// Idea: two words are anagrams when they use exactly the same letters.
// Sort the letters of both words. If the sorted versions are equal, the
// words are anagrams.
// Time: O(n log n).
func isAnagram(s string, t string) bool {
	// different lengths can never be anagrams
	if len(s) != len(t) {
		return false
	}

	// a string cannot be sorted directly, so turn it into a list of
	// characters (runes) first
	runes := []rune(s)
	runesT := []rune(t)
	slices.Sort(runes)
	slices.Sort(runesT)

	// same letters in the same sorted order means anagram
	return slices.Equal(runes, runesT)
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: same letters, different order.
	check("example 1: racecar / carrace", `isAnagram("racecar", "carrace")`,
		isAnagram("racecar", "carrace"), true)

	// Example 2 from the problem: 'r' and 'm' differ.
	check("example 2: jar / jam", `isAnagram("jar", "jam")`,
		isAnagram("jar", "jam"), false)

	// Two empty strings have the same (zero) letters.
	check("both empty", `isAnagram("", "")`,
		isAnagram("", ""), true)

	// One letter, the same.
	check("same single letter", `isAnagram("a", "a")`,
		isAnagram("a", "a"), true)

	// One letter, different.
	check("different single letter", `isAnagram("a", "b")`,
		isAnagram("a", "b"), false)

	// Different lengths can never be anagrams.
	check("different lengths", `isAnagram("ab", "a")`,
		isAnagram("ab", "a"), false)

	// Same letters but a different COUNT of each letter.
	check("same letters, different counts", `isAnagram("aab", "abb")`,
		isAnagram("aab", "abb"), false)

	// A classic pair.
	check("listen / silent", `isAnagram("listen", "silent")`,
		isAnagram("listen", "silent"), true)

	// Upper and lower case are different letters.
	check("case matters", `isAnagram("a", "A")`,
		isAnagram("a", "A"), false)

	// A word is an anagram of itself.
	check("identical words", `isAnagram("hello", "hello")`,
		isAnagram("hello", "hello"), true)

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
