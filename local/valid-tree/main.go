// Problem: Graph Valid Tree
// Link:    https://neetcode.io/problems/valid-tree
// Source:  Data Structures & Algorithms/valid-tree/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./valid-tree
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: a graph with n nodes is a tree when:
//  1. it has exactly n-1 edges, and
//  2. every node can be reached from node 0 (it is connected).
//
// With exactly n-1 edges, a connected graph cannot have a cycle.
// Time: O(nodes + edges).
func validTree(n int, edges [][]int) bool {
	// rule 1
	if len(edges) != n-1 {
		return false
	}

	// build the neighbor list of every node
	neighbors := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		neighbors[a] = append(neighbors[a], b)
		neighbors[b] = append(neighbors[b], a)
	}

	// rule 2: visit everything reachable from node 0
	seen := make([]bool, n)
	visit(0, neighbors, seen)

	for i := 0; i < n; i++ {
		if !seen[i] {
			return false // a node was not reached
		}
	}
	return true
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
	// Example 1 from the problem: connected and no cycle.
	check("example 1: a tree", validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}), true)

	// Example 2 from the problem: 1 - 2 - 3 - 1 is a cycle.
	check("example 2: has a cycle", validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}), false)

	// One node and no edges is a tree.
	check("one node", validTree(1, [][]int{}), true)

	// Two nodes without an edge are not connected.
	check("two nodes, not connected", validTree(2, [][]int{}), false)

	// Two nodes with one edge.
	check("two nodes, connected", validTree(2, [][]int{{0, 1}}), true)

	// A triangle is a cycle.
	check("triangle", validTree(3, [][]int{{0, 1}, {1, 2}, {2, 0}}), false)

	// Two separate pairs: no cycle, but not connected.
	check("two separate pairs", validTree(4, [][]int{{0, 1}, {2, 3}}), false)

	// Tricky: the number of edges is right (n-1 = 3), but there is a cycle
	// AND node 3 is alone. Counting edges is not enough.
	check("right edge count, still not a tree", validTree(4, [][]int{{0, 1}, {1, 2}, {2, 0}}), false)

	// A star: node 0 in the middle.
	check("star", validTree(4, [][]int{{0, 1}, {0, 2}, {0, 3}}), true)

	// A chain.
	check("chain", validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}), true)

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
