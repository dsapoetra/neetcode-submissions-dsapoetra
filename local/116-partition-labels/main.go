// Problem:    Partition Labels
// Difficulty: Medium
// Link:       https://neetcode.io/problems/partition-labels
// Source:     Data Structures & Algorithms/partition-labels/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Cut the string into as many parts as possible so that every letter
// appears in only ONE part. Return the lengths of the parts, in order.
//
// Example 1:
//
//	Input:  s = "xyxxyzbzbbisl"
//	Output: [5, 5, 1, 1, 1]      ("xyxxy", "zbzbb", "i", "s", "l")
//
// Example 2:
//
//	Input:  s = "abcabc"
//	Output: [6]
//
// How to run (from the "local" folder):
//
//	go run ./116-partition-labels
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: a part must contain ALL copies of every letter in it. So first
// remember the last position of each letter.
// Then walk through the string. The current part must reach at least to
// the last position of every letter we have seen in it (end).
// When we arrive at "end", no letter of this part appears later, so we
// can cut here.
// Time: O(n).
func partitionLabels(s string) []int {
	var last [26]int // last[x] = last index of letter x
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

	res := []int{}
	start := 0 // where the current part begins
	end := 0   // where the current part must at least end

	for i := 0; i < len(s); i++ {
		if last[s[i]-'a'] > end {
			end = last[s[i]-'a']
		}

		if i == end {
			res = append(res, end-start+1) // cut here
			start = i + 1
		}
	}
	return res
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: "xyxxy", "zbzbb", "i", "s", "l".
	check("example 1", `partitionLabels("xyxxyzbzbbisl")`,
		partitionLabels("xyxxyzbzbbisl"), []int{5, 5, 1, 1, 1})

	// Example 2 from the problem: every letter appears at both ends.
	check("example 2", `partitionLabels("abcabc")`,
		partitionLabels("abcabc"), []int{6})

	// The classic example.
	check("classic", `partitionLabels("ababcbacadefegdehijhklij")`,
		partitionLabels("ababcbacadefegdehijhklij"), []int{9, 7, 8})

	// One letter.
	check("one letter", `partitionLabels("a")`,
		partitionLabels("a"), []int{1})

	// All different: every letter is its own part.
	check("all different", `partitionLabels("abc")`,
		partitionLabels("abc"), []int{1, 1, 1})

	// All the same.
	check("all the same", `partitionLabels("aaa")`,
		partitionLabels("aaa"), []int{3})

	// The first letter comes back at the end.
	check("first letter at the end", `partitionLabels("abca")`,
		partitionLabels("abca"), []int{4})

	// "c" is alone, the rest is chained together.
	check("chained letters", `partitionLabels("caedbdedda")`,
		partitionLabels("caedbdedda"), []int{1, 9})

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
