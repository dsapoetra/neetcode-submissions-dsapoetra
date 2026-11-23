// Problem: Longest Common Subsequence
// Link:    https://neetcode.io/problems/longest-common-subsequence
// Source:  Data Structures & Algorithms/longest-common-subsequence/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./longest-common-subsequence
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (2D dynamic programming):
// lcs[i][j] = answer for the first i letters of text1 and the first j
// letters of text2.
//   - if the last letters are the same, use it: lcs[i-1][j-1] + 1
//   - otherwise drop one last letter from either text and take the better
//
// Time: O(len(text1) * len(text2)).
func longestCommonSubsequence(text1 string, text2 string) int {
	// one extra row and column of zeros for the "empty text" case
	lcs := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		lcs[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				// same letter: it extends the common subsequence
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else if lcs[i-1][j] > lcs[i][j-1] {
				lcs[i][j] = lcs[i-1][j] // drop last letter of text1
			} else {
				lcs[i][j] = lcs[i][j-1] // drop last letter of text2
			}
		}
	}
	return lcs[len(text1)][len(text2)]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: "cat".
	check("example 1", longestCommonSubsequence("cat", "crabt"), 3)

	// Example 2 from the problem: the texts are the same.
	check("example 2: same text", longestCommonSubsequence("abcd", "abcd"), 4)

	// Example 3 from the problem: nothing in common.
	check("example 3: nothing in common", longestCommonSubsequence("abcd", "efgh"), 0)

	// The classic example: "ace".
	check("classic", longestCommonSubsequence("abcde", "ace"), 3)

	// Single letters.
	check("single letters, same", longestCommonSubsequence("a", "a"), 1)
	check("single letters, different", longestCommonSubsequence("a", "b"), 0)

	// One text is empty.
	check("first empty", longestCommonSubsequence("", "abc"), 0)
	check("second empty", longestCommonSubsequence("abc", ""), 0)

	// The order matters: "ab" and "ba" share only one letter in order.
	check("order matters", longestCommonSubsequence("ab", "ba"), 1)

	// A well-known example: "gtab".
	check("gtab", longestCommonSubsequence("aggtab", "gxtxayb"), 4)

	// The whole first text is inside the second one.
	check("first inside second", longestCommonSubsequence("abcba", "abcbcba"), 5)

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
