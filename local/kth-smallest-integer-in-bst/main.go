// Problem: Kth Smallest Element in a BST
// Link:    https://neetcode.io/problems/kth-smallest-integer-in-bst
// Source:  Data Structures & Algorithms/kth-smallest-integer-in-bst/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./kth-smallest-integer-in-bst
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

// Idea: an inorder traversal (left, node, right) of a BST visits the
// values in sorted order. So collect them and take the k-th one.
// Time: O(n).
func kthSmallest(root *TreeNode, k int) int {
	values := inorderValues(root)
	return values[k-1] // k starts from 1, the slice starts from 0
}

// inorderValues returns all values of the tree in sorted order.
func inorderValues(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	values := inorderValues(node.Left)                    // smaller values first
	values = append(values, node.Val)                     // then this node
	values = append(values, inorderValues(node.Right)...) // then bigger values
	return values
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: sorted values are 1, 2, 3.
	check("example 1: smallest", kthSmallest(makeTree(2, 1, 3), 1), 1)

	// Example 2 from the problem: sorted values are 2, 3, 4, 5.
	check("example 2: biggest", kthSmallest(makeTree(4, 3, 5, 2), 4), 5)

	// One node.
	check("one node", kthSmallest(makeTree(7), 1), 7)

	// The middle value.
	check("middle value", kthSmallest(makeTree(2, 1, 3), 2), 2)

	// The classic example: sorted values are 1, 2, 3, 4.
	check("classic", kthSmallest(makeTree(3, 1, 4, nil, 2), 1), 1)

	// A bigger tree: sorted values are 1, 2, 3, 4, 5, 6.
	check("bigger tree, k = 3", kthSmallest(makeTree(5, 3, 6, 2, 4, nil, nil, 1), 3), 3)
	check("bigger tree, k = 6", kthSmallest(makeTree(5, 3, 6, 2, 4, nil, nil, 1), 6), 6)

	// A chain going left: 3 -> 2 -> 1.
	check("left chain", kthSmallest(makeTree(3, 2, nil, 1), 2), 2)

	// A chain going right: 1 -> 2 -> 3.
	check("right chain", kthSmallest(makeTree(1, nil, 2, nil, 3), 3), 3)

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
