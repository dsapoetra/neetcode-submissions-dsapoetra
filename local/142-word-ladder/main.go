// Problem:    Word Ladder
// Difficulty: Hard
// Link:       https://neetcode.io/problems/word-ladder
// Source:     Data Structures & Algorithms/word-ladder/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Change beginWord into endWord one letter at a time. Every word in
// between must be in wordList (and so must endWord). Return the number of
// words in the shortest chain, counting beginWord and endWord, or 0 if
// there is no chain.
//
// Example 1:
//
//	Input:  beginWord = "cat", endWord = "sag",
//	        wordList = ["bat", "bag", "sag", "dag", "dot"]
//	Output: 4             (cat -> bat -> bag -> sag)
//
// Example 2:
//
//	Input:  beginWord = "cat", endWord = "sag",
//	        wordList = ["bat", "bag", "sat", "dag", "dot"]
//	Output: 0
//
// How to run (from the "local" folder):
//
//	go run ./142-word-ladder
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (BFS): BFS finds the shortest way. Start with beginWord. In every
// step, go from each word to all unused words that differ by exactly one
// letter. The first time we reach endWord, the number of steps is the
// shortest possible.
// Time: O(n^2 * word length), fine for the small word lists here.
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// used[i] is true once wordList[i] is in the chain; a word is never
	// used twice (a second, later visit can only be a longer way)
	used := make([]bool, len(wordList))
	queue := []string{beginWord}
	length := 1 // number of words in the chain so far

	for len(queue) > 0 {
		// the words we can reach with one more step
		next := []string{}

		for _, word := range queue {
			// BFS reaches endWord for the first time on the shortest way
			if word == endWord {
				return length
			}

			// every unused word one letter away comes in the next round
			for i, candidate := range wordList {
				if !used[i] && differByOne(word, candidate) {
					used[i] = true
					next = append(next, candidate)
				}
			}
		}

		// nothing found in this round: continue with the words one step further
		queue = next
		length++
	}
	return 0 // endWord was never reached
}

// differByOne is true if the two words differ in exactly one position.
func differByOne(a string, b string) bool {
	// words of different length cannot differ by just one letter
	if len(a) != len(b) {
		return false
	}
	// count the positions where the letters are not the same
	different := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			different++
		}
	}
	return different == 1
}

// ===================== Test cases =====================

func main() {
	// The answer counts WORDS in the chain, including the first and the last.

	// Example 1 from the problem: cat -> bat -> bag -> sag.
	check("example 1", `ladderLength("cat", "sag", []string{"bat", "bag", "sag", "dag", "dot"})`,
		ladderLength("cat", "sag", []string{"bat", "bag", "sag", "dag", "dot"}), 4)

	// Example 2 from the problem: "sag" is not in the list.
	check("example 2: end word missing", `ladderLength("cat", "sag", []string{"bat", "bag", "sat", "dag", "dot"})`,
		ladderLength("cat", "sag", []string{"bat", "bag", "sat", "dag", "dot"}), 0)

	// The classic example: hit -> hot -> dot -> dog -> cog.
	check("classic", `ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})`,
		ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}), 5)

	// The classic example without "cog".
	check("classic, end word missing", `ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"})`,
		ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}), 0)

	// One step.
	check("one step", `ladderLength("hot", "dot", []string{"dot"})`,
		ladderLength("hot", "dot", []string{"dot"}), 2)

	// One-letter words: a -> c directly.
	check("one-letter words", `ladderLength("a", "c", []string{"a", "b", "c"})`,
		ladderLength("a", "c", []string{"a", "b", "c"}), 2)

	// "dog" is in the list, but there is no word between "hot" and "dog".
	check("no connection", `ladderLength("hot", "dog", []string{"hot", "dog"})`,
		ladderLength("hot", "dog", []string{"hot", "dog"}), 0)

	// Other words in the list are not needed.
	check("direct neighbor", `ladderLength("lost", "cost", []string{"most", "fist", "lost", "cost", "fish"})`,
		ladderLength("lost", "cost", []string{"most", "fist", "lost", "cost", "fish"}), 2)

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
