// Problem: Longest Repeating Character Replacement
// Link:    https://neetcode.io/problems/longest-repeating-substring-with-replacement
// Source:  Data Structures & Algorithms/longest-repeating-substring-with-replacement/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./longest-repeating-substring-with-replacement
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (sliding window): a window is OK if we can make all its letters
// the same with at most k changes. The cheapest way is to keep the most
// common letter and change the rest, so:
//
//	changes needed = window length - count of most common letter
//
// Grow the window to the right; while it needs more than k changes,
// shrink it from the left.
// Time: O(26 * n).
func characterReplacement(s string, k int) int {
	count := make([]int, 26) // count[i] = how many of letter 'A'+i in the window
	left := 0
	best := 0

	for right := 0; right < len(s); right++ {
		// add the new letter to the window
		count[s[right]-'A']++

		// shrink while the window needs more than k changes
		for (right-left+1)-mostCommon(count) > k {
			count[s[left]-'A']--
			left++
		}

		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

// mostCommon returns the biggest value in count.
func mostCommon(count []int) int {
	biggest := 0
	for _, c := range count {
		if c > biggest {
			biggest = c
		}
	}
	return biggest
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: change both X (or both Y).
	check("example 1", characterReplacement("XYYX", 2), 4)

	// Example 2 from the problem: "AAABA" with the B changed.
	check("example 2", characterReplacement("AAABABB", 1), 5)

	// One letter, no changes.
	check("one letter", characterReplacement("A", 0), 1)

	// The classic example: change both A or both B.
	check("classic ABAB", characterReplacement("ABAB", 2), 4)

	// The classic example: "AABA" -> "AAAA" (or "BABB" -> "BBBB").
	check("classic AABABBA", characterReplacement("AABABBA", 1), 4)

	// Already all the same.
	check("already the same", characterReplacement("AAAA", 0), 4)

	// All different and no changes allowed.
	check("all different, k = 0", characterReplacement("ABCD", 0), 1)

	// All different, one change.
	check("all different, k = 1", characterReplacement("ABCD", 1), 2)

	// More changes than letters: the whole string.
	check("k bigger than the string", characterReplacement("ABCDE", 10), 5)

	// No changes allowed: the longest run of equal letters.
	check("longest run, k = 0", characterReplacement("AABBB", 0), 3)

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
