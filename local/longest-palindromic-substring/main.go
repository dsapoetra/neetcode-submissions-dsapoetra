// Problem: Longest Palindromic Substring
// Link:    https://neetcode.io/problems/longest-palindromic-substring
// Source:  Data Structures & Algorithms/longest-palindromic-substring/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./longest-palindromic-substring
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: a palindrome is a mirror around its center. Try every position
// as a center and grow outwards while the letters on both sides match.
// There are two kinds of center:
//   - one letter      (odd length, like "aba")
//   - between letters (even length, like "abba")
//
// Time: O(n^2).
func longestPalindrome(s string) string {
	bestStart := 0
	bestLen := 0

	for i := 0; i < len(s); i++ {
		// odd length: center is letter i
		left, right := growPalindrome(s, i, i)
		if right-left+1 > bestLen {
			bestStart = left
			bestLen = right - left + 1
		}

		// even length: center is between i and i+1
		left, right = growPalindrome(s, i, i+1)
		if right-left+1 > bestLen {
			bestStart = left
			bestLen = right - left + 1
		}
	}
	return s[bestStart : bestStart+bestLen]
}

// growPalindrome grows outwards from (left, right) while the letters match.
// It returns the first and last index of the palindrome it found.
func growPalindrome(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// the loop went one step too far on both sides
	return left + 1, right - 1
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: both "aba" and "bab" are correct answers.
	got := longestPalindrome("ababd")
	check("example 1: aba or bab", got == "aba" || got == "bab", true)

	// Example 2 from the problem: an even-length palindrome.
	check("example 2: even length", longestPalindrome("abbc"), "bb")

	// One letter is a palindrome.
	check("one letter", longestPalindrome("a"), "a")

	// Empty string.
	check("empty string", longestPalindrome(""), "")

	// The whole string is a palindrome.
	check("whole string", longestPalindrome("racecar"), "racecar")

	// All letters the same.
	check("all the same", longestPalindrome("aaaa"), "aaaa")

	// No palindrome longer than one letter: any single letter is correct.
	check("no long palindrome", len(longestPalindrome("abc")), 1)

	// A long even palindrome in the middle.
	check("in the middle", longestPalindrome("forgeeksskeegfor"), "geeksskeeg")

	// An odd palindrome in the middle.
	check("odd, in the middle", longestPalindrome("xabay"), "aba")

	// The palindrome is at the very end.
	check("at the end", longestPalindrome("noonracecar"), "racecar")

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
