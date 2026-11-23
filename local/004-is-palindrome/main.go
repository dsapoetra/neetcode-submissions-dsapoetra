// Problem:    Valid Palindrome
// Difficulty: Easy
// Link:       https://neetcode.io/problems/is-palindrome
// Source:     Data Structures & Algorithms/is-palindrome/submission-5.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a string, return true if it is a palindrome when you only look at
// letters and digits and ignore upper/lower case.
//
// Example 1:
//
//	Input:  s = "Was it a car or a cat I saw?"
//	Output: true
//
// Example 2:
//
//	Input:  s = "tab a cat"
//	Output: false
//
// How to run (from the "local" folder):
//
//	go run ./004-is-palindrome
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"unicode"
)

// ===================== Solution =====================

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlphaNum(rune(s[l])) {
			l++
		}
		for r > l && !isAlphaNum(rune(s[r])) {
			r--
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: spaces and '?' are ignored, upper/lower case too.
	check("example 1", `isPalindrome("Was it a car or a cat I saw?")`,
		isPalindrome("Was it a car or a cat I saw?"), true)

	// Example 2 from the problem: "tabacat" is not a palindrome.
	check("example 2", `isPalindrome("tab a cat")`,
		isPalindrome("tab a cat"), false)

	// An empty string is a palindrome.
	check("empty string", `isPalindrome("")`,
		isPalindrome(""), true)

	// One letter.
	check("one letter", `isPalindrome("a")`,
		isPalindrome("a"), true)

	// The classic examples.
	check("classic true", `isPalindrome("A man, a plan, a canal: Panama")`,
		isPalindrome("A man, a plan, a canal: Panama"), true)
	check("classic false", `isPalindrome("race a car")`,
		isPalindrome("race a car"), false)

	// Digits count as characters: '0' is not 'P'.
	check("digit and letter", `isPalindrome("0P")`,
		isPalindrome("0P"), false)

	// Only punctuation: nothing is left to compare.
	check("only punctuation", `isPalindrome(".,")`,
		isPalindrome(".,"), true)

	// Only digits.
	check("digits", `isPalindrome("12321")`,
		isPalindrome("12321"), true)

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
