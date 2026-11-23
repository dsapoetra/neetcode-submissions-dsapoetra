// Problem: Reverse Linked List
// Link:    https://neetcode.io/problems/reverse-a-linked-list
// Source:  Data Structures & Algorithms/reverse-a-linked-list/submission-2.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./reverse-a-linked-list
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

// Idea: walk through the list and turn every arrow around, so each
// node points to the node BEFORE it instead of the node after it.
// Time: O(n).
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode // the node before cur (nothing at the start)
	cur := head

	for cur != nil {
		next := cur.Next // save the next node before we lose it
		cur.Next = prev  // turn the arrow around
		prev = cur       // move both pointers one step forward
		cur = next
	}

	// prev is the old last node, which is the new head
	return prev
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", listToSlice(reverseList(makeList(0, 1, 2, 3))), []int{3, 2, 1, 0})

	// Example 2 from the problem: an empty list stays empty.
	check("example 2: empty list", listToSlice(reverseList(makeList())), []int{})

	// One node: nothing changes.
	check("one node", listToSlice(reverseList(makeList(1))), []int{1})

	// Two nodes swap places.
	check("two nodes", listToSlice(reverseList(makeList(1, 2))), []int{2, 1})

	// Five nodes.
	check("five nodes", listToSlice(reverseList(makeList(1, 2, 3, 4, 5))), []int{5, 4, 3, 2, 1})

	// All values the same: looks the same after reversing.
	check("all the same", listToSlice(reverseList(makeList(7, 7, 7))), []int{7, 7, 7})

	// Negative values.
	check("negative values", listToSlice(reverseList(makeList(-1, 0, 1))), []int{1, 0, -1})

	// Reversing twice gives the original list back.
	check("reverse twice", listToSlice(reverseList(reverseList(makeList(1, 2, 3)))), []int{1, 2, 3})

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
