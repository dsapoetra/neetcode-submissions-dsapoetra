// Problem:    Serialize and Deserialize Binary Tree
// Difficulty: Hard
// Link:       https://neetcode.io/problems/serialize-and-deserialize-binary-tree
// Source:     Data Structures & Algorithms/serialize-and-deserialize-binary-tree/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Write two functions: serialize turns a binary tree into a string, and
// deserialize turns that string back into the same tree. The format of the
// string is up to you.
//
// Example 1:
//
//	Input:  root = [1, 2, 3, null, null, 4, 5]
//	Output: [1, 2, 3, null, null, 4, 5]      (after serialize and deserialize)
//
// Example 2:
//
//	Input:  root = []
//	Output: []
//
// How to run (from the "local" folder):
//
//	go run ./138-serialize-and-deserialize-binary-tree
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
	check("serialized format", `codec.serialize(makeTree(1, 2, 3))`,
		codec.serialize(makeTree(1, 2, 3)), "1,2,N,N,3,N,N")
	check("serialized empty tree", `codec.serialize(makeTree())`,
		codec.serialize(makeTree()), "N")

	// Example 1 from the problem.
	check("example 1", `roundTrip(1, 2, 3, nil, nil, 4, 5)`,
		roundTrip(1, 2, 3, nil, nil, 4, 5), "[1,2,3,null,null,4,5]")

	// Example 2 from the problem: an empty tree.
	check("example 2: empty tree", `roundTrip()`,
		roundTrip(), "[]")

	// One node.
	check("one node", `roundTrip(1)`,
		roundTrip(1), "[1]")

	// Only a left child / only a right child must not be mixed up.
	check("left child only", `roundTrip(1, 2)`,
		roundTrip(1, 2), "[1,2]")
	check("right child only", `roundTrip(1, nil, 2)`,
		roundTrip(1, nil, 2), "[1,null,2]")

	// Negative values (the "-" must survive).
	check("negative values", `roundTrip(-1, -2, -3)`,
		roundTrip(-1, -2, -3), "[-1,-2,-3]")

	// A chain going left.
	check("left chain", `roundTrip(1, 2, nil, 3, nil, 4)`,
		roundTrip(1, 2, nil, 3, nil, 4), "[1,2,null,3,null,4]")

	// Equal values.
	check("equal values", `roundTrip(1, 1, 1)`,
		roundTrip(1, 1, 1), "[1,1,1]")

	// Values with several digits.
	check("big values", `roundTrip(1000, -1000, 999)`,
		roundTrip(1000, -1000, 999), "[1000,-1000,999]")

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
