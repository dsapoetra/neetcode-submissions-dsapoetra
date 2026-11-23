// Problem: Insert Interval
// Link:    https://neetcode.io/problems/insert-new-interval
// Source:  Data Structures & Algorithms/insert-new-interval/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./insert-new-interval
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: the intervals are already sorted, so do it in 3 steps:
// 1) copy every interval that ends before the new one starts
// 2) merge every interval that overlaps the new one into it
// 3) copy every interval that is left
// Time: O(n).
func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	i := 0

	// step 1: intervals completely before the new one
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	// step 2: intervals that overlap the new one.
	// grow newInterval so it covers them.
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	res = append(res, newInterval)

	// step 3: intervals completely after the new one
	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}
	return res
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: the new interval joins both.
	check("example 1", insert([][]int{{1, 3}, {4, 6}}, []int{2, 5}), [][]int{{1, 6}})

	// Example 2 from the problem: the new interval fits in a gap.
	check("example 2: fits in a gap", insert([][]int{{1, 2}, {3, 5}, {9, 10}}, []int{6, 7}), [][]int{{1, 2}, {3, 5}, {6, 7}, {9, 10}})

	// No intervals yet.
	check("empty list", insert([][]int{}, []int{5, 7}), [][]int{{5, 7}})

	// The new interval is inside an existing one.
	check("inside an existing one", insert([][]int{{1, 5}}, []int{2, 3}), [][]int{{1, 5}})

	// The new interval comes before everything.
	check("before all", insert([][]int{{3, 5}}, []int{1, 2}), [][]int{{1, 2}, {3, 5}})

	// The new interval comes after everything.
	check("after all", insert([][]int{{1, 2}}, []int{3, 4}), [][]int{{1, 2}, {3, 4}})

	// The new interval only TOUCHES its neighbors: touching counts as overlap.
	check("touching both sides", insert([][]int{{1, 2}, {5, 6}}, []int{2, 5}), [][]int{{1, 6}})

	// The classic example: the new interval swallows three intervals.
	check("classic", insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}), [][]int{{1, 2}, {3, 10}, {12, 16}})

	// The new interval covers everything.
	check("covers everything", insert([][]int{{2, 3}, {4, 5}}, []int{1, 10}), [][]int{{1, 10}})

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
