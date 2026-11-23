// Problem: Construct Binary Tree from Preorder and Inorder Traversal
// Link:    https://neetcode.io/problems/binary-tree-from-preorder-and-inorder-traversal
// Source:  Data Structures & Algorithms/binary-tree-from-preorder-and-inorder-traversal/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./binary-tree-from-preorder-and-inorder-traversal
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

// Idea:
//   - preorder is: root, left subtree, right subtree
//     -> the FIRST value of preorder is the root
//   - inorder is:  left subtree, root, right subtree
//     -> everything before the root in inorder is the left subtree,
//     everything after it is the right subtree
//
// So: take the root, find it in inorder, split both lists, and build
// the left and right subtrees the same way (recursion).
// Time: O(n^2) because we search for the root in inorder each time.
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// find the root in inorder
	mid := 0
	for inorder[mid] != rootVal {
		mid++
	}
	// mid is also the number of nodes in the left subtree

	root.Left = buildTree(preorder[1:1+mid], inorder[:mid])
	root.Right = buildTree(preorder[1+mid:], inorder[mid+1:])
	return root
}

// ===================== Test cases =====================

func main() {
	// The result is written level by level, the same way NeetCode shows trees.

	// Example 1 from the problem.
	check("example 1", treeString(buildTree([]int{1, 2, 3, 4}, []int{2, 1, 3, 4})), "[1,2,3,null,null,null,4]")

	// Example 2 from the problem: one node.
	check("example 2: one node", treeString(buildTree([]int{1}, []int{1})), "[1]")

	// The classic example.
	check("classic", treeString(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})), "[3,9,20,null,null,15,7]")

	// Everything goes left: inorder is preorder backwards.
	check("left chain", treeString(buildTree([]int{1, 2, 3}, []int{3, 2, 1})), "[1,2,null,3]")

	// Everything goes right: inorder is the same as preorder.
	check("right chain", treeString(buildTree([]int{1, 2, 3}, []int{1, 2, 3})), "[1,null,2,null,3]")

	// No nodes.
	check("empty tree", treeString(buildTree([]int{}, []int{})), "[]")

	// A full tree with 3 levels.
	check("full tree", treeString(buildTree([]int{4, 2, 1, 3, 6, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7})), "[4,2,6,1,3,5,7]")

	// Negative values.
	check("negative values", treeString(buildTree([]int{-1, -2, -3}, []int{-2, -1, -3})), "[-1,-2,-3]")

	report()
}

// ===================== Helpers for the tests =====================

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
