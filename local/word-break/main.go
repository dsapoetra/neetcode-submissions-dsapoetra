// Problem: Word Break
// Link:    https://neetcode.io/problems/word-break
// Source:  Data Structures & Algorithms/word-break/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./word-break
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (dynamic programming): canBreak[i] = true if the first i letters
// of s can be split into dictionary words.
// The first i letters work if some word w ends exactly at i AND the
// letters before that word (the first i-len(w) letters) also work.
// Time: O(len(s) * number of words * word length).
func wordBreak(s string, wordDict []string) bool {
	canBreak := make([]bool, len(s)+1)
	canBreak[0] = true // the empty string needs no words

	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			start := i - len(word) // where this word would begin
			if start >= 0 && canBreak[start] && s[start:i] == word {
				canBreak[i] = true
				break // one way is enough
			}
		}
	}
	return canBreak[len(s)]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: "neet" + "code".
	check("example 1", `wordBreak("neetcode", []string{"neet", "code"})`,
		wordBreak("neetcode", []string{"neet", "code"}), true)

	// Example 2 from the problem: "apple" is used twice.
	check("example 2: a word used twice", `wordBreak("applepenapple", []string{"apple", "pen", "ape"})`,
		wordBreak("applepenapple", []string{"apple", "pen", "ape"}), true)

	// Example 3 from the problem.
	check("example 3: impossible", `wordBreak("catsincars", []string{"cats", "cat", "sin", "in", "car"})`,
		wordBreak("catsincars", []string{"cats", "cat", "sin", "in", "car"}), false)

	// The classic "true" example.
	check("classic true", `wordBreak("leetcode", []string{"leet", "code"})`,
		wordBreak("leetcode", []string{"leet", "code"}), true)

	// The classic "false" example: "cats and og" / "cat sand og" both leave "og".
	check("classic false", `wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})`,
		wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}), false)

	// One letter.
	check("one letter, in the dictionary", `wordBreak("a", []string{"a"})`,
		wordBreak("a", []string{"a"}), true)
	check("one letter, not in the dictionary", `wordBreak("a", []string{"b"})`,
		wordBreak("a", []string{"b"}), false)

	// 7 letters = "aaaa" + "aaa".
	check("two different lengths", `wordBreak("aaaaaaa", []string{"aaaa", "aaa"})`,
		wordBreak("aaaaaaa", []string{"aaaa", "aaa"}), true)

	// 7 letters cannot be built from pieces of 2.
	check("odd length from pairs", `wordBreak("aaaaaaa", []string{"aa"})`,
		wordBreak("aaaaaaa", []string{"aa"}), false)

	// Tricky: starting with the longer word "car" leaves "s"; "ca" + "rs" works.
	check("shorter word first", `wordBreak("cars", []string{"car", "ca", "rs"})`,
		wordBreak("cars", []string{"car", "ca", "rs"}), true)

	// The empty string needs no words.
	check("empty string", `wordBreak("", []string{"a"})`,
		wordBreak("", []string{"a"}), true)

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
