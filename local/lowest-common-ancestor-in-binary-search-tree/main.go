// Problem: Lowest Common Ancestor of a Binary Search Tree
// Link:    https://neetcode.io/problems/lowest-common-ancestor-in-binary-search-tree
// Source:  Data Structures & Algorithms/lowest-common-ancestor-in-binary-search-tree/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./lowest-common-ancestor-in-binary-search-tree
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

// Idea: use the BST order. Start at the root:
//   - if p and q are both smaller, the answer is in the left subtree
//   - if p and q are both bigger, the answer is in the right subtree
//   - otherwise they split here (or one of them IS this node) -> answer
//
// Time: O(height of tree).
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left // both on the left
		} else if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right // both on the right
		} else {
			return cur // they split here
		}
	}
	return nil
}

// ===================== Test cases =====================

func main() {
	// The tree used by all tests (from the problem):
	//
	//	        5
	//	      /   \
	//	     3     8
	//	    / \   / \
	//	   1   4 7   9
	//	    \
	//	     2
	root := makeTree(5, 3, 8, 1, 4, 7, 9, nil, 2)
	fmt.Println("-- tree (level by level):", treeString(root))

	// lca is a small shortcut: find the two nodes by value and return the
	// VALUE of their lowest common ancestor.
	lca := func(p int, q int) int {
		return lowestCommonAncestor(root, findNode(root, p), findNode(root, q)).Val
	}

	// Example 1 from the problem: 3 is on the left, 8 on the right, so they split at the root.
	check("example 1: split at the root", `lca(3, 8)`,
		lca(3, 8), 5)

	// Example 2 from the problem: 4 is a child of 3, so 3 itself is the answer.
	check("example 2: one is the parent of the other", `lca(3, 4)`,
		lca(3, 4), 3)

	// Both in the left subtree.
	check("both on the left", `lca(1, 4)`,
		lca(1, 4), 3)

	// Both in the right subtree.
	check("both on the right", `lca(7, 9)`,
		lca(7, 9), 8)

	// 2 is deeper than 4, they still meet at 3.
	check("different depths", `lca(2, 4)`,
		lca(2, 4), 3)

	// Far apart: deepest left node and the rightmost node.
	check("far apart", `lca(2, 9)`,
		lca(2, 9), 5)

	// One of them is the root.
	check("one is the root", `lca(5, 9)`,
		lca(5, 9), 5)

	// 2 is the right child of 1.
	check("parent and child at the bottom", `lca(2, 1)`,
		lca(2, 1), 1)

	// The order of p and q must not matter.
	check("p and q swapped", `lca(8, 3)`,
		lca(8, 3), 5)

	// The same node twice: the node is its own ancestor.
	check("same node twice", `lca(7, 7)`,
		lca(7, 7), 7)

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

// findNode returns the node with the given value (or nil if there is none).
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if found := findNode(root.Left, val); found != nil {
		return found
	}
	return findNode(root.Right, val)
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
