// Problem:    Permutations
// Difficulty: Medium
// Link:       https://neetcode.io/problems/permutations
// Source:     Data Structures & Algorithms/permutations/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a list of different numbers, return all possible orderings
// (permutations) of the list. The order of the answer does not matter.
//
// Example 1:
//
//	Input:  nums = [1, 2, 3]
//	Output: [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]]
//
// Example 2:
//
//	Input:  nums = [7]
//	Output: [[7]]
//
// How to run (from the "local" folder):
//
//	go run ./068-permutations
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

// Idea (backtracking): fill the permutation position by position. For each
// position try every number that is not used yet, go deeper, then take
// the number back out and try the next one.
// Time: O(n * n!), there are n! permutations.
func permute(nums []int) [][]int {
	res := [][]int{}
	current := []int{}
	used := make([]bool, len(nums))

	var build func()
	build = func() {
		if len(current) == len(nums) {
			// full permutation: save a COPY, because current keeps changing
			res = append(res, append([]int{}, current...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true // choose
			current = append(current, nums[i])

			build() // fill the next position

			current = current[:len(current)-1] // undo the choice
			used[i] = false
		}
	}

	build()
	return res
}

// ===================== Test cases =====================

func main() {
	// The permutations may come back in any order, so we sort the list of
	// permutations (not the numbers inside them) before comparing.

	// Example 1 from the problem.
	check("example 1", `sortLists(permute([]int{1, 2, 3}))`,
		sortLists(permute([]int{1, 2, 3})), [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}})

	// Example 2 from the problem: one number.
	check("example 2: one number", `sortLists(permute([]int{7}))`,
		sortLists(permute([]int{7})), [][]int{{7}})

	// Two numbers.
	check("two numbers", `sortLists(permute([]int{0, 1}))`,
		sortLists(permute([]int{0, 1})), [][]int{{0, 1}, {1, 0}})

	// Negative numbers.
	check("negative numbers", `sortLists(permute([]int{-1, 1}))`,
		sortLists(permute([]int{-1, 1})), [][]int{{-1, 1}, {1, -1}})

	// The input does not have to be sorted.
	check("unsorted input", `sortLists(permute([]int{3, 1}))`,
		sortLists(permute([]int{3, 1})), [][]int{{1, 3}, {3, 1}})

	// n numbers always give n! permutations.
	check("4 numbers give 24", `len(permute([]int{1, 2, 3, 4}))`,
		len(permute([]int{1, 2, 3, 4})), 24)
	check("5 numbers give 120", `len(permute([]int{1, 2, 3, 4, 5}))`,
		len(permute([]int{1, 2, 3, 4, 5})), 120)

	report()
}

// ===================== Helpers for the tests =====================

// sortLists sorts the rows, but does NOT change the order inside a row.
// We use it when the rows may come back in any order.
func sortLists(rows [][]int) [][]int {
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
