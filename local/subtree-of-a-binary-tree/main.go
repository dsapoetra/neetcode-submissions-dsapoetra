// Problem: Subtree of Another Tree
// Link:    https://neetcode.io/problems/subtree-of-a-binary-tree
// Source:  Data Structures & Algorithms/subtree-of-a-binary-tree/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./subtree-of-a-binary-tree
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

// Idea: for every node in the big tree, check if the tree starting at
// that node is exactly the same as subRoot.
// Time: O(nodes in root * nodes in subRoot).
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// an empty tree is a subtree of anything
	if subRoot == nil {
		return true
	}
	// nothing left to search in
	if root == nil {
		return false
	}
	// does the tree starting here match?
	if sameTree(root, subRoot) {
		return true
	}
	// otherwise try the left and right children
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// sameTree returns true if the two trees are exactly the same.
func sameTree(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return sameTree(a.Left, b.Left) && sameTree(a.Right, b.Right)
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: the left part of the big tree is exactly the small tree.
	check("example 1: is a subtree", isSubtree(makeTree(1, 2, 3, 4, 5), makeTree(2, 4, 5)), true)

	// Example 2 from the problem: the big tree has an extra node (6) under 4,
	// so that part is NOT exactly the same.
	check("example 2: extra node below", isSubtree(makeTree(1, 2, 3, 4, 5, nil, nil, 6), makeTree(2, 4, 5)), false)

	// A tree is a subtree of itself.
	check("same tree", isSubtree(makeTree(1, 2, 3), makeTree(1, 2, 3)), true)

	// Single nodes.
	check("single nodes, same", isSubtree(makeTree(1), makeTree(1)), true)
	check("single nodes, different", isSubtree(makeTree(1), makeTree(2)), false)

	// An empty tree is a subtree of every tree.
	check("empty subRoot", isSubtree(makeTree(1, 2, 3), makeTree()), true)

	// Nothing (except an empty tree) is a subtree of an empty tree.
	check("empty root", isSubtree(makeTree(), makeTree(1)), false)

	// The classic example.
	check("classic true", isSubtree(makeTree(3, 4, 5, 1, 2), makeTree(4, 1, 2)), true)

	// The classic "false" example: an extra 0 under the 2.
	check("classic false", isSubtree(makeTree(3, 4, 5, 1, 2, nil, nil, nil, nil, 0), makeTree(4, 1, 2)), false)

	// A single leaf is a subtree.
	check("a leaf", isSubtree(makeTree(1, 2, 3), makeTree(3)), true)

	// Same values everywhere: the first try (at the root) fails, the child matches.
	check("repeated values", isSubtree(makeTree(1, 1), makeTree(1)), true)

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
