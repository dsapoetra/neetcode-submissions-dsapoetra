// Problem:    Subsets
// Difficulty: Medium
// Link:       https://neetcode.io/problems/subsets
// Source:     Data Structures & Algorithms/subsets/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a list of different numbers, return all possible subsets (every way
// to pick some of the numbers, including none and all). No subset may
// appear twice; the order does not matter.
//
// Example 1:
//
//	Input:  nums = [1, 2, 3]
//	Output: [[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]]
//
// Example 2:
//
//	Input:  nums = [7]
//	Output: [[], [7]]
//
// How to run (from the "local" folder):
//
//	go run ./065-subsets
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"sort"
)

// ===================== Solution =====================

// Idea: start with only the empty subset. Then take the numbers one by
// one: for every subset we already have, make a copy that ALSO contains
// the new number. That doubles the list each time.
// Example with [1, 2]:
//
//	start:    []
//	add 1:    [], [1]
//	add 2:    [], [1], [2], [1,2]
//
// Time: O(n * 2^n), there are 2^n subsets.
func subsets(nums []int) [][]int {
	res := [][]int{{}} // just the empty subset

	for _, num := range nums {
		count := len(res) // only loop over the subsets that exist right now
		for i := 0; i < count; i++ {
			// copy the old subset and put the new number at the end
			withNum := make([]int, len(res[i]), len(res[i])+1)
			copy(withNum, res[i])
			withNum = append(withNum, num)
			res = append(res, withNum)
		}
	}
	return res
}

// ===================== Test cases =====================

func main() {
	// The subsets may come back in any order, so we sort them before comparing.

	// Example 1 from the problem: 3 numbers give 8 subsets.
	check("example 1", `sortRows(subsets([]int{1, 2, 3}))`,
		sortRows(subsets([]int{1, 2, 3})), [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}})

	// Example 2 from the problem: the empty subset and the number itself.
	check("example 2: one number", `sortRows(subsets([]int{7}))`,
		sortRows(subsets([]int{7})), [][]int{{}, {7}})

	// No numbers: only the empty subset.
	check("no numbers", `sortRows(subsets([]int{}))`,
		sortRows(subsets([]int{})), [][]int{{}})

	// Two numbers.
	check("two numbers", `sortRows(subsets([]int{1, 2}))`,
		sortRows(subsets([]int{1, 2})), [][]int{{}, {1}, {1, 2}, {2}})

	// Negative numbers.
	check("negative numbers", `sortRows(subsets([]int{-1, 0}))`,
		sortRows(subsets([]int{-1, 0})), [][]int{{}, {-1}, {-1, 0}, {0}})

	// n numbers always give 2^n subsets.
	check("4 numbers give 16 subsets", `len(subsets([]int{1, 2, 3, 4}))`,
		len(subsets([]int{1, 2, 3, 4})), 16)
	check("10 numbers give 1024 subsets", `len(subsets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))`,
		len(subsets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})), 1024)

	report()
}

// ===================== Helpers for the tests =====================

// sortRows sorts the numbers inside every row and then sorts the rows.
// We use it when the problem says "return the answer in any order":
// after sorting, two correct answers always look the same.
func sortRows(rows [][]int) [][]int {
	for _, row := range rows {
		sort.Ints(row)
	}
	sort.Slice(rows, func(a, b int) bool {
		for i := 0; i < len(rows[a]) && i < len(rows[b]); i++ {
			if rows[a][i] != rows[b][i] {
				return rows[a][i] < rows[b][i]
			}
		}
		return len(rows[a]) < len(rows[b])
	})
	return rows
}

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
