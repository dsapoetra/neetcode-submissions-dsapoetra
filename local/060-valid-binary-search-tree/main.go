// Problem:    Validate Binary Search Tree
// Difficulty: Medium
// Link:       https://neetcode.io/problems/valid-binary-search-tree
// Source:     Data Structures & Algorithms/valid-binary-search-tree/submission-1.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given the root of a binary tree, return true if it is a valid binary
// search tree: for EVERY node, all values in its left subtree are smaller
// and all values in its right subtree are bigger.
//
// Example 1:
//
//	Input:  root = [2, 1, 3]
//	Output: true
//
// Example 2:
//
//	Input:  root = [1, 2, 3]
//	Output: false         (2 is on the left of 1 but bigger)
//
// How to run (from the "local" folder):
//
//	go run ./060-valid-binary-search-tree
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
	check("example 1: valid", `isValidBST(makeTree(2, 1, 3))`,
		isValidBST(makeTree(2, 1, 3)), true)

	// Example 2 from the problem: the left child 2 is bigger than the root 1.
	check("example 2: left child too big", `isValidBST(makeTree(1, 2, 3))`,
		isValidBST(makeTree(1, 2, 3)), false)

	// One node is always valid.
	check("one node", `isValidBST(makeTree(1))`,
		isValidBST(makeTree(1)), true)

	// An empty tree is valid.
	check("empty tree", `isValidBST(makeTree())`,
		isValidBST(makeTree()), true)

	// The classic example: right child 4 is smaller than the root 5.
	check("right child too small", `isValidBST(makeTree(5, 1, 4, nil, nil, 3, 6))`,
		isValidBST(makeTree(5, 1, 4, nil, nil, 3, 6)), false)

	// Tricky: every child looks fine next to its parent (3 < 6), but 3 sits
	// in the RIGHT subtree of 5, so it must be bigger than 5.
	check("grandchild breaks the rule", `isValidBST(makeTree(5, 4, 6, nil, nil, 3, 7))`,
		isValidBST(makeTree(5, 4, 6, nil, nil, 3, 7)), false)

	// Equal values are not allowed.
	check("equal values", `isValidBST(makeTree(2, 2, 2))`,
		isValidBST(makeTree(2, 2, 2)), false)
	check("left child equal to root", `isValidBST(makeTree(1, 1))`,
		isValidBST(makeTree(1, 1)), false)
	check("right child equal to root", `isValidBST(makeTree(1, nil, 1))`,
		isValidBST(makeTree(1, nil, 1)), false)

	// A bigger valid tree.
	check("bigger valid tree", `isValidBST(makeTree(10, 5, 15, 3, 7, 12, 20))`,
		isValidBST(makeTree(10, 5, 15, 3, 7, 12, 20)), true)

	// A chain going left: 3 > 2 > 1.
	check("left chain", `isValidBST(makeTree(3, 2, nil, 1))`,
		isValidBST(makeTree(3, 2, nil, 1)), true)

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
