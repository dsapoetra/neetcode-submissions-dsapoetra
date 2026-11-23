// Problem:    Reverse Nodes in k-Group
// Difficulty: Hard
// Link:       https://neetcode.io/problems/reverse-nodes-in-k-group
// Source:     Data Structures & Algorithms/reverse-nodes-in-k-group/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a linked list and a number k, reverse the nodes k at a time and
// return the new first node. If fewer than k nodes are left at the end,
// they stay in their original order.
//
// Example 1:
//
//	Input:  head = [1, 2, 3, 4, 5, 6], k = 3
//	Output: [3, 2, 1, 6, 5, 4]
//
// Example 2:
//
//	Input:  head = [1, 2, 3, 4, 5], k = 2
//	Output: [2, 1, 4, 3, 5]
//
// How to run (from the "local" folder):
//
//	go run ./135-reverse-nodes-in-k-group
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

// Idea: handle one group of k nodes at a time.
//  1. check that k nodes are left (if not, we are done)
//  2. reverse the arrows inside the group
//  3. connect the reversed group with the part before and after it
//
// groupPrev is always the node right before the group we are working on.
// Time: O(n), no extra memory.
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head} // fake node before the head
	groupPrev := dummy

	for {
		// 1) find the last node of this group
		kth := groupPrev
		for i := 0; i < k && kth != nil; i++ {
			kth = kth.Next
		}
		if kth == nil {
			break // fewer than k nodes left: leave them as they are
		}
		groupNext := kth.Next // first node AFTER the group

		// 2) reverse the group. Starting with prev = groupNext makes the
		//    old first node point to the rest of the list.
		prev := groupNext
		cur := groupPrev.Next
		for cur != groupNext {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

		// 3) the old first node is now the last node of the group
		oldFirst := groupPrev.Next
		groupPrev.Next = kth // kth is the new first node
		groupPrev = oldFirst // the next group starts after it
	}
	return dummy.Next
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: two full groups of 3.
	check("example 1", `listToSlice(reverseKGroup(makeList(1, 2, 3, 4, 5, 6), 3))`,
		listToSlice(reverseKGroup(makeList(1, 2, 3, 4, 5, 6), 3)), []int{3, 2, 1, 6, 5, 4})

	// Example 2 from the problem: the last node has no partner and stays.
	check("example 2: one node left over", `listToSlice(reverseKGroup(makeList(1, 2, 3, 4, 5), 2))`,
		listToSlice(reverseKGroup(makeList(1, 2, 3, 4, 5), 2)), []int{2, 1, 4, 3, 5})

	// Two nodes are left over and keep their order.
	check("two nodes left over", `listToSlice(reverseKGroup(makeList(1, 2, 3, 4, 5), 3))`,
		listToSlice(reverseKGroup(makeList(1, 2, 3, 4, 5), 3)), []int{3, 2, 1, 4, 5})

	// k = 1: nothing changes.
	check("k = 1", `listToSlice(reverseKGroup(makeList(1, 2, 3), 1))`,
		listToSlice(reverseKGroup(makeList(1, 2, 3), 1)), []int{1, 2, 3})

	// One node.
	check("one node", `listToSlice(reverseKGroup(makeList(1), 1))`,
		listToSlice(reverseKGroup(makeList(1), 1)), []int{1})

	// k = length: the whole list is reversed.
	check("k = length", `listToSlice(reverseKGroup(makeList(1, 2, 3, 4), 4))`,
		listToSlice(reverseKGroup(makeList(1, 2, 3, 4), 4)), []int{4, 3, 2, 1})

	// Two nodes, k = 2.
	check("two nodes", `listToSlice(reverseKGroup(makeList(1, 2), 2))`,
		listToSlice(reverseKGroup(makeList(1, 2), 2)), []int{2, 1})

	// k is bigger than the list: nothing changes.
	check("k bigger than the list", `listToSlice(reverseKGroup(makeList(1, 2, 3), 4))`,
		listToSlice(reverseKGroup(makeList(1, 2, 3), 4)), []int{1, 2, 3})

	// An empty list.
	check("empty list", `listToSlice(reverseKGroup(makeList(), 2))`,
		listToSlice(reverseKGroup(makeList(), 2)), []int{})

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
