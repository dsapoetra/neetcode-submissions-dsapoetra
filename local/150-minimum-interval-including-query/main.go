// Problem:    Minimum Interval to Include Each Query
// Difficulty: Hard
// Link:       https://neetcode.io/problems/minimum-interval-including-query
// Source:     Data Structures & Algorithms/minimum-interval-including-query/submission-1.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a list of intervals [left, right] and a list of queries (numbers),
// return for every query the length of the SHORTEST interval that contains
// it (length = right - left + 1), or -1 if no interval contains it.
//
// Example 1:
//
//	Input:  intervals = [[1, 3], [2, 3], [3, 7], [6, 6]], queries = [2, 3, 1, 7, 6, 8]
//	Output: [2, 2, 3, 5, 1, -1]
//
// How to run (from the "local" folder):
//
//	go run ./150-minimum-interval-including-query
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

// Idea: the SHORTEST interval should answer as many queries as it can,
// then the second shortest answers the queries that are still open, ...
//  1. sort the intervals by length, shortest first
//  2. sort the queries by value (we remember where each one came from)
//  3. for each interval, find the first sorted query that is >= left
//     (binary search) and walk right while the query is <= right.
//     Every open query we meet gets this interval's length.
//
// To not walk over answered queries again and again, skip[i] points to
// the next query at or after i that might still be open. That way every
// query is answered once and then jumped over.
// Time: O(n log n + q log q).
func minInterval(intervals [][]int, queries []int) []int {
	// 1) shortest interval first
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][1]-intervals[a][0] < intervals[b][1]-intervals[b][0]
	})

	// 2) order[i] = index (in queries) of the i-th smallest query
	q := len(queries)
	order := make([]int, q)
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(a, b int) bool {
		return queries[order[a]] < queries[order[b]]
	})

	res := make([]int, q)
	for i := range res {
		res[i] = -1 // -1 = no interval contains this query
	}

	// skip[i] = i means sorted query i is still open.
	// Otherwise it points to a later position to continue from.
	skip := make([]int, q+1)
	for i := range skip {
		skip[i] = i
	}
	// nextOpen returns the first open sorted query at or after i (or q if none).
	var nextOpen func(i int) int
	nextOpen = func(i int) int {
		if skip[i] != i {
			skip[i] = nextOpen(skip[i]) // shortcut for the next time
		}
		return skip[i]
	}

	// 3) hand out the answers, shortest interval first
	for _, interval := range intervals {
		left, right := interval[0], interval[1]

		// binary search: first sorted query with value >= left
		low, high := 0, q
		for low < high {
			mid := (low + high) / 2
			if queries[order[mid]] < left {
				low = mid + 1
			} else {
				high = mid
			}
		}

		for i := nextOpen(low); i < q && queries[order[i]] <= right; i = nextOpen(i) {
			res[order[i]] = right - left + 1
			skip[i] = i + 1 // answered: jump over it from now on
		}
	}
	return res
}

// ===================== Test cases =====================

func main() {
	// The length of an interval [left, right] is right - left + 1.

	// Example 1 from the problem.
	check("example 1", `minInterval([][]int{{1, 3}, {2, 3}, {3, 7}, {6, 6}}, []int{2, 3, 1, 7, 6, 8})`,
		minInterval([][]int{{1, 3}, {2, 3}, {3, 7}, {6, 6}}, []int{2, 3, 1, 7, 6, 8}), []int{2, 2, 3, 5, 1, -1})

	// The classic examples.
	check("classic 1", `minInterval([][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5})`,
		minInterval([][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5}), []int{3, 3, 1, 4})
	check("classic 2", `minInterval([][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, []int{2, 19, 5, 22})`,
		minInterval([][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, []int{2, 19, 5, 22}), []int{2, -1, 4, 6})

	// One interval of length 1.
	check("point interval, hit", `minInterval([][]int{{1, 1}}, []int{1})`,
		minInterval([][]int{{1, 1}}, []int{1}), []int{1})
	check("point interval, miss", `minInterval([][]int{{1, 1}}, []int{2})`,
		minInterval([][]int{{1, 1}}, []int{2}), []int{-1})

	// Intervals inside each other: the innermost one that fits wins.
	check("nested intervals", `minInterval([][]int{{1, 10}, {2, 9}, {3, 8}}, []int{5, 1, 10, 2})`,
		minInterval([][]int{{1, 10}, {2, 9}, {3, 8}}, []int{5, 1, 10, 2}), []int{6, 10, 10, 8})

	// The same query several times.
	check("same query three times", `minInterval([][]int{{1, 5}}, []int{3, 3, 3})`,
		minInterval([][]int{{1, 5}}, []int{3, 3, 3}), []int{5, 5, 5})

	// The two ends of an interval count as inside.
	check("the ends count", `minInterval([][]int{{4, 6}}, []int{4, 6, 3, 7})`,
		minInterval([][]int{{4, 6}}, []int{4, 6, 3, 7}), []int{3, 3, -1, -1})

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
