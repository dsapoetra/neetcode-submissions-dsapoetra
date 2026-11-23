// Problem: Minimum Window Substring
// Link:    https://neetcode.io/problems/minimum-window-with-characters
// Source:  Data Structures & Algorithms/minimum-window-with-characters/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./minimum-window-with-characters
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (sliding window):
//
//	need[ch] = how many more of ch the window still needs
//	           (it becomes negative when the window has extra)
//	missing  = how many letters of t are still missing in total
//
// Grow the window to the right. When nothing is missing the window is
// valid, so shrink it from the left as long as it stays valid and
// remember the smallest one.
// Time: O(len(s) + len(t)).
func minWindow(s string, t string) string {
	need := map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	missing := len(t)

	bestStart := 0
	bestLen := 0 // 0 means we have not found any window yet
	left := 0

	for right := 0; right < len(s); right++ {
		// add s[right] to the window
		if need[s[right]] > 0 {
			missing-- // it was a letter we still needed
		}
		need[s[right]]--

		// the window has everything: try to make it smaller
		for missing == 0 {
			size := right - left + 1
			if bestLen == 0 || size < bestLen {
				bestStart = left
				bestLen = size
			}

			// remove s[left] from the window
			need[s[left]]++
			if need[s[left]] > 0 {
				missing++ // we just lost a letter we need
			}
			left++
		}
	}
	return s[bestStart : bestStart+bestLen]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", minWindow("OUZODYXAZV", "XYZ"), "YXAZ")

	// Example 2 from the problem: the whole string is the window.
	check("example 2: whole string", minWindow("xyz", "xyz"), "xyz")

	// Example 3 from the problem: s has no 'y'.
	check("example 3: impossible", minWindow("x", "xy"), "")

	// The classic example.
	check("classic", minWindow("ADOBECODEBANC", "ABC"), "BANC")

	// One letter.
	check("one letter", minWindow("a", "a"), "a")

	// t needs TWO a's, s has only one.
	check("not enough copies", minWindow("a", "aa"), "")

	// t needs two a's and s has exactly two.
	check("exactly enough copies", minWindow("aa", "aa"), "aa")

	// The window is at the end.
	check("window at the end", minWindow("ab", "b"), "b")

	// The letter is not in s at all.
	check("letter not in s", minWindow("abc", "d"), "")

	// Upper and lower case are different letters.
	check("case matters", minWindow("aA", "A"), "A")

	// "bba" contains a and b, but the shorter "ba" does too.
	check("shrink from the left", minWindow("bba", "ab"), "ba")

	// t needs three a's: the three at the end are closest together.
	check("three copies", minWindow("aaflslflsldkalskaaa", "aaa"), "aaa")

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
