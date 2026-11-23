// Problem: Alien Dictionary
// Link:    https://neetcode.io/problems/foreign-dictionary
// Source:  Data Structures & Algorithms/foreign-dictionary/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./foreign-dictionary
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
	check("example 1", foreignDictionary([]string{"z", "o"}), "zo")

	// Example 2 from the problem: the rules are h<e, e<r, r<n, n<f.
	check("example 2", foreignDictionary([]string{"hrn", "hrf", "er", "enn", "rfnn"}), "hernf")

	// The classic example: w<e, e<r, r<t, t<f.
	check("classic", foreignDictionary([]string{"wrt", "wrf", "er", "ett", "rftt"}), "wertf")

	// A longer word before its own prefix is never a valid order.
	check("word before its prefix", foreignDictionary([]string{"abc", "ab"}), "")

	// a<b and b<a at the same time: a cycle, so no valid order.
	check("cycle", foreignDictionary([]string{"a", "b", "a"}), "")

	// The same word twice gives no rule, but the letter must be in the answer.
	check("same word twice", foreignDictionary([]string{"z", "z"}), "z")

	// Three one-letter words: a<b, b<c.
	check("simple chain", foreignDictionary([]string{"a", "b", "c"}), "abc")

	// One word gives no rules: any order of its letters is correct.
	check("one word: all letters present", sortLetters(foreignDictionary([]string{"abc"})), "abc")

	// A prefix before the longer word is fine, and gives no rules.
	check("prefix first: all letters present", sortLetters(foreignDictionary([]string{"ab", "abc"})), "abc")

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
