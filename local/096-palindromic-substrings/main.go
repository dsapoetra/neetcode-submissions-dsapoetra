// Problem:    Palindromic Substrings
// Difficulty: Medium
// Link:       https://neetcode.io/problems/palindromic-substrings
// Source:     Data Structures & Algorithms/palindromic-substrings/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a string, count how many of its pieces (letters next to each other)
// are palindromes. Pieces at different positions count separately, even if
// they look the same.
//
// Example 1:
//
//	Input:  s = "abc"
//	Output: 3             ("a", "b", "c")
//
// Example 2:
//
//	Input:  s = "aaa"
//	Output: 6             ("a", "a", "a", "aa", "aa", "aaa")
//
// How to run (from the "local" folder):
//
//	go run ./096-palindromic-substrings
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
	// every position can be a center
	for i := 0; i < len(s); i++ {
		count += countFromCenter(s, i, i)   // odd length palindromes
		count += countFromCenter(s, i, i+1) // even length palindromes
	}
	return count
}

// countFromCenter counts the palindromes that have this center.
func countFromCenter(s string, left int, right int) int {
	count := 0
	// grow outwards from the center; stop at the edge of the string, or when the two sides do not match
	for left >= 0 && right < len(s) && s[left] == s[right] {
		// s[left..right] is a palindrome
		count++
		// try one letter wider on both sides
		left--
		right++
	}
	return count
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: only the single letters.
	check("example 1: abc", `countSubstrings("abc")`,
		countSubstrings("abc"), 3)

	// Example 2 from the problem: a, a, a, aa, aa, aaa.
	check("example 2: aaa", `countSubstrings("aaa")`,
		countSubstrings("aaa"), 6)

	// One letter.
	check("one letter", `countSubstrings("a")`,
		countSubstrings("a"), 1)

	// Empty string.
	check("empty string", `countSubstrings("")`,
		countSubstrings(""), 0)

	// a, a, aa.
	check("aa", `countSubstrings("aa")`,
		countSubstrings("aa"), 3)

	// a, b, a, aba.
	check("aba", `countSubstrings("aba")`,
		countSubstrings("aba"), 4)

	// 4 single letters + bb + abba.
	check("abba", `countSubstrings("abba")`,
		countSubstrings("abba"), 6)

	// 7 single letters + cec + aceca + racecar.
	check("racecar", `countSubstrings("racecar")`,
		countSubstrings("racecar"), 10)

	// All different: only the single letters.
	check("all different", `countSubstrings("abcd")`,
		countSubstrings("abcd"), 4)

	// 4 + 3 + 2 + 1.
	check("aaaa", `countSubstrings("aaaa")`,
		countSubstrings("aaaa"), 10)

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
