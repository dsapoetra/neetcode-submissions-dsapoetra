// Problem: Maximum Depth of Binary Tree
// Link:    https://neetcode.io/problems/depth-of-binary-tree
// Source:  Data Structures & Algorithms/depth-of-binary-tree/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./depth-of-binary-tree
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

// Idea: depth of a tree = 1 (this node) + the deeper of its two subtrees.
// Time: O(n).
func maxDepth(root *TreeNode) int {
	// an empty tree has depth 0
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	// take the deeper side and add 1 for the current node
	if left > right {
		return left + 1
	}
	return right + 1
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 1 -> 3 -> 4.
	check("example 1", maxDepth(makeTree(1, 2, 3, nil, nil, 4)), 3)

	// Example 2 from the problem: an empty tree.
	check("example 2: empty tree", maxDepth(makeTree()), 0)

	// One node.
	check("one node", maxDepth(makeTree(1)), 1)

	// A chain going left.
	check("left chain", maxDepth(makeTree(1, 2, nil, 3)), 3)

	// A chain going right.
	check("right chain", maxDepth(makeTree(1, nil, 2, nil, 3, nil, 4)), 4)

	// The classic example.
	check("classic", maxDepth(makeTree(3, 9, 20, nil, nil, 15, 7)), 3)

	// A full tree with 3 levels.
	check("full tree", maxDepth(makeTree(1, 2, 3, 4, 5, 6, 7)), 3)

	// The left side is much deeper than the right side.
	check("deep left side", maxDepth(makeTree(1, 2, 3, 4, nil, nil, nil, 5)), 4)

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
