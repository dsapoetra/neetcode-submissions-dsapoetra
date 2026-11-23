// Problem: Non-overlapping Intervals
// Link:    https://neetcode.io/problems/non-overlapping-intervals
// Source:  Data Structures & Algorithms/non-overlapping-intervals/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./non-overlapping-intervals
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

// Idea (greedy): sort by END time. Always keep the interval that ends
// earliest, because it leaves the most space for the others.
// Any interval that starts before the last kept one ends must be removed.
// Time: O(n log n).
func eraseOverlapIntervals(intervals [][]int) int {
	// sort by end time
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][1] < intervals[b][1]
	})

	removed := 0
	lastEnd := intervals[0][1] // end of the last interval we kept

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < lastEnd {
			// overlaps the one we kept -> remove this one
			removed++
		} else {
			// no overlap -> keep it
			lastEnd = intervals[i][1]
		}
	}
	return removed
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: remove [1,4].
	check("example 1", eraseOverlapIntervals([][]int{{1, 2}, {2, 4}, {1, 4}}), 1)

	// Example 2 from the problem: touching is not overlapping.
	check("example 2: touching is fine", eraseOverlapIntervals([][]int{{1, 2}, {2, 4}}), 0)

	// One interval.
	check("one interval", eraseOverlapIntervals([][]int{{1, 2}}), 0)

	// The same interval three times: keep one.
	check("three copies", eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}), 2)

	// The classic example: remove [1,3].
	check("classic", eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}), 1)

	// Tricky: keep [1,11] and [11,22], remove the other two.
	check("one long interval", eraseOverlapIntervals([][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}), 2)

	// Removing the one long interval is better than removing three short ones.
	check("remove the long one", eraseOverlapIntervals([][]int{{1, 10}, {2, 3}, {4, 5}, {6, 7}}), 1)

	// A staircase: keep every second interval.
	check("staircase", eraseOverlapIntervals([][]int{{0, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}}), 2)

	// Negative numbers.
	check("negative numbers", eraseOverlapIntervals([][]int{{-5, 0}, {-1, 1}}), 1)

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
