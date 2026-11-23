// Problem: Validate Binary Search Tree
// Link:    https://neetcode.io/problems/valid-binary-search-tree
// Source:  Data Structures & Algorithms/valid-binary-search-tree/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./valid-binary-search-tree
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

// Idea: in a BST every node must be bigger than everything on its left
// and smaller than everything on its right. So each node has an allowed
// range (low, high). Going left lowers the upper limit, going right
// raises the lower limit.
// Time: O(n).
func isValidBST(root *TreeNode) bool {
	// the root has no limits (nil means no limit)
	return checkBST(root, nil, nil)
}

// checkBST returns true if every node in this subtree has a value
// strictly between low.Val and high.Val.
func checkBST(node *TreeNode, low *TreeNode, high *TreeNode) bool {
	if node == nil {
		return true
	}
	// must be bigger than the lower limit
	if low != nil && node.Val <= low.Val {
		return false
	}
	// must be smaller than the upper limit
	if high != nil && node.Val >= high.Val {
		return false
	}
	// left subtree: values must be smaller than this node
	// right subtree: values must be bigger than this node
	return checkBST(node.Left, low, node) && checkBST(node.Right, node, high)
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 1 < 2 < 3.
	check("example 1: valid", isValidBST(makeTree(2, 1, 3)), true)

	// Example 2 from the problem: the left child 2 is bigger than the root 1.
	check("example 2: left child too big", isValidBST(makeTree(1, 2, 3)), false)

	// One node is always valid.
	check("one node", isValidBST(makeTree(1)), true)

	// An empty tree is valid.
	check("empty tree", isValidBST(makeTree()), true)

	// The classic example: right child 4 is smaller than the root 5.
	check("right child too small", isValidBST(makeTree(5, 1, 4, nil, nil, 3, 6)), false)

	// Tricky: every child looks fine next to its parent (3 < 6), but 3 sits
	// in the RIGHT subtree of 5, so it must be bigger than 5.
	check("grandchild breaks the rule", isValidBST(makeTree(5, 4, 6, nil, nil, 3, 7)), false)

	// Equal values are not allowed.
	check("equal values", isValidBST(makeTree(2, 2, 2)), false)
	check("left child equal to root", isValidBST(makeTree(1, 1)), false)
	check("right child equal to root", isValidBST(makeTree(1, nil, 1)), false)

	// A bigger valid tree.
	check("bigger valid tree", isValidBST(makeTree(10, 5, 15, 3, 7, 12, 20)), true)

	// A chain going left: 3 > 2 > 1.
	check("left chain", isValidBST(makeTree(3, 2, nil, 1)), true)

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
