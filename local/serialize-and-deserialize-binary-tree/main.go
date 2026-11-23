// Problem: Serialize and Deserialize Binary Tree
// Link:    https://neetcode.io/problems/serialize-and-deserialize-binary-tree
// Source:  Data Structures & Algorithms/serialize-and-deserialize-binary-tree/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./serialize-and-deserialize-binary-tree
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

// Idea: write the tree in preorder (node, left, right) separated by
// commas, and write "N" for an empty child.
// Example:   1
//           / \
//          2   3     ->  "1,2,N,N,3,N,N"
// Because the empty children are written too, we can rebuild the tree
// by reading the values back in the same order.

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "N"
	}
	// node, then left subtree, then right subtree
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	parts := strings.Split(data, ",")
	index := 0 // the next value to read

	var build func() *TreeNode
	build = func() *TreeNode {
		value := parts[index]
		index++

		if value == "N" {
			return nil
		}

		number, _ := strconv.Atoi(value)
		node := &TreeNode{Val: number}
		node.Left = build()  // same order as serialize: left first
		node.Right = build() // then right
		return node
	}

	return build()
}

// ===================== Test cases =====================

func main() {
	codec := Constructor()

	// roundTrip builds a tree, turns it into a string, turns the string back
	// into a tree, and writes that tree level by level. We must get the same
	// tree back.
	roundTrip := func(vals ...interface{}) string {
		text := codec.serialize(makeTree(vals...))
		return treeString(codec.deserialize(text))
	}

	// What the serialized string looks like: preorder, "N" for an empty child.
	check("serialized format", codec.serialize(makeTree(1, 2, 3)), "1,2,N,N,3,N,N")
	check("serialized empty tree", codec.serialize(makeTree()), "N")

	// Example 1 from the problem.
	check("example 1", roundTrip(1, 2, 3, nil, nil, 4, 5), "[1,2,3,null,null,4,5]")

	// Example 2 from the problem: an empty tree.
	check("example 2: empty tree", roundTrip(), "[]")

	// One node.
	check("one node", roundTrip(1), "[1]")

	// Only a left child / only a right child must not be mixed up.
	check("left child only", roundTrip(1, 2), "[1,2]")
	check("right child only", roundTrip(1, nil, 2), "[1,null,2]")

	// Negative values (the "-" must survive).
	check("negative values", roundTrip(-1, -2, -3), "[-1,-2,-3]")

	// A chain going left.
	check("left chain", roundTrip(1, 2, nil, 3, nil, 4), "[1,2,null,3,null,4]")

	// Equal values.
	check("equal values", roundTrip(1, 1, 1), "[1,1,1]")

	// Values with several digits.
	check("big values", roundTrip(1000, -1000, 999), "[1000,-1000,999]")

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
