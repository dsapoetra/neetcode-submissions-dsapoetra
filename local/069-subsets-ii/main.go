// Problem:    Subsets II
// Difficulty: Medium
// Link:       https://neetcode.io/problems/subsets-ii
// Source:     Data Structures & Algorithms/subsets-ii/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a list of numbers that may contain duplicates, return all possible
// subsets. The same subset must not appear twice; the order does not matter.
//
// Example 1:
//
//	Input:  nums = [1, 2, 1]
//	Output: [[], [1], [1, 2], [1, 1], [1, 2, 1], [2]]
//
// Example 2:
//
//	Input:  nums = [7, 7]
//	Output: [[], [7], [7, 7]]
//
// How to run (from the "local" folder):
//
//	go run ./069-subsets-ii
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

// Idea (backtracking): sort the numbers so equal numbers are next to each
// other. Every state of "current" is one subset, so save it at every step.
// Then try to add each later number.
// To avoid the same subset twice: at one position never try the same
// VALUE twice (skip equal neighbors).
// Time: O(n * 2^n).
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	current := []int{}

	var build func(start int)
	build = func(start int) {
		// current is a subset: save a copy
		res = append(res, append([]int{}, current...))

		for i := start; i < len(nums); i++ {
			// same value as the previous try at this position: skip
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			current = append(current, nums[i]) // choose
			build(i + 1)                       // go deeper
			current = current[:len(current)-1] // undo the choice
		}
	}

	build(0)
	return res
}

// ===================== Test cases =====================

func main() {
	// The subsets may come back in any order, so we sort them before comparing.

	// Example 1 from the problem: two 1s, but [1] and [1,2] appear only once.
	check("example 1", `sortRows(subsetsWithDup([]int{1, 2, 1}))`,
		sortRows(subsetsWithDup([]int{1, 2, 1})), [][]int{{}, {1}, {1, 1}, {1, 1, 2}, {1, 2}, {2}})

	// Example 2 from the problem.
	check("example 2", `sortRows(subsetsWithDup([]int{7, 7}))`,
		sortRows(subsetsWithDup([]int{7, 7})), [][]int{{}, {7}, {7, 7}})

	// The classic example.
	check("classic", `sortRows(subsetsWithDup([]int{1, 2, 2}))`,
		sortRows(subsetsWithDup([]int{1, 2, 2})), [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}})

	// One number.
	check("one number", `sortRows(subsetsWithDup([]int{0}))`,
		sortRows(subsetsWithDup([]int{0})), [][]int{{}, {0}})

	// Three equal numbers: one subset per size.
	check("three equal numbers", `sortRows(subsetsWithDup([]int{1, 1, 1}))`,
		sortRows(subsetsWithDup([]int{1, 1, 1})), [][]int{{}, {1}, {1, 1}, {1, 1, 1}})

	// No duplicates: the same as the normal Subsets problem.
	check("no duplicates", `sortRows(subsetsWithDup([]int{1, 2}))`,
		sortRows(subsetsWithDup([]int{1, 2})), [][]int{{}, {1}, {1, 2}, {2}})

	// One 1 (use it or not) and four 4s (use 0 to 4 of them): 2 * 5 = 10.
	check("counting", `len(subsetsWithDup([]int{4, 4, 4, 1, 4}))`,
		len(subsetsWithDup([]int{4, 4, 4, 1, 4})), 10)

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
