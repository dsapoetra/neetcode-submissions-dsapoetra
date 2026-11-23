// Problem:    Reorder List
// Difficulty: Medium
// Link:       https://neetcode.io/problems/reorder-linked-list
// Source:     Data Structures & Algorithms/reorder-linked-list/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a linked list  L0 -> L1 -> ... -> Ln, reorder it in place to
// L0 -> Ln -> L1 -> Ln-1 -> L2 -> ...  (first, last, second, second-last, ...).
// Nothing is returned; the list itself is changed.
//
// Example 1:
//
//	Input:  head = [2, 4, 6, 8]
//	Output: [2, 8, 4, 6]
//
// Example 2:
//
//	Input:  head = [2, 4, 6, 8, 10]
//	Output: [2, 10, 4, 8, 6]
//
// How to run (from the "local" folder):
//
//	go run ./046-reorder-linked-list
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ListNode is one node of a singly linked list.
// NeetCode gives us this type; locally we have to define it ourselves.
type ListNode struct {
	Val  int
	Next *ListNode
}

// ===================== Solution =====================

// Idea: put all the nodes in a slice so we can reach them by index.
// Then use two pointers, one at the front and one at the back, and
// link them in the order: first, last, second, second-last, ...
// Time: O(n), extra space: O(n).
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// collect all nodes in a slice
	nodes := []*ListNode{}
	for cur := head; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
	}

	front := 0
	back := len(nodes) - 1
	for front < back {
		// front node points to the back node
		nodes[front].Next = nodes[back]
		front++
		if front == back {
			break
		}
		// back node points to the next front node
		nodes[back].Next = nodes[front]
		back--
	}

	// the node in the middle is now the last node
	nodes[front].Next = nil
}

// ===================== Test cases =====================

func main() {
	// reorderList changes the list in place and returns nothing, so this
	// small shortcut builds the list, reorders it and returns the values.
	reorder := func(vals ...int) []int {
		head := makeList(vals...)
		reorderList(head)
		return listToSlice(head)
	}

	// Example 1 from the problem: first, last, second, second-last.
	check("example 1: even length", `reorder(2, 4, 6, 8)`,
		reorder(2, 4, 6, 8), []int{2, 8, 4, 6})

	// Example 2 from the problem: the middle node ends up last.
	check("example 2: odd length", `reorder(2, 4, 6, 8, 10)`,
		reorder(2, 4, 6, 8, 10), []int{2, 10, 4, 8, 6})

	// An empty list.
	check("empty list", `reorder()`,
		reorder(), []int{})

	// One node.
	check("one node", `reorder(1)`,
		reorder(1), []int{1})

	// Two nodes: nothing changes.
	check("two nodes", `reorder(1, 2)`,
		reorder(1, 2), []int{1, 2})

	// Three nodes.
	check("three nodes", `reorder(1, 2, 3)`,
		reorder(1, 2, 3), []int{1, 3, 2})

	// Four nodes.
	check("four nodes", `reorder(1, 2, 3, 4)`,
		reorder(1, 2, 3, 4), []int{1, 4, 2, 3})

	// Six nodes.
	check("six nodes", `reorder(1, 2, 3, 4, 5, 6)`,
		reorder(1, 2, 3, 4, 5, 6), []int{1, 6, 2, 5, 3, 4})

	// Seven nodes.
	check("seven nodes", `reorder(1, 2, 3, 4, 5, 6, 7)`,
		reorder(1, 2, 3, 4, 5, 6, 7), []int{1, 7, 2, 6, 3, 5, 4})

	report()
}

// ===================== Helpers for the tests =====================

// makeList builds a linked list from the values.
// Example: makeList(1, 2, 3) gives 1 -> 2 -> 3. makeList() gives an empty list (nil).
func makeList(vals ...int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

// listToSlice turns a linked list into a slice, so it is easy to print and compare.
// It stops after 1000 nodes so a list with a cycle cannot loop forever.
func listToSlice(head *ListNode) []int {
	vals := []int{}
	for cur := head; cur != nil && len(vals) < 1000; cur = cur.Next {
		vals = append(vals, cur.Val)
	}
	return vals
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
