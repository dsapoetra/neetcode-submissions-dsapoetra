// Problem:    Binary Tree Right Side View
// Difficulty: Medium
// Link:       https://neetcode.io/problems/binary-tree-right-side-view
// Source:     Data Structures & Algorithms/binary-tree-right-side-view/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Imagine standing on the right side of a binary tree. Return the values of
// the nodes you can see, from top to bottom (the last node of every level).
//
// Example 1:
//
//	Input:  root = [1, 2, 3]
//	Output: [1, 3]
//
// Example 2:
//
//	Input:  root = [1, 2, 3, 4, 5, 6, 7]
//	Output: [1, 3, 7]
//
// How to run (from the "local" folder):
//
//	go run ./058-binary-tree-right-side-view
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

// Idea: standing on the right side we see the LAST node of every level.
// So walk through the tree level by level (BFS with a queue) and keep the
// last value of each level.
// Time: O(n).
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue) // number of nodes in this level

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			// the last node of the level is the one we can see
			if i == size-1 {
				res = append(res, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", `rightSideView(makeTree(1, 2, 3))`,
		rightSideView(makeTree(1, 2, 3)), []int{1, 3})

	// Example 2 from the problem: a full tree.
	check("example 2", `rightSideView(makeTree(1, 2, 3, 4, 5, 6, 7))`,
		rightSideView(makeTree(1, 2, 3, 4, 5, 6, 7)), []int{1, 3, 7})

	// An empty tree.
	check("empty tree", `rightSideView(makeTree())`,
		rightSideView(makeTree()), []int{})

	// One node.
	check("one node", `rightSideView(makeTree(1))`,
		rightSideView(makeTree(1)), []int{1})

	// The classic example.
	check("classic", `rightSideView(makeTree(1, 2, 3, nil, 5, nil, 4))`,
		rightSideView(makeTree(1, 2, 3, nil, 5, nil, 4)), []int{1, 3, 4})

	// Tricky: the LEFT side is deeper, so its last node (4) is visible.
	check("left side is deeper", `rightSideView(makeTree(1, 2, 3, 4))`,
		rightSideView(makeTree(1, 2, 3, 4)), []int{1, 3, 4})

	// Only a left child: we can see it.
	check("only a left child", `rightSideView(makeTree(1, 2))`,
		rightSideView(makeTree(1, 2)), []int{1, 2})

	// Only a right child.
	check("only a right child", `rightSideView(makeTree(1, nil, 3))`,
		rightSideView(makeTree(1, nil, 3)), []int{1, 3})

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
