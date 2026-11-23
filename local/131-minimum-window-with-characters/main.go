// Problem:    Minimum Window Substring
// Difficulty: Hard
// Link:       https://neetcode.io/problems/minimum-window-with-characters
// Source:     Data Structures & Algorithms/minimum-window-with-characters/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given two strings s and t, return the shortest piece of s that contains
// every character of t (including duplicates). Return "" if there is none.
//
// Example 1:
//
//	Input:  s = "OUZODYXAZV", t = "XYZ"
//	Output: "YXAZ"
//
// Example 2:
//
//	Input:  s = "xyz", t = "xyz"
//	Output: "xyz"
//
// Example 3:
//
//	Input:  s = "x", t = "xy"
//	Output: ""
//
// How to run (from the "local" folder):
//
//	go run ./131-minimum-window-with-characters
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
	check("example 1", `minWindow("OUZODYXAZV", "XYZ")`,
		minWindow("OUZODYXAZV", "XYZ"), "YXAZ")

	// Example 2 from the problem: the whole string is the window.
	check("example 2: whole string", `minWindow("xyz", "xyz")`,
		minWindow("xyz", "xyz"), "xyz")

	// Example 3 from the problem: s has no 'y'.
	check("example 3: impossible", `minWindow("x", "xy")`,
		minWindow("x", "xy"), "")

	// The classic example.
	check("classic", `minWindow("ADOBECODEBANC", "ABC")`,
		minWindow("ADOBECODEBANC", "ABC"), "BANC")

	// One letter.
	check("one letter", `minWindow("a", "a")`,
		minWindow("a", "a"), "a")

	// t needs TWO a's, s has only one.
	check("not enough copies", `minWindow("a", "aa")`,
		minWindow("a", "aa"), "")

	// t needs two a's and s has exactly two.
	check("exactly enough copies", `minWindow("aa", "aa")`,
		minWindow("aa", "aa"), "aa")

	// The window is at the end.
	check("window at the end", `minWindow("ab", "b")`,
		minWindow("ab", "b"), "b")

	// The letter is not in s at all.
	check("letter not in s", `minWindow("abc", "d")`,
		minWindow("abc", "d"), "")

	// Upper and lower case are different letters.
	check("case matters", `minWindow("aA", "A")`,
		minWindow("aA", "A"), "A")

	// "bba" contains a and b, but the shorter "ba" does too.
	check("shrink from the left", `minWindow("bba", "ab")`,
		minWindow("bba", "ab"), "ba")

	// t needs three a's: the three at the end are closest together.
	check("three copies", `minWindow("aaflslflsldkalskaaa", "aaa")`,
		minWindow("aaflslflsldkalskaaa", "aaa"), "aaa")

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
