// Problem:    Min Cost to Connect All Points
// Difficulty: Medium
// Link:       https://neetcode.io/problems/min-cost-to-connect-points
// Source:     Data Structures & Algorithms/min-cost-to-connect-points/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given points [x, y] on a grid, connect them all (directly or through
// other points). Connecting two points costs their Manhattan distance
// |x1 - x2| + |y1 - y2|. Return the smallest total cost.
//
// Example 1:
//
//	Input:  points = [[0, 0], [2, 2], [3, 3], [2, 4], [4, 2]]
//	Output: 10
//
// How to run (from the "local" folder):
//
//	go run ./091-min-cost-to-connect-points
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (Prim's algorithm): grow one connected group, starting with point 0.
// In every step add the point that is CHEAPEST to connect to the group.
// cost[p] = the cheapest known distance from point p to any point that is
// already in the group. After adding a point, update the others.
// Time: O(n^2).
func minCostConnectPoints(points [][]int) int {
	const inf = 1000000000
	n := len(points)

	// inGroup[p] is true once point p is connected
	inGroup := make([]bool, n)
	cost := make([]int, n)
	for i := range cost {
		cost[i] = inf
	}
	cost[0] = 0 // start with point 0, it costs nothing

	total := 0
	for step := 0; step < n; step++ {
		// find the cheapest point that is not in the group yet
		best := -1
		for p := 0; p < n; p++ {
			// take p if it is the first free point we see, or cheaper than the best so far
			if !inGroup[p] && (best == -1 || cost[p] < cost[best]) {
				best = p
			}
		}

		// add it
		// pay its cost and mark it as connected
		inGroup[best] = true
		total += cost[best]

		// maybe the new point is a cheaper way to reach the others
		for p := 0; p < n; p++ {
			if !inGroup[p] {
				// Manhattan distance: difference in x plus difference in y
				distance := absDiff(points[best][0], points[p][0]) + absDiff(points[best][1], points[p][1])
				// cheaper than every connection we knew for p: remember it
				if distance < cost[p] {
					cost[p] = distance
				}
			}
		}
	}
	// all points are in the group
	return total
}

// absDiff returns the distance between two numbers (never negative).
func absDiff(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// ===================== Test cases =====================

func main() {
	// The cost between two points is |x1 - x2| + |y1 - y2|.

	// Example 1 from the problem.
	check("example 1", `minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 3}, {2, 4}, {4, 2}})`,
		minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 3}, {2, 4}, {4, 2}}), 10)

	// The classic examples.
	check("classic 1", `minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}})`,
		minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}), 20)
	check("classic 2", `minCostConnectPoints([][]int{{3, 12}, {-2, 5}, {-4, 1}})`,
		minCostConnectPoints([][]int{{3, 12}, {-2, 5}, {-4, 1}}), 18)

	// One point: nothing to connect.
	check("one point", `minCostConnectPoints([][]int{{0, 0}})`,
		minCostConnectPoints([][]int{{0, 0}}), 0)

	// Two points.
	check("two points", `minCostConnectPoints([][]int{{0, 0}, {1, 1}})`,
		minCostConnectPoints([][]int{{0, 0}, {1, 1}}), 2)

	// Three points on a line: connect neighbors, not the two ends.
	check("points on a line", `minCostConnectPoints([][]int{{0, 0}, {1, 0}, {2, 0}})`,
		minCostConnectPoints([][]int{{0, 0}, {1, 0}, {2, 0}}), 2)

	// On a line, but not given in order.
	check("line, not in order", `minCostConnectPoints([][]int{{0, 0}, {0, 10}, {0, 3}})`,
		minCostConnectPoints([][]int{{0, 0}, {0, 10}, {0, 3}}), 10)

	// A square: 3 sides are enough.
	check("square", `minCostConnectPoints([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}})`,
		minCostConnectPoints([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}), 3)

	// Negative coordinates.
	check("negative coordinates", `minCostConnectPoints([][]int{{-1, -1}, {1, 1}})`,
		minCostConnectPoints([][]int{{-1, -1}, {1, 1}}), 4)

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
