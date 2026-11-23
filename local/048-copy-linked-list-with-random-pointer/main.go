// Problem:    Copy List with Random Pointer
// Difficulty: Medium
// Link:       https://neetcode.io/problems/copy-linked-list-with-random-pointer
// Source:     Data Structures & Algorithms/copy-linked-list-with-random-pointer/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Every node of a linked list has a Next pointer and also a Random pointer
// that points to any node of the list (or to nothing). Make a deep copy of
// the list: all new nodes, with Next and Random pointing to the matching
// NEW nodes.
// A list is written as pairs [value, index of the random node] (null = nothing).
//
// Example 1:
//
//	Input:  head = [[3, null], [7, 3], [4, 0], [5, 1]]
//	Output: [[3, null], [7, 3], [4, 0], [5, 1]]
//
// Example 2:
//
//	Input:  head = [[1, null], [2, 2], [3, 2]]
//	Output: [[1, null], [2, 2], [3, 2]]
//
// How to run (from the "local" folder):
//
//	go run ./048-copy-linked-list-with-random-pointer
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// Node is one node of a linked list. Besides Next it has a Random pointer
// that can point to any node of the list (or to nothing).
// NeetCode gives us this type; locally we have to define it ourselves.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// ===================== Solution =====================

// Idea: copy in two passes, with a map from each old node to its copy.
//
//	pass 1: create a copy of every node (only the value)
//	pass 2: set Next and Random of every copy, using the map to find
//	        the copy that belongs to the old node's Next / Random
//
// Time: O(n), extra space: O(n).
func copyRandomList(head *Node) *Node {
	copies := map[*Node]*Node{} // old node -> its copy

	// pass 1: create all the copies
	for cur := head; cur != nil; cur = cur.Next {
		copies[cur] = &Node{Val: cur.Val}
	}

	// pass 2: connect the copies.
	// copies[nil] gives nil, which is exactly what we want for "no node".
	for cur := head; cur != nil; cur = cur.Next {
		copies[cur].Next = copies[cur.Next]
		copies[cur].Random = copies[cur.Random]
	}

	return copies[head]
}

// ===================== Test cases =====================

func main() {
	// A list is written as pairs [value, random]: "random" is the INDEX of
	// the node the random pointer points to, or -1 for "nothing".

	// build creates the list from the pairs.
	build := func(pairs [][]int) *Node {
		nodes := []*Node{}
		for _, pair := range pairs {
			nodes = append(nodes, &Node{Val: pair[0]})
		}
		for i, pair := range pairs {
			if i+1 < len(nodes) {
				nodes[i].Next = nodes[i+1]
			}
			if pair[1] >= 0 {
				nodes[i].Random = nodes[pair[1]]
			}
		}
		if len(nodes) == 0 {
			return nil
		}
		return nodes[0]
	}

	// describe is the opposite of build: it turns a list back into pairs.
	describe := func(head *Node) [][]int {
		index := map[*Node]int{}
		i := 0
		for cur := head; cur != nil; cur = cur.Next {
			index[cur] = i
			i++
		}
		pairs := [][]int{}
		for cur := head; cur != nil; cur = cur.Next {
			random := -1
			if cur.Random != nil {
				random = index[cur.Random]
			}
			pairs = append(pairs, []int{cur.Val, random})
		}
		return pairs
	}

	// copyOf builds the list, copies it, and describes the COPY.
	// It must look exactly like the pairs we started with.
	copyOf := func(pairs [][]int) [][]int {
		return describe(copyRandomList(build(pairs)))
	}

	// copyIsNew is true if the copy uses only NEW nodes: no node of the
	// copy (and no random pointer of the copy) may be a node of the original.
	copyIsNew := func(pairs [][]int) bool {
		original := build(pairs)
		old := map[*Node]bool{}
		for cur := original; cur != nil; cur = cur.Next {
			old[cur] = true
		}
		for cur := copyRandomList(original); cur != nil; cur = cur.Next {
			if old[cur] || old[cur.Random] {
				return false
			}
		}
		return true
	}

	// Example 1 from the problem.
	check("example 1: same shape", `copyOf([][]int{{3, -1}, {7, 3}, {4, 0}, {5, 1}})`,
		copyOf([][]int{{3, -1}, {7, 3}, {4, 0}, {5, 1}}), [][]int{{3, -1}, {7, 3}, {4, 0}, {5, 1}})
	check("example 1: new nodes", `copyIsNew([][]int{{3, -1}, {7, 3}, {4, 0}, {5, 1}})`,
		copyIsNew([][]int{{3, -1}, {7, 3}, {4, 0}, {5, 1}}), true)

	// Example 2 from the problem: two nodes point at the same node, one points at itself.
	check("example 2: same shape", `copyOf([][]int{{1, -1}, {2, 2}, {3, 2}})`,
		copyOf([][]int{{1, -1}, {2, 2}, {3, 2}}), [][]int{{1, -1}, {2, 2}, {3, 2}})
	check("example 2: new nodes", `copyIsNew([][]int{{1, -1}, {2, 2}, {3, 2}})`,
		copyIsNew([][]int{{1, -1}, {2, 2}, {3, 2}}), true)

	// An empty list.
	check("empty list", `copyOf([][]int{})`,
		copyOf([][]int{}), [][]int{})

	// One node whose random pointer points at itself.
	check("single node pointing at itself", `copyOf([][]int{{1, 0}})`,
		copyOf([][]int{{1, 0}}), [][]int{{1, 0}})
	check("single node: new node", `copyIsNew([][]int{{1, 0}})`,
		copyIsNew([][]int{{1, 0}}), true)

	// No random pointers at all.
	check("no random pointers", `copyOf([][]int{{1, -1}, {2, -1}, {3, -1}})`,
		copyOf([][]int{{1, -1}, {2, -1}, {3, -1}}), [][]int{{1, -1}, {2, -1}, {3, -1}})

	// Every random pointer points at the head.
	check("all point at the head", `copyOf([][]int{{1, 0}, {2, 0}, {3, 0}})`,
		copyOf([][]int{{1, 0}, {2, 0}, {3, 0}}), [][]int{{1, 0}, {2, 0}, {3, 0}})

	// Equal values: the nodes must still be told apart correctly.
	check("equal values", `copyOf([][]int{{7, 1}, {7, 0}})`,
		copyOf([][]int{{7, 1}, {7, 0}}), [][]int{{7, 1}, {7, 0}})

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
