// Problem: Combination Sum
// Link:    https://neetcode.io/problems/combination-target-sum
// Source:  Data Structures & Algorithms/combination-target-sum/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./combination-target-sum
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

// Idea (backtracking): build a combination one number at a time.
// At each step try every number from index "start" onwards:
//
//	pick it -> go deeper -> un-pick it (backtrack) and try the next one.
//
// We pass the same index again when going deeper because a number can
// be used many times. Starting from "start" avoids duplicates like
// [2,3] and [3,2].
func combinationSum(nums []int, target int) [][]int {
	res := [][]int{}
	current := []int{} // the combination we are building

	var backtrack func(start int, remaining int)
	backtrack = func(start int, remaining int) {
		// we hit the target exactly: save a COPY of current
		if remaining == 0 {
			saved := make([]int, len(current))
			copy(saved, current)
			res = append(res, saved)
			return
		}

		for i := start; i < len(nums); i++ {
			// this number is too big, skip it
			if nums[i] > remaining {
				continue
			}
			current = append(current, nums[i]) // pick
			backtrack(i, remaining-nums[i])    // go deeper (i again: reuse allowed)
			current = current[:len(current)-1] // un-pick
		}
	}

	backtrack(0, target)
	return res
}

// ===================== Test cases =====================

func main() {
	// The combinations may come back in any order, so we sort them before comparing.

	// Example 1 from the problem: 2+2+5 and 9.
	check("example 1", `sortRows(combinationSum([]int{2, 5, 6, 9}, 9))`,
		sortRows(combinationSum([]int{2, 5, 6, 9}, 9)), [][]int{{2, 2, 5}, {9}})

	// Example 2 from the problem.
	check("example 2", `sortRows(combinationSum([]int{3, 4, 5}, 16))`,
		sortRows(combinationSum([]int{3, 4, 5}, 16)), [][]int{{3, 3, 3, 3, 4}, {3, 3, 5, 5}, {3, 4, 4, 5}, {4, 4, 4, 4}})

	// Example 3 from the problem: 5 cannot be made from 3s.
	check("example 3: impossible", `sortRows(combinationSum([]int{3}, 5))`,
		sortRows(combinationSum([]int{3}, 5)), [][]int{})

	// The classic example.
	check("classic", `sortRows(combinationSum([]int{2, 3, 6, 7}, 7))`,
		sortRows(combinationSum([]int{2, 3, 6, 7}, 7)), [][]int{{2, 2, 3}, {7}})

	// One number used many times.
	check("same number four times", `sortRows(combinationSum([]int{2}, 8))`,
		sortRows(combinationSum([]int{2}, 8)), [][]int{{2, 2, 2, 2}})

	// Three different ways.
	check("three ways", `sortRows(combinationSum([]int{2, 3, 5}, 8))`,
		sortRows(combinationSum([]int{2, 3, 5}, 8)), [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}})

	// The number is the target itself.
	check("number equals target", `sortRows(combinationSum([]int{5}, 5))`,
		sortRows(combinationSum([]int{5}, 5)), [][]int{{5}})

	// Every number is bigger than the target.
	check("all numbers too big", `sortRows(combinationSum([]int{4, 5}, 3))`,
		sortRows(combinationSum([]int{4, 5}, 3)), [][]int{})

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
