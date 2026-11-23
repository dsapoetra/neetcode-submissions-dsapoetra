// Problem: Binary Tree Level Order Traversal
// Link:    https://neetcode.io/problems/level-order-traversal-of-binary-tree
// Source:  Data Structures & Algorithms/level-order-traversal-of-binary-tree/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./level-order-traversal-of-binary-tree
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

// Idea: BFS with a queue. The queue holds the nodes of the current level.
// For each level, take out exactly that many nodes and put their children in.
// Time: O(n).
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := []int{}
		size := len(queue) // number of nodes in this level

		for i := 0; i < size; i++ {
			// take the first node out of the queue
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			// its children belong to the next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: three full levels.
	check("example 1", levelOrder(makeTree(1, 2, 3, 4, 5, 6, 7)), [][]int{{1}, {2, 3}, {4, 5, 6, 7}})

	// Example 2 from the problem: one node.
	check("example 2: one node", levelOrder(makeTree(1)), [][]int{{1}})

	// Example 3 from the problem: an empty tree has no levels.
	check("example 3: empty tree", levelOrder(makeTree()), [][]int{})

	// The classic example: the left child 9 has no children.
	check("classic", levelOrder(makeTree(3, 9, 20, nil, nil, 15, 7)), [][]int{{3}, {9, 20}, {15, 7}})

	// A chain going left: one node per level.
	check("left chain", levelOrder(makeTree(1, 2, nil, 3)), [][]int{{1}, {2}, {3}})

	// A chain going right.
	check("right chain", levelOrder(makeTree(1, nil, 2, nil, 3)), [][]int{{1}, {2}, {3}})

	// Gaps in the last level: 4 and 5 still come left to right.
	check("gaps in last level", levelOrder(makeTree(1, 2, 3, nil, 4, nil, 5)), [][]int{{1}, {2, 3}, {4, 5}})

	// Negative values.
	check("negative values", levelOrder(makeTree(-1, -2, -3)), [][]int{{-1}, {-2, -3}})

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
