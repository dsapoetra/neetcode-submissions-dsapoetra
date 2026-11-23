// Problem:    Longest Common Subsequence
// Difficulty: Medium
// Link:       https://neetcode.io/problems/longest-common-subsequence
// Source:     Data Structures & Algorithms/longest-common-subsequence/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given two strings, return the length of their longest common subsequence:
// the longest string that can be made from BOTH by deleting letters (the
// order of the remaining letters must stay the same).
//
// Example 1:
//
//	Input:  text1 = "cat", text2 = "crabt"
//	Output: 3             ("cat")
//
// Example 2:
//
//	Input:  text1 = "abcd", text2 = "abcd"
//	Output: 4
//
// Example 3:
//
//	Input:  text1 = "abcd", text2 = "efgh"
//	Output: 0
//
// How to run (from the "local" folder):
//
//	go run ./104-longest-common-subsequence
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
	check("example 1", `longestCommonSubsequence("cat", "crabt")`,
		longestCommonSubsequence("cat", "crabt"), 3)

	// Example 2 from the problem: the texts are the same.
	check("example 2: same text", `longestCommonSubsequence("abcd", "abcd")`,
		longestCommonSubsequence("abcd", "abcd"), 4)

	// Example 3 from the problem: nothing in common.
	check("example 3: nothing in common", `longestCommonSubsequence("abcd", "efgh")`,
		longestCommonSubsequence("abcd", "efgh"), 0)

	// The classic example: "ace".
	check("classic", `longestCommonSubsequence("abcde", "ace")`,
		longestCommonSubsequence("abcde", "ace"), 3)

	// Single letters.
	check("single letters, same", `longestCommonSubsequence("a", "a")`,
		longestCommonSubsequence("a", "a"), 1)
	check("single letters, different", `longestCommonSubsequence("a", "b")`,
		longestCommonSubsequence("a", "b"), 0)

	// One text is empty.
	check("first empty", `longestCommonSubsequence("", "abc")`,
		longestCommonSubsequence("", "abc"), 0)
	check("second empty", `longestCommonSubsequence("abc", "")`,
		longestCommonSubsequence("abc", ""), 0)

	// The order matters: "ab" and "ba" share only one letter in order.
	check("order matters", `longestCommonSubsequence("ab", "ba")`,
		longestCommonSubsequence("ab", "ba"), 1)

	// A well-known example: "gtab".
	check("gtab", `longestCommonSubsequence("aggtab", "gxtxayb")`,
		longestCommonSubsequence("aggtab", "gxtxayb"), 4)

	// The whole first text is inside the second one.
	check("first inside second", `longestCommonSubsequence("abcba", "abcbcba")`,
		longestCommonSubsequence("abcba", "abcbcba"), 5)

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
