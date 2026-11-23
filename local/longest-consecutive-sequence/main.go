// Problem: Longest Consecutive Sequence
// Link:    https://neetcode.io/problems/longest-consecutive-sequence
// Source:  Data Structures & Algorithms/longest-consecutive-sequence/submission-10.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./longest-consecutive-sequence
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

func longestConsecutive(nums []int) int {
	seen := make(map[int]bool)
	for _, n := range nums {
		seen[n] = true
	}

	longest := 0
	for n := range seen {
		// Only start counting from the beginning of a run.
		if seen[n-1] {
			continue
		}

		length := 1
		for seen[n+length] {
			length++
		}

		if length > longest {
			longest = length
		}
	}

	return longest
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 2, 3, 4, 5.
	check("example 1", `longestConsecutive([]int{2, 20, 4, 10, 3, 4, 5})`,
		longestConsecutive([]int{2, 20, 4, 10, 3, 4, 5}), 4)

	// Example 2 from the problem: 0..6, the duplicate 1 does not count twice.
	check("example 2", `longestConsecutive([]int{0, 3, 2, 5, 4, 6, 1, 1})`,
		longestConsecutive([]int{0, 3, 2, 5, 4, 6, 1, 1}), 7)

	// No numbers: length 0.
	check("empty list", `longestConsecutive([]int{})`,
		longestConsecutive([]int{}), 0)

	// One number is a sequence of length 1.
	check("one number", `longestConsecutive([]int{5})`,
		longestConsecutive([]int{5}), 1)

	// Duplicates only: still length 1.
	check("all the same", `longestConsecutive([]int{1, 1, 1})`,
		longestConsecutive([]int{1, 1, 1}), 1)

	// The classic example: 1, 2, 3, 4.
	check("classic", `longestConsecutive([]int{100, 4, 200, 1, 3, 2})`,
		longestConsecutive([]int{100, 4, 200, 1, 3, 2}), 4)

	// The sequence crosses zero.
	check("negative to positive", `longestConsecutive([]int{-1, 0, 1})`,
		longestConsecutive([]int{-1, 0, 1}), 3)

	// No two numbers are neighbors.
	check("no neighbors", `longestConsecutive([]int{10, 20, 30})`,
		longestConsecutive([]int{10, 20, 30}), 1)

	// Two sequences (-1,0 and 3..9): the longer one wins.
	check("two sequences", `longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6})`,
		longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}), 7)

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
