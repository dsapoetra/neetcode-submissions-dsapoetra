// Problem:    Count Good Nodes in Binary Tree
// Difficulty: Medium
// Link:       https://neetcode.io/problems/count-good-nodes-in-binary-tree
// Source:     Data Structures & Algorithms/count-good-nodes-in-binary-tree/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// A node is "good" if no node on the way from the root down to it has a
// bigger value. Return the number of good nodes in the tree.
//
// Example 1:
//
//	Input:  root = [2, 1, 1, 3, null, 1, 5]
//	Output: 3
//
// Example 2:
//
//	Input:  root = [1, 2, -1, 3, 4]
//	Output: 4
//
// How to run (from the "local" folder):
//
//	go run ./059-count-good-nodes-in-binary-tree
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// TreeNode is one node of a binary tree.
// NeetCode gives us this type; locally we have to define it ourselves.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ===================== Solution =====================

// Idea: a node is "good" if no node on the way from the root down to it
// has a bigger value. So walk down the tree and carry along the biggest
// value seen so far on this path. A node is good if its value is at
// least that big.
// Time: O(n).
func goodNodes(root *TreeNode) int {
	return countGood(root, root.Val)
}

// countGood counts the good nodes in this subtree.
// maxSoFar is the biggest value on the path from the root to this node.
func countGood(node *TreeNode, maxSoFar int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val >= maxSoFar {
		count = 1           // this node is good
		maxSoFar = node.Val // and it is the new biggest value for the nodes below
	}

	count += countGood(node.Left, maxSoFar)
	count += countGood(node.Right, maxSoFar)
	return count
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: the root 2, the 3 and the 5.
	check("example 1", `goodNodes(makeTree(2, 1, 1, 3, nil, 1, 5))`,
		goodNodes(makeTree(2, 1, 1, 3, nil, 1, 5)), 3)

	// Example 2 from the problem: everything except the -1.
	check("example 2", `goodNodes(makeTree(1, 2, -1, 3, 4))`,
		goodNodes(makeTree(1, 2, -1, 3, 4)), 4)

	// The classic examples.
	check("classic 1", `goodNodes(makeTree(3, 1, 4, 3, nil, 1, 5))`,
		goodNodes(makeTree(3, 1, 4, 3, nil, 1, 5)), 4)
	check("classic 2", `goodNodes(makeTree(3, 3, nil, 4, 2))`,
		goodNodes(makeTree(3, 3, nil, 4, 2)), 3)

	// One node: the root is always good.
	check("one node", `goodNodes(makeTree(1))`,
		goodNodes(makeTree(1)), 1)

	// The root is the biggest: only the root is good.
	check("root is the biggest", `goodNodes(makeTree(5, 4, 3))`,
		goodNodes(makeTree(5, 4, 3)), 1)

	// Values grow going down: every node is good.
	check("growing values", `goodNodes(makeTree(1, 2, 3))`,
		goodNodes(makeTree(1, 2, 3)), 3)

	// Equal values count as good.
	check("equal values", `goodNodes(makeTree(2, 2, 2))`,
		goodNodes(makeTree(2, 2, 2)), 3)

	// Negative values.
	check("negative values", `goodNodes(makeTree(-1, -2, -3))`,
		goodNodes(makeTree(-1, -2, -3)), 1)

	// 6 is bigger than its parent 3, but smaller than the root 9.
	check("bigger than the parent only", `goodNodes(makeTree(9, nil, 3, 6))`,
		goodNodes(makeTree(9, nil, 3, 6)), 1)

	report()
}

// ===================== Helpers for the tests =====================

// makeTree builds a binary tree from values given level by level (the same
// format NeetCode/LeetCode use). nil means "no node here".
// Example: makeTree(1, 2, 3, nil, 4) gives
//
//	  1
//	 / \
//	2   3
//	 \
//	  4
func makeTree(vals ...interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root} // nodes that are still waiting for their children
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		// next value is the left child
		if vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// the value after that is the right child
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
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
