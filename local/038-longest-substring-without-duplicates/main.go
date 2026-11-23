// Problem:    Longest Substring Without Repeating Characters
// Difficulty: Medium
// Link:       https://neetcode.io/problems/longest-substring-without-duplicates
// Source:     Data Structures & Algorithms/longest-substring-without-duplicates/submission-2.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a string, return the length of the longest piece (letters next to
// each other) in which no character appears twice.
//
// Example 1:
//
//	Input:  s = "zxyzxyz"
//	Output: 3             ("xyz")
//
// Example 2:
//
//	Input:  s = "xxxx"
//	Output: 1             ("x")
//
// How to run (from the "local" folder):
//
//	go run ./038-longest-substring-without-duplicates
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
	// left = first index of the window
	left := 0
	// the longest window found so far
	best := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]

		// seen before AND still inside the window -> jump left past it
		index, seen := lastSeen[ch]
		if seen && index >= left {
			left = index + 1
		}

		// remember where we saw this character (overwrites the older position)
		lastSeen[ch] = right

		// the window is valid again: is it the longest so far?
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
	check("example 1", `lengthOfLongestSubstring("zxyzxyz")`,
		lengthOfLongestSubstring("zxyzxyz"), 3)

	// Example 2 from the problem: every letter is the same.
	check("example 2: all the same", `lengthOfLongestSubstring("xxxx")`,
		lengthOfLongestSubstring("xxxx"), 1)

	// Empty string.
	check("empty string", `lengthOfLongestSubstring("")`,
		lengthOfLongestSubstring(""), 0)

	// One letter.
	check("one letter", `lengthOfLongestSubstring("a")`,
		lengthOfLongestSubstring("a"), 1)

	// The classic example: "abc".
	check("classic abcabcbb", `lengthOfLongestSubstring("abcabcbb")`,
		lengthOfLongestSubstring("abcabcbb"), 3)

	// "wke" - note "pwke" is not a substring (letters must be next to each other).
	check("classic pwwkew", `lengthOfLongestSubstring("pwwkew")`,
		lengthOfLongestSubstring("pwwkew"), 3)

	// No letter repeats: the whole string.
	check("no repeats", `lengthOfLongestSubstring("abcdef")`,
		lengthOfLongestSubstring("abcdef"), 6)

	// Tricky: after "abb" the left side must NOT jump back to the old 'a'.
	check("abba", `lengthOfLongestSubstring("abba")`,
		lengthOfLongestSubstring("abba"), 2)

	// Tricky: the best window "vdf" starts in the middle.
	check("dvdf", `lengthOfLongestSubstring("dvdf")`,
		lengthOfLongestSubstring("dvdf"), 3)

	// A space is a character like any other.
	check("with spaces", `lengthOfLongestSubstring("a b c")`,
		lengthOfLongestSubstring("a b c"), 3)

	// "mzuxt" is the longest.
	check("tmmzuxt", `lengthOfLongestSubstring("tmmzuxt")`,
		lengthOfLongestSubstring("tmmzuxt"), 5)

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
