// Problem:    Alien Dictionary
// Difficulty: Hard
// Link:       https://neetcode.io/problems/foreign-dictionary
// Source:     Data Structures & Algorithms/foreign-dictionary/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// A foreign language uses the letters a-z, but in a different order. You
// get a list of words that is sorted by the rules of that language. Return
// a string with the letters in the order of that language. If the list is
// not possible, return "". If several orders fit, any of them is correct.
//
// Example 1:
//
//	Input:  words = ["z", "o"]
//	Output: "zo"
//
// Example 2:
//
//	Input:  words = ["hrn", "hrf", "er", "enn", "rfnn"]
//	Output: "hernf"
//
// How to run (from the "local" folder):
//
//	go run ./145-foreign-dictionary
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// ===================== Solution =====================

// Idea: the words are sorted in the alien order. Compare each word with
// the next one: the FIRST letter where they differ tells us one rule,
// "letter x comes before letter y".
// The rules form a graph, and a valid alphabet is a topological sort
// of it (same BFS method as Course Schedule). If we cannot place every
// letter, the rules contain a cycle, so there is no valid order.
func foreignDictionary(words []string) string {
	after := map[byte][]byte{} // after[x] = letters that must come after x
	need := map[byte]int{}     // need[y] = how many letters must come before y

	// register every letter that appears
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			need[word[i]] = 0
		}
	}

	// get one rule from each pair of neighboring words
	for i := 0; i < len(words)-1; i++ {
		first := words[i]
		second := words[i+1]

		shorter := len(first)
		if len(second) < shorter {
			shorter = len(second)
		}

		foundDifference := false
		for j := 0; j < shorter; j++ {
			if first[j] != second[j] {
				// first[j] comes before second[j]
				after[first[j]] = append(after[first[j]], second[j])
				need[second[j]]++
				foundDifference = true
				break
			}
		}

		// like "apple" before "app": that is never a valid order
		if !foundDifference && len(first) > len(second) {
			return ""
		}
	}

	// start with the letters that have nothing before them
	queue := []byte{}
	for letter, count := range need {
		if count == 0 {
			queue = append(queue, letter)
		}
	}

	order := []byte{}
	for len(queue) > 0 {
		letter := queue[0]
		queue = queue[1:]
		order = append(order, letter)

		// every letter after this one now waits for one less letter
		for _, next := range after[letter] {
			need[next]--
			if need[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// some letters were never placed -> there is a cycle
	if len(order) != len(need) {
		return ""
	}
	return string(order)
}

// ===================== Test cases =====================

func main() {
	// sortLetters sorts the letters of a string. We use it when MANY orders
	// are correct: then we only check that the right letters are in the answer.
	sortLetters := func(s string) string {
		letters := strings.Split(s, "")
		sort.Strings(letters)
		return strings.Join(letters, "")
	}

	// Example 1 from the problem: "z" comes before "o".
	check("example 1", `foreignDictionary([]string{"z", "o"})`,
		foreignDictionary([]string{"z", "o"}), "zo")

	// Example 2 from the problem: the rules are h<e, e<r, r<n, n<f.
	check("example 2", `foreignDictionary([]string{"hrn", "hrf", "er", "enn", "rfnn"})`,
		foreignDictionary([]string{"hrn", "hrf", "er", "enn", "rfnn"}), "hernf")

	// The classic example: w<e, e<r, r<t, t<f.
	check("classic", `foreignDictionary([]string{"wrt", "wrf", "er", "ett", "rftt"})`,
		foreignDictionary([]string{"wrt", "wrf", "er", "ett", "rftt"}), "wertf")

	// A longer word before its own prefix is never a valid order.
	check("word before its prefix", `foreignDictionary([]string{"abc", "ab"})`,
		foreignDictionary([]string{"abc", "ab"}), "")

	// a<b and b<a at the same time: a cycle, so no valid order.
	check("cycle", `foreignDictionary([]string{"a", "b", "a"})`,
		foreignDictionary([]string{"a", "b", "a"}), "")

	// The same word twice gives no rule, but the letter must be in the answer.
	check("same word twice", `foreignDictionary([]string{"z", "z"})`,
		foreignDictionary([]string{"z", "z"}), "z")

	// Three one-letter words: a<b, b<c.
	check("simple chain", `foreignDictionary([]string{"a", "b", "c"})`,
		foreignDictionary([]string{"a", "b", "c"}), "abc")

	// One word gives no rules: any order of its letters is correct.
	check("one word: all letters present", `sortLetters(foreignDictionary([]string{"abc"}))`,
		sortLetters(foreignDictionary([]string{"abc"})), "abc")

	// A prefix before the longer word is fine, and gives no rules.
	check("prefix first: all letters present", `sortLetters(foreignDictionary([]string{"ab", "abc"}))`,
		sortLetters(foreignDictionary([]string{"ab", "abc"})), "abc")

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
