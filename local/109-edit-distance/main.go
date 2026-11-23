// Problem:    Edit Distance
// Difficulty: Medium
// Link:       https://neetcode.io/problems/edit-distance
// Source:     Data Structures & Algorithms/edit-distance/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given two words, return the smallest number of steps needed to turn
// word1 into word2. One step is: insert a letter, delete a letter, or
// replace a letter.
//
// Example 1:
//
//	Input:  word1 = "monkeys", word2 = "money"
//	Output: 2             (delete "k", delete "s")
//
// Example 2:
//
//	Input:  word1 = "neatcdee", word2 = "neetcode"
//	Output: 3
//
// How to run (from the "local" folder):
//
//	go run ./109-edit-distance
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (dynamic programming):
// dist[i][j] = fewest steps to turn the first i letters of word1 into the
// first j letters of word2.
//   - if the last letters are equal, nothing to do:  dist[i-1][j-1]
//   - otherwise 1 step plus the best of:
//     replace the last letter:  dist[i-1][j-1]
//     delete from word1:        dist[i-1][j]
//     insert into word1:        dist[i][j-1]
//
// Start values: turning i letters into nothing takes i deletes, turning
// nothing into j letters takes j inserts.
// Time: O(len(word1) * len(word2)).
func minDistance(word1 string, word2 string) int {
	dist := make([][]int, len(word1)+1)
	for i := range dist {
		dist[i] = make([]int, len(word2)+1)
		dist[i][0] = i // delete everything
	}
	for j := 0; j <= len(word2); j++ {
		dist[0][j] = j // insert everything
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dist[i][j] = dist[i-1][j-1] // same letter, no step needed
				continue
			}

			best := dist[i-1][j-1] // replace
			if dist[i-1][j] < best {
				best = dist[i-1][j] // delete
			}
			if dist[i][j-1] < best {
				best = dist[i][j-1] // insert
			}
			dist[i][j] = best + 1
		}
	}
	return dist[len(word1)][len(word2)]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: delete "k" and "s".
	check("example 1", `minDistance("monkeys", "money")`,
		minDistance("monkeys", "money"), 2)

	// Example 2 from the problem.
	check("example 2", `minDistance("neatcdee", "neetcode")`,
		minDistance("neatcdee", "neetcode"), 3)

	// The classic examples.
	check("classic 1", `minDistance("horse", "ros")`,
		minDistance("horse", "ros"), 3)
	check("classic 2", `minDistance("intention", "execution")`,
		minDistance("intention", "execution"), 5)

	// Empty words.
	check("both empty", `minDistance("", "")`,
		minDistance("", ""), 0)
	check("insert everything", `minDistance("", "abc")`,
		minDistance("", "abc"), 3)
	check("delete everything", `minDistance("abc", "")`,
		minDistance("abc", ""), 3)

	// The same word.
	check("same word", `minDistance("abc", "abc")`,
		minDistance("abc", "abc"), 0)

	// One replace.
	check("one replace", `minDistance("a", "b")`,
		minDistance("a", "b"), 1)

	// A well-known example: k->s, e->i, add g.
	check("kitten -> sitting", `minDistance("kitten", "sitting")`,
		minDistance("kitten", "sitting"), 3)

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
