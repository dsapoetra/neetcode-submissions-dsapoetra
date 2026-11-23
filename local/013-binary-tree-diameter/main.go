// Problem:    Diameter of Binary Tree
// Difficulty: Easy
// Link:       https://neetcode.io/problems/binary-tree-diameter
// Source:     Data Structures & Algorithms/binary-tree-diameter/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// The diameter of a binary tree is the length of the longest path between
// any two nodes, counted in edges. The path does not have to go through the
// root. Return the diameter.
//
// Example 1:
//
//	Input:  root = [1, null, 2, 3, 4, 5]
//	Output: 3
//
// Example 2:
//
//	Input:  root = [1, 2, 3]
//	Output: 2
//
// How to run (from the "local" folder):
//
//	go run ./013-binary-tree-diameter
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

// Idea: the longest path that bends at a node goes down its left side as
// deep as possible and down its right side as deep as possible:
//
//	path through node = height of left subtree + height of right subtree
//
// So compute the height of every subtree (recursion) and, at every node,
// check whether left height + right height is the best so far.
// Time: O(n).
func diameterOfBinaryTree(root *TreeNode) int {
	best := 0

	// height returns the number of nodes on the longest way down from node.
	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := height(node.Left)
		right := height(node.Right)

		// the path that bends at this node has left + right edges
		if left+right > best {
			best = left + right
		}

		// height of this node = 1 + the taller side
		if left > right {
			return left + 1
		}
		return right + 1
	}

	height(root)
	return best
}

// ===================== Test cases =====================

func main() {
	// The diameter is counted in EDGES (lines between nodes), not nodes.

	// Example 1 from the problem: the longest path does not go through the root.
	check("example 1", `diameterOfBinaryTree(makeTree(1, nil, 2, 3, 4, 5))`,
		diameterOfBinaryTree(makeTree(1, nil, 2, 3, 4, 5)), 3)

	// Example 2 from the problem: 2 - 1 - 3.
	check("example 2", `diameterOfBinaryTree(makeTree(1, 2, 3))`,
		diameterOfBinaryTree(makeTree(1, 2, 3)), 2)

	// One node: no edges.
	check("one node", `diameterOfBinaryTree(makeTree(1))`,
		diameterOfBinaryTree(makeTree(1)), 0)

	// An empty tree.
	check("empty tree", `diameterOfBinaryTree(makeTree())`,
		diameterOfBinaryTree(makeTree()), 0)

	// The classic example: 4 - 2 - 1 - 3 (or 5 - 2 - 1 - 3).
	check("classic", `diameterOfBinaryTree(makeTree(1, 2, 3, 4, 5))`,
		diameterOfBinaryTree(makeTree(1, 2, 3, 4, 5)), 3)

	// Two nodes.
	check("two nodes", `diameterOfBinaryTree(makeTree(1, 2))`,
		diameterOfBinaryTree(makeTree(1, 2)), 1)

	// The longest path (5 - 3 - 2 - 4 - 6) is completely inside the left subtree.
	check("path below the root", `diameterOfBinaryTree(makeTree(1, 2, nil, 3, 4, 5, nil, nil, 6))`,
		diameterOfBinaryTree(makeTree(1, 2, nil, 3, 4, 5, nil, nil, 6)), 4)

	// A chain: the path is the whole chain.
	check("chain", `diameterOfBinaryTree(makeTree(1, 2, nil, 3, nil, 4))`,
		diameterOfBinaryTree(makeTree(1, 2, nil, 3, nil, 4)), 3)

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
