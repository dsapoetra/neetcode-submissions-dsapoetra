// Problem: Clone Graph
// Link:    https://neetcode.io/problems/clone-graph
// Source:  Data Structures & Algorithms/clone-graph/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./clone-graph
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// Node is one node of a graph.
// NeetCode gives us this type; locally we have to define it ourselves.
type Node struct {
	Val       int
	Neighbors []*Node
}

// ===================== Solution =====================

// Idea: DFS through the graph and create a copy of each node.
// A map (old node -> its copy) remembers what we already copied,
// so a node is never copied twice and cycles do not loop forever.
// Time: O(nodes + edges).
func cloneGraph(node *Node) *Node {
	copies := map[*Node]*Node{}
	return cloneNode(node, copies)
}

// cloneNode returns the copy of node, creating it if needed.
func cloneNode(node *Node, copies map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	// already copied -> reuse it
	existing, found := copies[node]
	if found {
		return existing
	}

	// make the copy and save it in the map BEFORE visiting neighbors
	newNode := &Node{Val: node.Val}
	copies[node] = newNode

	// copy every neighbor and connect it
	for _, neighbor := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, cloneNode(neighbor, copies))
	}
	return newNode
}

// ===================== Test cases =====================

func main() {
	// A graph is written as a neighbor list: entry 0 is the neighbors of
	// node 1, entry 1 the neighbors of node 2, and so on.
	// For every graph we check two things with these two shortcuts:

	// cloneShape builds the graph, clones it, and returns the neighbor list
	// of the CLONE. It must be the same as the neighbor list we started with.
	cloneShape := func(adj [][]int) [][]int {
		return graphAdj(cloneGraph(buildGraph(adj)))
	}

	// cloneIsNew builds the graph, clones it, and returns true if the clone
	// is made of NEW nodes (it shares no node with the original).
	cloneIsNew := func(adj [][]int) bool {
		original := buildGraph(adj)
		return !sharesNodes(original, cloneGraph(original))
	}

	// Example 1 from the problem: 1 - 2 - 3 in a line.
	check("example 1: same shape", `cloneShape([][]int{{2}, {1, 3}, {2}})`,
		cloneShape([][]int{{2}, {1, 3}, {2}}), [][]int{{2}, {1, 3}, {2}})
	check("example 1: new nodes", `cloneIsNew([][]int{{2}, {1, 3}, {2}})`,
		cloneIsNew([][]int{{2}, {1, 3}, {2}}), true)

	// Example 2 from the problem: one node without neighbors.
	check("example 2: single node, same shape", `cloneShape([][]int{{}})`,
		cloneShape([][]int{{}}), [][]int{{}})
	check("example 2: single node, new node", `cloneIsNew([][]int{{}})`,
		cloneIsNew([][]int{{}}), true)

	// Example 3 from the problem: an empty graph stays empty.
	check("example 3: empty graph", `cloneGraph(nil) == nil`,
		cloneGraph(nil) == nil, true)

	// Two nodes pointing at each other.
	check("two nodes: same shape", `cloneShape([][]int{{2}, {1}})`,
		cloneShape([][]int{{2}, {1}}), [][]int{{2}, {1}})
	check("two nodes: new nodes", `cloneIsNew([][]int{{2}, {1}})`,
		cloneIsNew([][]int{{2}, {1}}), true)

	// A square (a cycle of 4). The DFS must not loop forever.
	check("square: same shape", `cloneShape([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})`,
		cloneShape([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}), [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
	check("square: new nodes", `cloneIsNew([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})`,
		cloneIsNew([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}), true)

	// Every node is connected to every other node.
	check("triangle: same shape", `cloneShape([][]int{{2, 3}, {1, 3}, {1, 2}})`,
		cloneShape([][]int{{2, 3}, {1, 3}, {1, 2}}), [][]int{{2, 3}, {1, 3}, {1, 2}})
	check("triangle: new nodes", `cloneIsNew([][]int{{2, 3}, {1, 3}, {1, 2}})`,
		cloneIsNew([][]int{{2, 3}, {1, 3}, {1, 2}}), true)

	// A node that is its own neighbor.
	check("self loop: same shape", `cloneShape([][]int{{1}})`,
		cloneShape([][]int{{1}}), [][]int{{1}})
	check("self loop: new node", `cloneIsNew([][]int{{1}})`,
		cloneIsNew([][]int{{1}}), true)

	report()
}

// ===================== Helpers for the tests =====================

// collectNodes returns every node we can reach from start.
func collectNodes(start *Node) []*Node {
	if start == nil {
		return nil
	}
	seen := map[*Node]bool{start: true}
	nodes := []*Node{start}
	for i := 0; i < len(nodes); i++ {
		for _, neighbor := range nodes[i].Neighbors {
			if !seen[neighbor] {
				seen[neighbor] = true
				nodes = append(nodes, neighbor)
			}
		}
	}
	return nodes
}

// buildGraph builds a graph from a neighbor list. Node values start at 1:
// adj[0] lists the neighbors of node 1, adj[1] the neighbors of node 2, ...
// It returns node 1 (or nil for an empty graph).
func buildGraph(adj [][]int) *Node {
	if len(adj) == 0 {
		return nil
	}
	nodes := make([]*Node, len(adj))
	for i := range adj {
		nodes[i] = &Node{Val: i + 1}
	}
	for i, neighbors := range adj {
		for _, val := range neighbors {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[val-1])
		}
	}
	return nodes[0]
}

// graphAdj is the opposite of buildGraph: it turns a graph back into a
// neighbor list, so it is easy to print and compare.
func graphAdj(start *Node) [][]int {
	nodes := collectNodes(start)
	adj := make([][]int, len(nodes))
	for _, node := range nodes {
		neighbors := []int{}
		for _, neighbor := range node.Neighbors {
			neighbors = append(neighbors, neighbor.Val)
		}
		adj[node.Val-1] = neighbors
	}
	return adj
}

// sharesNodes is true if graph b uses at least one node that also belongs
// to graph a. A real (deep) copy must not share any node with the original.
func sharesNodes(a *Node, b *Node) bool {
	inA := map[*Node]bool{}
	for _, node := range collectNodes(a) {
		inA[node] = true
	}
	for _, node := range collectNodes(b) {
		if inA[node] {
			return true
		}
	}
	return false
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
