// Problem:    Combination Sum II
// Difficulty: Medium
// Link:       https://neetcode.io/problems/combination-target-sum-ii
// Source:     Data Structures & Algorithms/combination-target-sum-ii/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a list of numbers (duplicates are possible) and a target, return
// every combination of numbers that adds up to the target. Each ELEMENT of
// the list can be used only once. The same combination must not appear
// twice; the order does not matter.
//
// Example 1:
//
//	Input:  candidates = [9, 2, 2, 4, 6, 1, 5], target = 8
//	Output: [[1, 2, 5], [2, 2, 4], [2, 6]]
//
// Example 2:
//
//	Input:  candidates = [1, 2, 3, 4, 5], target = 7
//	Output: [[1, 2, 4], [2, 5], [3, 4]]
//
// How to run (from the "local" folder):
//
//	go run ./067-combination-target-sum-ii
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

// Idea (backtracking): sort the numbers, then build combinations one
// number at a time. Each number can be used only once, so after choosing
// index i we continue from i+1.
// To avoid the same combination twice: at one position of the combination
// never try the same VALUE twice (skip equal neighbors in the sorted list).
// Time: O(2^n) in the worst case.
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	current := []int{}

	// build tries every number from index start on as the next number.
	var build func(start int, remaining int)
	build = func(start int, remaining int) {
		if remaining == 0 {
			// found one: save a COPY, because current keeps changing
			res = append(res, append([]int{}, current...))
			return
		}

		for i := start; i < len(candidates); i++ {
			// same value as the previous try at this position: skip
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			// sorted list: everything after this is too big as well
			if candidates[i] > remaining {
				break
			}

			current = append(current, candidates[i]) // choose
			build(i+1, remaining-candidates[i])      // go deeper
			current = current[:len(current)-1]       // undo the choice
		}
	}

	build(0, target)
	return res
}

// ===================== Test cases =====================

func main() {
	// The combinations may come back in any order, so we sort them before comparing.

	// Example 1 from the problem.
	check("example 1", `sortRows(combinationSum2([]int{9, 2, 2, 4, 6, 1, 5}, 8))`,
		sortRows(combinationSum2([]int{9, 2, 2, 4, 6, 1, 5}, 8)), [][]int{{1, 2, 5}, {2, 2, 4}, {2, 6}})

	// Example 2 from the problem.
	check("example 2", `sortRows(combinationSum2([]int{1, 2, 3, 4, 5}, 7))`,
		sortRows(combinationSum2([]int{1, 2, 3, 4, 5}, 7)), [][]int{{1, 2, 4}, {2, 5}, {3, 4}})

	// The classic example: there are two 1s, but [1,7] must appear only once.
	check("classic", `sortRows(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))`,
		sortRows(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)), [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}})

	// Three 2s: [1,2,2] must appear only once.
	check("many duplicates", `sortRows(combinationSum2([]int{2, 5, 2, 1, 2}, 5))`,
		sortRows(combinationSum2([]int{2, 5, 2, 1, 2}, 5)), [][]int{{1, 2, 2}, {5}})

	// One number that is the target.
	check("one number, match", `sortRows(combinationSum2([]int{1}, 1))`,
		sortRows(combinationSum2([]int{1}, 1)), [][]int{{1}})

	// A number cannot be used twice.
	check("one number, cannot reuse", `sortRows(combinationSum2([]int{1}, 2))`,
		sortRows(combinationSum2([]int{1}, 2)), [][]int{})

	// Three 1s: only ONE combination [1,1].
	check("three equal numbers", `sortRows(combinationSum2([]int{1, 1, 1}, 2))`,
		sortRows(combinationSum2([]int{1, 1, 1}, 2)), [][]int{{1, 1}})

	// Everything is too big.
	check("all too big", `sortRows(combinationSum2([]int{3, 5}, 1))`,
		sortRows(combinationSum2([]int{3, 5}, 1)), [][]int{})

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
