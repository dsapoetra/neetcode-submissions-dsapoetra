// Problem:    Binary Tree Maximum Path Sum
// Difficulty: Hard
// Link:       https://neetcode.io/problems/binary-tree-maximum-path-sum
// Source:     Data Structures & Algorithms/binary-tree-maximum-path-sum/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// A path in a binary tree is a chain of connected nodes; it can start and
// end anywhere and does not have to go through the root. Return the biggest
// sum of node values over all paths (a path has at least one node).
//
// Example 1:
//
//	Input:  root = [1, 2, 3]
//	Output: 6             (2 -> 1 -> 3)
//
// Example 2:
//
//	Input:  root = [-15, 10, 20, null, null, 15, 5, -5]
//	Output: 40            (15 -> 20 -> 5)
//
// How to run (from the "local" folder):
//
//	go run ./137-binary-tree-maximum-path-sum
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

// Idea: for every node, imagine it is the TOP of the path. Then the
// path is: best path going down-left + the node + best path going down-right.
// We try every node as the top and keep the biggest total.
//
// The helper returns the best path that starts at a node and goes DOWN
// on only one side, because the parent can only connect to one side.
// A side with a negative sum is ignored (we use 0 instead).
// Time: O(n).
func maxPathSum(root *TreeNode) int {
	best := root.Val

	var bestDown func(node *TreeNode) int
	bestDown = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// best path going down each side (0 if it would make things worse)
		left := bestDown(node.Left)
		if left < 0 {
			left = 0
		}
		right := bestDown(node.Right)
		if right < 0 {
			right = 0
		}

		// path with this node as the top, using both sides
		if node.Val+left+right > best {
			best = node.Val + left + right
		}

		// for the parent: this node plus only the better side
		if left > right {
			return node.Val + left
		}
		return node.Val + right
	}

	bestDown(root)
	return best
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 2 -> 1 -> 3.
	check("example 1", `maxPathSum(makeTree(1, 2, 3))`,
		maxPathSum(makeTree(1, 2, 3)), 6)

	// Example 2 from the problem: 15 -> 20 -> 5. The path does not use the root.
	check("example 2: path below the root", `maxPathSum(makeTree(-15, 10, 20, nil, nil, 15, 5, -5))`,
		maxPathSum(makeTree(-15, 10, 20, nil, nil, 15, 5, -5)), 40)

	// One node.
	check("one node", `maxPathSum(makeTree(5))`,
		maxPathSum(makeTree(5)), 5)

	// One negative node: the path must have at least one node.
	check("one negative node", `maxPathSum(makeTree(-3))`,
		maxPathSum(makeTree(-3)), -3)

	// All negative: the best path is the single biggest node.
	check("all negative", `maxPathSum(makeTree(-2, -1))`,
		maxPathSum(makeTree(-2, -1)), -1)

	// The classic example: 15 -> 20 -> 7.
	check("classic", `maxPathSum(makeTree(-10, 9, 20, nil, nil, 15, 7))`,
		maxPathSum(makeTree(-10, 9, 20, nil, nil, 15, 7)), 42)

	// A negative child is left out.
	check("skip a negative child", `maxPathSum(makeTree(2, -1))`,
		maxPathSum(makeTree(2, -1)), 2)

	// Skip the negative left child, take root + right.
	check("skip one side", `maxPathSum(makeTree(1, -2, 3))`,
		maxPathSum(makeTree(1, -2, 3)), 4)

	// A chain: the path can go straight down.
	check("chain", `maxPathSum(makeTree(1, 2, nil, 3))`,
		maxPathSum(makeTree(1, 2, nil, 3)), 6)

	// A negative node in the middle is still worth passing through: 10 -> -1 -> 10.
	check("through a negative node", `maxPathSum(makeTree(-1, 10, 10))`,
		maxPathSum(makeTree(-1, 10, 10)), 19)

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
