// Problem:    Number of Connected Components in an Undirected Graph
// Difficulty: Medium
// Link:       https://neetcode.io/problems/count-connected-components
// Source:     Data Structures & Algorithms/count-connected-components/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given n nodes (0 .. n-1) and a list of undirected edges, return the
// number of connected components (groups of nodes that are linked together).
//
// Example 1:
//
//	Input:  n = 3, edges = [[0, 1], [0, 2]]
//	Output: 1
//
// Example 2:
//
//	Input:  n = 6, edges = [[0, 1], [1, 2], [2, 3], [4, 5]]
//	Output: 2
//
// How to run (from the "local" folder):
//
//	go run ./088-count-connected-components
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: go through every node. If we have not seen it yet, it belongs
// to a new component: count it, then use DFS to mark every node in
// that component as seen.
// Time: O(nodes + edges).
func countComponents(n int, edges [][]int) int {
	// build the neighbor list of every node
	neighbors := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		neighbors[a] = append(neighbors[a], b)
		neighbors[b] = append(neighbors[b], a)
	}

	seen := make([]bool, n)
	count := 0
	for node := 0; node < n; node++ {
		if !seen[node] {
			count++ // found a new component
			visit(node, neighbors, seen)
		}
	}
	return count
}

// visit marks node and everything connected to it as seen (DFS).
func visit(node int, neighbors [][]int, seen []bool) {
	seen[node] = true
	for _, next := range neighbors[node] {
		if !seen[next] {
			visit(next, neighbors, seen)
		}
	}
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: everything is connected through node 0.
	check("example 1", `countComponents(3, [][]int{{0, 1}, {0, 2}})`,
		countComponents(3, [][]int{{0, 1}, {0, 2}}), 1)

	// Example 2 from the problem: {0,1,2,3} and {4,5}.
	check("example 2", `countComponents(6, [][]int{{0, 1}, {1, 2}, {2, 3}, {4, 5}})`,
		countComponents(6, [][]int{{0, 1}, {1, 2}, {2, 3}, {4, 5}}), 2)

	// One node.
	check("one node", `countComponents(1, [][]int{})`,
		countComponents(1, [][]int{}), 1)

	// No edges: every node is its own component.
	check("no edges", `countComponents(4, [][]int{})`,
		countComponents(4, [][]int{}), 4)

	// The classic example: {0,1,2} and {3,4}.
	check("classic", `countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}})`,
		countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}}), 2)

	// A cycle is still one component.
	check("cycle", `countComponents(3, [][]int{{0, 1}, {1, 2}, {2, 0}})`,
		countComponents(3, [][]int{{0, 1}, {1, 2}, {2, 0}}), 1)

	// Two pairs and one node alone.
	check("two pairs and a single", `countComponents(5, [][]int{{0, 1}, {2, 3}})`,
		countComponents(5, [][]int{{0, 1}, {2, 3}}), 3)

	// Two nodes, one edge.
	check("two nodes", `countComponents(2, [][]int{{0, 1}})`,
		countComponents(2, [][]int{{0, 1}}), 1)

	// The edges are given "backwards" (bigger node first).
	check("edges backwards", `countComponents(4, [][]int{{3, 2}, {1, 0}})`,
		countComponents(4, [][]int{{3, 2}, {1, 0}}), 2)

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
