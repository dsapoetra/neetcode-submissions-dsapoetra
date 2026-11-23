// Problem:    Palindrome Partitioning
// Difficulty: Medium
// Link:       https://neetcode.io/problems/palindrome-partitioning
// Source:     Data Structures & Algorithms/palindrome-partitioning/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a string, cut it into pieces so that every piece is a palindrome.
// Return all possible ways to do that.
//
// Example 1:
//
//	Input:  s = "aab"
//	Output: [["a", "a", "b"], ["aa", "b"]]
//
// Example 2:
//
//	Input:  s = "a"
//	Output: [["a"]]
//
// How to run (from the "local" folder):
//
//	go run ./072-palindrome-partitioning
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (backtracking): cut the string from left to right. At position
// "start", try every possible first piece s[start..end]. If that piece is
// a palindrome, keep it and cut the rest of the string the same way.
// When we reach the end of the string, the pieces form one answer.
func partition(s string) [][]string {
	res := [][]string{}
	current := []string{}

	var build func(start int)
	build = func(start int) {
		if start == len(s) {
			// the whole string is cut: save a copy of the pieces
			res = append(res, append([]string{}, current...))
			return
		}

		for end := start; end < len(s); end++ {
			if isPalindromePiece(s, start, end) {
				current = append(current, s[start:end+1]) // choose this piece
				build(end + 1)                            // cut the rest
				current = current[:len(current)-1]        // undo the choice
			}
		}
	}

	build(0)
	return res
}

// isPalindromePiece checks if s[left..right] reads the same in both directions.
func isPalindromePiece(s string, left int, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", `partition("aab")`,
		partition("aab"), [][]string{{"a", "a", "b"}, {"aa", "b"}})

	// Example 2 from the problem: one letter.
	check("example 2: one letter", `partition("a")`,
		partition("a"), [][]string{{"a"}})

	// The whole string is a palindrome too.
	check("aba", `partition("aba")`,
		partition("aba"), [][]string{{"a", "b", "a"}, {"aba"}})

	// No palindrome longer than one letter.
	check("abc", `partition("abc")`,
		partition("abc"), [][]string{{"a", "b", "c"}})

	// Every cut works.
	check("aaa", `partition("aaa")`,
		partition("aaa"), [][]string{{"a", "a", "a"}, {"a", "aa"}, {"aa", "a"}, {"aaa"}})

	// An even palindrome.
	check("abba", `partition("abba")`,
		partition("abba"), [][]string{{"a", "b", "b", "a"}, {"a", "bb", "a"}, {"abba"}})

	// 4 equal letters: 3 places to cut or not cut = 2*2*2 = 8 ways.
	check("aaaa gives 8", `len(partition("aaaa"))`,
		len(partition("aaaa")), 8)

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
