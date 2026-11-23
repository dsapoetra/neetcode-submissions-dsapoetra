// Problem: Palindromic Substrings
// Link:    https://neetcode.io/problems/palindromic-substrings
// Source:  Data Structures & Algorithms/palindromic-substrings/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./palindromic-substrings
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: same as Longest Palindromic Substring. Try every center and grow
// outwards. Every time both sides match we found one more palindrome.
// Time: O(n^2).
func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count += countFromCenter(s, i, i)   // odd length palindromes
		count += countFromCenter(s, i, i+1) // even length palindromes
	}
	return count
}

// countFromCenter counts the palindromes that have this center.
func countFromCenter(s string, left int, right int) int {
	count := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: only the single letters.
	check("example 1: abc", countSubstrings("abc"), 3)

	// Example 2 from the problem: a, a, a, aa, aa, aaa.
	check("example 2: aaa", countSubstrings("aaa"), 6)

	// One letter.
	check("one letter", countSubstrings("a"), 1)

	// Empty string.
	check("empty string", countSubstrings(""), 0)

	// a, a, aa.
	check("aa", countSubstrings("aa"), 3)

	// a, b, a, aba.
	check("aba", countSubstrings("aba"), 4)

	// 4 single letters + bb + abba.
	check("abba", countSubstrings("abba"), 6)

	// 7 single letters + cec + aceca + racecar.
	check("racecar", countSubstrings("racecar"), 10)

	// All different: only the single letters.
	check("all different", countSubstrings("abcd"), 4)

	// 4 + 3 + 2 + 1.
	check("aaaa", countSubstrings("aaaa"), 10)

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
