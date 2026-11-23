// Problem: Merge Intervals
// Link:    https://neetcode.io/problems/merge-intervals
// Source:  Data Structures & Algorithms/merge-intervals/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./merge-intervals
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

// Idea: sort the intervals by start time. Then walk through them:
// if the current interval overlaps the last one we kept, merge them,
// otherwise keep it as a new interval.
// Time: O(n log n) because of the sort.
func merge(intervals [][]int) [][]int {
	// sort by start time
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	// start the answer with the first interval
	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		last := res[len(res)-1] // the last interval in the answer

		if cur[0] <= last[1] {
			// cur starts before last ends -> overlap.
			// extend the end of last if cur ends later.
			if cur[1] > last[1] {
				last[1] = cur[1]
			}
		} else {
			// no overlap -> cur is a new separate interval
			res = append(res, cur)
		}
	}
	return res
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", `merge([][]int{{1, 3}, {1, 5}, {6, 7}})`,
		merge([][]int{{1, 3}, {1, 5}, {6, 7}}), [][]int{{1, 5}, {6, 7}})

	// Example 2 from the problem: touching intervals are merged.
	check("example 2: touching", `merge([][]int{{1, 2}, {2, 3}})`,
		merge([][]int{{1, 2}, {2, 3}}), [][]int{{1, 3}})

	// One interval: nothing to merge.
	check("one interval", `merge([][]int{{1, 5}})`,
		merge([][]int{{1, 5}}), [][]int{{1, 5}})

	// The classic example.
	check("classic", `merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})`,
		merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}), [][]int{{1, 6}, {8, 10}, {15, 18}})

	// The input is not sorted.
	check("unsorted input", `merge([][]int{{8, 10}, {1, 3}, {2, 6}})`,
		merge([][]int{{8, 10}, {1, 3}, {2, 6}}), [][]int{{1, 6}, {8, 10}})

	// Small intervals inside a big one.
	check("inside a big one", `merge([][]int{{1, 10}, {2, 3}, {4, 5}})`,
		merge([][]int{{1, 10}, {2, 3}, {4, 5}}), [][]int{{1, 10}})

	// Nothing overlaps.
	check("no overlap", `merge([][]int{{1, 2}, {3, 4}, {5, 6}})`,
		merge([][]int{{1, 2}, {3, 4}, {5, 6}}), [][]int{{1, 2}, {3, 4}, {5, 6}})

	// The same interval twice.
	check("same interval twice", `merge([][]int{{1, 4}, {1, 4}})`,
		merge([][]int{{1, 4}, {1, 4}}), [][]int{{1, 4}})

	// A chain: each one touches the next, so all become one.
	check("chain", `merge([][]int{{1, 4}, {4, 5}, {5, 9}})`,
		merge([][]int{{1, 4}, {4, 5}, {5, 9}}), [][]int{{1, 9}})

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
