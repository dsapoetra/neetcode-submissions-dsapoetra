// Problem: Longest Substring Without Repeating Characters
// Link:    https://neetcode.io/problems/longest-substring-without-duplicates
// Source:  Data Structures & Algorithms/longest-substring-without-duplicates/submission-2.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./longest-substring-without-duplicates
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (sliding window): the window s[left..right] never contains a
// repeated character. Move right one step at a time. If the new
// character is already inside the window, move left to just after
// its previous position, so the window is valid again.
// Time: O(n).
func lengthOfLongestSubstring(s string) int {
	lastSeen := map[byte]int{} // character -> last index where we saw it
	left := 0
	best := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]

		// seen before AND still inside the window -> jump left past it
		index, seen := lastSeen[ch]
		if seen && index >= left {
			left = index + 1
		}

		lastSeen[ch] = right

		size := right - left + 1
		if size > best {
			best = size
		}
	}
	return best
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: "xyz" (or "zxy", "yzx").
	check("example 1", lengthOfLongestSubstring("zxyzxyz"), 3)

	// Example 2 from the problem: every letter is the same.
	check("example 2: all the same", lengthOfLongestSubstring("xxxx"), 1)

	// Empty string.
	check("empty string", lengthOfLongestSubstring(""), 0)

	// One letter.
	check("one letter", lengthOfLongestSubstring("a"), 1)

	// The classic example: "abc".
	check("classic abcabcbb", lengthOfLongestSubstring("abcabcbb"), 3)

	// "wke" - note "pwke" is not a substring (letters must be next to each other).
	check("classic pwwkew", lengthOfLongestSubstring("pwwkew"), 3)

	// No letter repeats: the whole string.
	check("no repeats", lengthOfLongestSubstring("abcdef"), 6)

	// Tricky: after "abb" the left side must NOT jump back to the old 'a'.
	check("abba", lengthOfLongestSubstring("abba"), 2)

	// Tricky: the best window "vdf" starts in the middle.
	check("dvdf", lengthOfLongestSubstring("dvdf"), 3)

	// A space is a character like any other.
	check("with spaces", lengthOfLongestSubstring("a b c"), 3)

	// "mzuxt" is the longest.
	check("tmmzuxt", lengthOfLongestSubstring("tmmzuxt"), 5)

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
