// Problem:    Letter Combinations of a Phone Number
// Difficulty: Medium
// Link:       https://neetcode.io/problems/combinations-of-a-phone-number
// Source:     Data Structures & Algorithms/combinations-of-a-phone-number/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// On a phone keypad every digit from 2 to 9 stands for some letters
// (2 = abc, 3 = def, 4 = ghi, 5 = jkl, 6 = mno, 7 = pqrs, 8 = tuv, 9 = wxyz).
// Given a string of digits, return every letter combination it could stand
// for. Return an empty list for an empty string.
//
// Example 1:
//
//	Input:  digits = "34"
//	Output: ["dg", "dh", "di", "eg", "eh", "ei", "fg", "fh", "fi"]
//
// Example 2:
//
//	Input:  digits = ""
//	Output: []
//
// How to run (from the "local" folder):
//
//	go run ./073-combinations-of-a-phone-number
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"sort"
)

// ===================== Solution =====================

// Idea: start with one empty string. For every digit, take all the
// strings we have so far and add each letter of that digit to each of
// them. Example "23": [""] -> ["a","b","c"] -> ["ad","ae","af","bd",...]
// Time: O(4^n), a digit has at most 4 letters.
func letterCombinations(digits string) []string {
	// no digits means no combinations (not one empty string)
	if digits == "" {
		return []string{}
	}

	// letters[d] = the letters on phone key d
	letters := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	// res = all the strings built so far; one empty string to start
	res := []string{""}
	for i := 0; i < len(digits); i++ {
		// digits[i] is a character like '2'; minus '0' gives the number 2
		options := letters[digits[i]-'0']

		// extend every string we have with every letter of this digit
		next := []string{}
		for _, prefix := range res {
			for _, letter := range options {
				next = append(next, prefix+string(letter))
			}
		}
		// the longer strings replace the old ones
		res = next
	}
	return res
}

// ===================== Test cases =====================

func main() {
	// The strings may come back in any order, so we sort them before comparing.

	// Example 1 from the problem: 3 = def, 4 = ghi.
	check("example 1", `sortedStrings(letterCombinations("34"))`,
		sortedStrings(letterCombinations("34")), []string{"dg", "dh", "di", "eg", "eh", "ei", "fg", "fh", "fi"})

	// Example 2 from the problem: no digits, no combinations.
	check("example 2: empty", `sortedStrings(letterCombinations(""))`,
		sortedStrings(letterCombinations("")), []string{})

	// One digit.
	check("one digit", `sortedStrings(letterCombinations("2"))`,
		sortedStrings(letterCombinations("2")), []string{"a", "b", "c"})

	// The classic example.
	check("classic", `sortedStrings(letterCombinations("23"))`,
		sortedStrings(letterCombinations("23")), []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"})

	// 7 and 9 have FOUR letters.
	check("7 has four letters", `sortedStrings(letterCombinations("7"))`,
		sortedStrings(letterCombinations("7")), []string{"p", "q", "r", "s"})
	check("9 has four letters", `sortedStrings(letterCombinations("9"))`,
		sortedStrings(letterCombinations("9")), []string{"w", "x", "y", "z"})

	// Counting: 4 * 4 and 3 * 3 * 3.
	check("79 gives 16", `len(letterCombinations("79"))`,
		len(letterCombinations("79")), 16)
	check("234 gives 27", `len(letterCombinations("234"))`,
		len(letterCombinations("234")), 27)

	// The same digit twice.
	check("same digit twice", `sortedStrings(letterCombinations("22"))`,
		sortedStrings(letterCombinations("22")), []string{"aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc"})

	report()
}

// ===================== Helpers for the tests =====================

// sortedStrings sorts the words. We use it when the answer may be in any order.
func sortedStrings(words []string) []string {
	sort.Strings(words)
	return words
}

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
