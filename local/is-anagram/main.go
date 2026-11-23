// Problem: Valid Anagram
// Link:    https://neetcode.io/problems/is-anagram
// Source:  Data Structures & Algorithms/is-anagram/submission-5.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./is-anagram
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

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	runes := []rune(s)
	runesT := []rune(t)
	slices.Sort(runes)
	slices.Sort(runesT)

	return slices.Equal(runes, runesT)
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: same letters, different order.
	check("example 1: racecar / carrace", isAnagram("racecar", "carrace"), true)

	// Example 2 from the problem: 'r' and 'm' differ.
	check("example 2: jar / jam", isAnagram("jar", "jam"), false)

	// Two empty strings have the same (zero) letters.
	check("both empty", isAnagram("", ""), true)

	// One letter, the same.
	check("same single letter", isAnagram("a", "a"), true)

	// One letter, different.
	check("different single letter", isAnagram("a", "b"), false)

	// Different lengths can never be anagrams.
	check("different lengths", isAnagram("ab", "a"), false)

	// Same letters but a different COUNT of each letter.
	check("same letters, different counts", isAnagram("aab", "abb"), false)

	// A classic pair.
	check("listen / silent", isAnagram("listen", "silent"), true)

	// Upper and lower case are different letters.
	check("case matters", isAnagram("a", "A"), false)

	// A word is an anagram of itself.
	check("identical words", isAnagram("hello", "hello"), true)

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
