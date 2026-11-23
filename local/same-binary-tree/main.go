// Problem: Same Tree
// Link:    https://neetcode.io/problems/same-binary-tree
// Source:  Data Structures & Algorithms/same-binary-tree/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./same-binary-tree
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

// Idea: two trees are the same if the roots have the same value
// and the left subtrees are the same and the right subtrees are the same.
// Time: O(n).
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// both empty -> same
	if p == nil && q == nil {
		return true
	}
	// only one is empty, or the values differ -> not the same
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	// check both subtrees
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1: identical", `isSameTree(makeTree(1, 2, 3), makeTree(1, 2, 3))`,
		isSameTree(makeTree(1, 2, 3), makeTree(1, 2, 3)), true)

	// Example 2 from the problem: 7 is a left child in one tree, a right child in the other.
	check("example 2: left vs right child", `isSameTree(makeTree(4, 7), makeTree(4, nil, 7))`,
		isSameTree(makeTree(4, 7), makeTree(4, nil, 7)), false)

	// Example 3 from the problem: same values, children swapped.
	check("example 3: children swapped", `isSameTree(makeTree(1, 2, 3), makeTree(1, 3, 2))`,
		isSameTree(makeTree(1, 2, 3), makeTree(1, 3, 2)), false)

	// Two empty trees are the same.
	check("both empty", `isSameTree(makeTree(), makeTree())`,
		isSameTree(makeTree(), makeTree()), true)

	// One empty, one not.
	check("first empty", `isSameTree(makeTree(), makeTree(1))`,
		isSameTree(makeTree(), makeTree(1)), false)
	check("second empty", `isSameTree(makeTree(1), makeTree())`,
		isSameTree(makeTree(1), makeTree()), false)

	// One node each, same value.
	check("single nodes, same", `isSameTree(makeTree(1), makeTree(1))`,
		isSameTree(makeTree(1), makeTree(1)), true)

	// One node each, different value.
	check("single nodes, different", `isSameTree(makeTree(1), makeTree(2))`,
		isSameTree(makeTree(1), makeTree(2)), false)

	// One tree has an extra node.
	check("extra node", `isSameTree(makeTree(1, 2), makeTree(1, 2, 3))`,
		isSameTree(makeTree(1, 2), makeTree(1, 2, 3)), false)

	// Bigger identical trees.
	check("bigger identical", `isSameTree(makeTree(1, 2, 3, 4, nil, nil, 5), makeTree(1, 2, 3, 4, nil, nil, 5))`,
		isSameTree(makeTree(1, 2, 3, 4, nil, nil, 5), makeTree(1, 2, 3, 4, nil, nil, 5)), true)

	// The difference is deep down in the tree.
	check("different deep down", `isSameTree(makeTree(1, 2, 3, 4), makeTree(1, 2, 3, nil, 4))`,
		isSameTree(makeTree(1, 2, 3, 4), makeTree(1, 2, 3, nil, 4)), false)

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
