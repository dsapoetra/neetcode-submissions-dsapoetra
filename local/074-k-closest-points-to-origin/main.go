// Problem:    K Closest Points to Origin
// Difficulty: Medium
// Link:       https://neetcode.io/problems/k-closest-points-to-origin
// Source:     Data Structures & Algorithms/k-closest-points-to-origin/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a list of points [x, y] and a number k, return the k points that
// are closest to the origin (0, 0). The answer may be in any order.
//
// Example 1:
//
//	Input:  points = [[0, 2], [2, 2]], k = 1
//	Output: [[0, 2]]
//
// Example 2:
//
//	Input:  points = [[0, 2], [2, 0], [2, 2]], k = 2
//	Output: [[0, 2], [2, 0]]
//
// How to run (from the "local" folder):
//
//	go run ./074-k-closest-points-to-origin
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

// Idea: sort the points by their distance to (0, 0) and take the first k.
// We compare x*x + y*y instead of the real distance: the square root does
// not change which one is smaller, so we can skip it.
// Time: O(n log n).
func kClosest(points [][]int, k int) [][]int {
	// sort.Slice asks us: "must point a come before point b?"
	sort.Slice(points, func(a, b int) bool {
		// squared distance to (0, 0): x*x + y*y
		distA := points[a][0]*points[a][0] + points[a][1]*points[a][1]
		distB := points[b][0]*points[b][0] + points[b][1]*points[b][1]
		// the closer point comes first
		return distA < distB
	})
	// the list is now ordered from close to far: the first k are the answer
	return points[:k]
}

// ===================== Test cases =====================

func main() {
	// The points may come back in any order, so we sort the list of points
	// (not the numbers inside a point) before comparing.

	// Example 1 from the problem: (0,2) is closer than (2,2).
	check("example 1", `sortLists(kClosest([][]int{{0, 2}, {2, 2}}, 1))`,
		sortLists(kClosest([][]int{{0, 2}, {2, 2}}, 1)), [][]int{{0, 2}})

	// Example 2 from the problem.
	check("example 2", `sortLists(kClosest([][]int{{0, 2}, {2, 0}, {2, 2}}, 2))`,
		sortLists(kClosest([][]int{{0, 2}, {2, 0}, {2, 2}}, 2)), [][]int{{0, 2}, {2, 0}})

	// The classic examples.
	check("classic 1", `sortLists(kClosest([][]int{{1, 3}, {-2, 2}}, 1))`,
		sortLists(kClosest([][]int{{1, 3}, {-2, 2}}, 1)), [][]int{{-2, 2}})
	check("classic 2", `sortLists(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))`,
		sortLists(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2)), [][]int{{-2, 4}, {3, 3}})

	// k = number of points: all of them.
	check("all points", `sortLists(kClosest([][]int{{2, 2}, {1, 1}}, 2))`,
		sortLists(kClosest([][]int{{2, 2}, {1, 1}}, 2)), [][]int{{1, 1}, {2, 2}})

	// The origin itself is the closest point possible.
	check("the origin", `sortLists(kClosest([][]int{{1, 1}, {0, 0}}, 1))`,
		sortLists(kClosest([][]int{{1, 1}, {0, 0}}, 1)), [][]int{{0, 0}})

	// Negative coordinates.
	check("negative coordinates", `sortLists(kClosest([][]int{{5, 5}, {-1, -1}}, 1))`,
		sortLists(kClosest([][]int{{5, 5}, {-1, -1}}, 1)), [][]int{{-1, -1}})

	// One point.
	check("one point", `sortLists(kClosest([][]int{{7, -7}}, 1))`,
		sortLists(kClosest([][]int{{7, -7}}, 1)), [][]int{{7, -7}})

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
