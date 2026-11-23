// Problem:    Redundant Connection
// Difficulty: Medium
// Link:       https://neetcode.io/problems/redundant-connection
// Source:     Data Structures & Algorithms/redundant-connection/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// A tree with n nodes (numbered from 1) got ONE extra edge, so the graph
// now has exactly one cycle. Given the list of edges, return an edge that
// can be removed to make it a tree again. If several edges work, return the
// one that comes last in the list.
//
// Example 1:
//
//	Input:  edges = [[1, 2], [1, 3], [3, 4], [2, 4]]
//	Output: [2, 4]
//
// Example 2:
//
//	Input:  edges = [[1, 2], [1, 3], [1, 4], [3, 4], [4, 5]]
//	Output: [3, 4]
//
// How to run (from the "local" folder):
//
//	go run ./089-redundant-connection
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (union-find): start with every node in its own group. Add the
// edges one by one; an edge joins the groups of its two nodes.
// If the two nodes are ALREADY in the same group, they are connected
// without this edge, so this edge closes a cycle: it is the answer.
// parent[x] points to another node of the same group; following the
// parents leads to the "boss" of the group.
func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1) // nodes are numbered from 1
	for i := range parent {
		parent[i] = i // every node is its own boss
	}

	// find returns the boss of x's group.
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) // shortcut: point straight to the boss
		}
		return parent[x]
	}

	for _, edge := range edges {
		bossA := find(edge[0])
		bossB := find(edge[1])

		if bossA == bossB {
			return edge // already connected: this edge is the extra one
		}
		parent[bossA] = bossB // join the two groups
	}
	return []int{}
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 1-2, 1-3, 3-4 are fine; 2-4 closes a cycle.
	check("example 1", `findRedundantConnection([][]int{{1, 2}, {1, 3}, {3, 4}, {2, 4}})`,
		findRedundantConnection([][]int{{1, 2}, {1, 3}, {3, 4}, {2, 4}}), []int{2, 4})

	// Example 2 from the problem: the extra edge is not the last one.
	check("example 2", `findRedundantConnection([][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}})`,
		findRedundantConnection([][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}), []int{3, 4})

	// The classic examples.
	check("classic 1", `findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}})`,
		findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}), []int{2, 3})
	check("classic 2", `findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}})`,
		findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}), []int{1, 4})

	// A triangle.
	check("triangle", `findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 1}})`,
		findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 1}}), []int{3, 1})

	// The cycle does not include node 1.
	check("cycle without node 1", `findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 2}})`,
		findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 2}}), []int{4, 2})

	// One big cycle through all nodes.
	check("one big cycle", `findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}})`,
		findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}}), []int{4, 1})

	// Edges written with the bigger node first.
	check("bigger node first", `findRedundantConnection([][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}})`,
		findRedundantConnection([][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}}), []int{1, 4})

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
