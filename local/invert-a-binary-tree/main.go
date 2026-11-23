// Problem: Invert Binary Tree
// Link:    https://neetcode.io/problems/invert-a-binary-tree
// Source:  Data Structures & Algorithms/invert-a-binary-tree/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./invert-a-binary-tree
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TreeNode is one node of a binary tree.
// NeetCode gives us this type; locally we have to define it ourselves.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ===================== Solution =====================

// Idea: swap the left and right child of every node (recursion).
// Time: O(n).
func invertTree(root *TreeNode) *TreeNode {
	// empty tree: nothing to do
	if root == nil {
		return nil
	}

	// swap the two children
	root.Left, root.Right = root.Right, root.Left

	// do the same for both subtrees
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// ===================== Test cases =====================

func main() {
	// Trees are written level by level, the same way NeetCode shows them.

	// Example 1 from the problem.
	check("example 1", treeString(invertTree(makeTree(1, 2, 3, 4, 5, 6, 7))), "[1,3,2,7,6,5,4]")

	// Example 2 from the problem.
	check("example 2", treeString(invertTree(makeTree(3, 2, 1))), "[3,1,2]")

	// Example 3 from the problem: an empty tree.
	check("example 3: empty tree", treeString(invertTree(makeTree())), "[]")

	// One node: nothing to swap.
	check("one node", treeString(invertTree(makeTree(1))), "[1]")

	// Only a left child: it becomes the right child.
	check("left child only", treeString(invertTree(makeTree(1, 2))), "[1,null,2]")

	// Only a right child: it becomes the left child.
	check("right child only", treeString(invertTree(makeTree(1, nil, 2))), "[1,2]")

	// The classic example.
	check("classic", treeString(invertTree(makeTree(4, 2, 7, 1, 3, 6, 9))), "[4,7,2,9,6,3,1]")

	// A chain going left becomes a chain going right.
	check("left chain", treeString(invertTree(makeTree(1, 2, nil, 3))), "[1,null,2,null,3]")

	// Inverting twice gives the original tree back.
	check("invert twice", treeString(invertTree(invertTree(makeTree(1, 2, 3, 4)))), "[1,2,3,4]")

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

// treeString writes a tree level by level, like "[1,2,3,null,4]".
// It is the opposite of makeTree, so trees are easy to print and compare.
func treeString(root *TreeNode) string {
	parts := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			parts = append(parts, "null")
			continue
		}
		parts = append(parts, strconv.Itoa(node.Val))
		queue = append(queue, node.Left, node.Right)
	}
	// the "null"s at the end carry no information, drop them
	for len(parts) > 0 && parts[len(parts)-1] == "null" {
		parts = parts[:len(parts)-1]
	}
	return "[" + strings.Join(parts, ",") + "]"
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
