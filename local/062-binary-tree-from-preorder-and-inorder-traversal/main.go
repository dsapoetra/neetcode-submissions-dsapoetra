// Problem:    Construct Binary Tree from Preorder and Inorder Traversal
// Difficulty: Medium
// Link:       https://neetcode.io/problems/binary-tree-from-preorder-and-inorder-traversal
// Source:     Data Structures & Algorithms/binary-tree-from-preorder-and-inorder-traversal/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given the preorder (node, left, right) and the inorder (left, node, right)
// list of the same binary tree with unique values, rebuild the tree and
// return its root.
//
// Example 1:
//
//	Input:  preorder = [1, 2, 3, 4], inorder = [2, 1, 3, 4]
//	Output: [1, 2, 3, null, null, null, 4]
//
// Example 2:
//
//	Input:  preorder = [1], inorder = [1]
//	Output: [1]
//
// How to run (from the "local" folder):
//
//	go run ./062-binary-tree-from-preorder-and-inorder-traversal
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
	// no values left: an empty subtree
	if len(preorder) == 0 {
		return nil
	}

	// the first value in preorder is always the root
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// find the root in inorder
	mid := 0
	for inorder[mid] != rootVal {
		mid++
	}
	// mid is also the number of nodes in the left subtree

	// left subtree: the next "mid" values of preorder (after the root),
	// and everything before the root in inorder
	root.Left = buildTree(preorder[1:1+mid], inorder[:mid])
	// right subtree: all the values that are left
	root.Right = buildTree(preorder[1+mid:], inorder[mid+1:])
	return root
}

// ===================== Test cases =====================

func main() {
	// The result is written level by level, the same way NeetCode shows trees.

	// Example 1 from the problem.
	check("example 1", `treeString(buildTree([]int{1, 2, 3, 4}, []int{2, 1, 3, 4}))`,
		treeString(buildTree([]int{1, 2, 3, 4}, []int{2, 1, 3, 4})), "[1,2,3,null,null,null,4]")

	// Example 2 from the problem: one node.
	check("example 2: one node", `treeString(buildTree([]int{1}, []int{1}))`,
		treeString(buildTree([]int{1}, []int{1})), "[1]")

	// The classic example.
	check("classic", `treeString(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))`,
		treeString(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})), "[3,9,20,null,null,15,7]")

	// Everything goes left: inorder is preorder backwards.
	check("left chain", `treeString(buildTree([]int{1, 2, 3}, []int{3, 2, 1}))`,
		treeString(buildTree([]int{1, 2, 3}, []int{3, 2, 1})), "[1,2,null,3]")

	// Everything goes right: inorder is the same as preorder.
	check("right chain", `treeString(buildTree([]int{1, 2, 3}, []int{1, 2, 3}))`,
		treeString(buildTree([]int{1, 2, 3}, []int{1, 2, 3})), "[1,null,2,null,3]")

	// No nodes.
	check("empty tree", `treeString(buildTree([]int{}, []int{}))`,
		treeString(buildTree([]int{}, []int{})), "[]")

	// A full tree with 3 levels.
	check("full tree", `treeString(buildTree([]int{4, 2, 1, 3, 6, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7}))`,
		treeString(buildTree([]int{4, 2, 1, 3, 6, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7})), "[4,2,6,1,3,5,7]")

	// Negative values.
	check("negative values", `treeString(buildTree([]int{-1, -2, -3}, []int{-2, -1, -3}))`,
		treeString(buildTree([]int{-1, -2, -3}, []int{-2, -1, -3})), "[-1,-2,-3]")

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
