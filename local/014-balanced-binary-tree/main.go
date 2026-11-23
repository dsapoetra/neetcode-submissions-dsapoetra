// Problem:    Balanced Binary Tree
// Difficulty: Easy
// Link:       https://neetcode.io/problems/balanced-binary-tree
// Source:     Data Structures & Algorithms/balanced-binary-tree/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// A binary tree is balanced if, for EVERY node, the heights of its left and
// right subtree differ by at most 1. Return true if the tree is balanced.
//
// Example 1:
//
//	Input:  root = [1, 2, 3, null, null, 4]
//	Output: true
//
// Example 2:
//
//	Input:  root = [1, 2, 3, null, null, 4, null, 5]
//	Output: false
//
// How to run (from the "local" folder):
//
//	go run ./014-balanced-binary-tree
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

// Idea: a tree is balanced if, at EVERY node, the heights of the left and
// right subtree differ by at most 1.
// The helper returns the height of a subtree, or -1 as a signal that
// something below is already unbalanced. That way we visit every node
// only once.
// Time: O(n).
func isBalanced(root *TreeNode) bool {
	return balancedHeight(root) != -1
}

// balancedHeight returns the height of the tree, or -1 if it is not balanced.
func balancedHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := balancedHeight(node.Left)
	right := balancedHeight(node.Right)

	// a subtree is already unbalanced: pass the bad news up
	if left == -1 || right == -1 {
		return -1
	}

	// the two sides differ by more than 1: unbalanced here
	if left-right > 1 || right-left > 1 {
		return -1
	}

	if left > right {
		return left + 1
	}
	return right + 1
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1: balanced", `isBalanced(makeTree(1, 2, 3, nil, nil, 4))`,
		isBalanced(makeTree(1, 2, 3, nil, nil, 4)), true)

	// Example 2 from the problem: the right side is 2 levels deeper than the left.
	check("example 2: not balanced", `isBalanced(makeTree(1, 2, 3, nil, nil, 4, nil, 5))`,
		isBalanced(makeTree(1, 2, 3, nil, nil, 4, nil, 5)), false)

	// An empty tree is balanced.
	check("empty tree", `isBalanced(makeTree())`,
		isBalanced(makeTree()), true)

	// One node.
	check("one node", `isBalanced(makeTree(1))`,
		isBalanced(makeTree(1)), true)

	// The classic examples.
	check("classic true", `isBalanced(makeTree(3, 9, 20, nil, nil, 15, 7))`,
		isBalanced(makeTree(3, 9, 20, nil, nil, 15, 7)), true)
	check("classic false", `isBalanced(makeTree(1, 2, 2, 3, 3, nil, nil, 4, 4))`,
		isBalanced(makeTree(1, 2, 2, 3, 3, nil, nil, 4, 4)), false)

	// A chain of three.
	check("chain of three", `isBalanced(makeTree(1, 2, nil, 3))`,
		isBalanced(makeTree(1, 2, nil, 3)), false)

	// Tricky: the root looks fine (both sides have height 3), but the node
	// 2 has a left side of height 2 and no right side.
	check("unbalanced below the root", `isBalanced(makeTree(1, 2, 3, 4, nil, nil, 5, 6, nil, nil, 7))`,
		isBalanced(makeTree(1, 2, 3, 4, nil, nil, 5, 6, nil, nil, 7)), false)

	// A difference of exactly 1 is allowed.
	check("difference of 1", `isBalanced(makeTree(1, 2, 3, 4))`,
		isBalanced(makeTree(1, 2, 3, 4)), true)

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
