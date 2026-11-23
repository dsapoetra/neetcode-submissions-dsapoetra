// Problem:    Reverse Linked List
// Difficulty: Easy
// Link:       https://neetcode.io/problems/reverse-a-linked-list
// Source:     Data Structures & Algorithms/reverse-a-linked-list/submission-2.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given the first node of a singly linked list, reverse the list and return
// its new first node.
//
// Example 1:
//
//	Input:  head = [0, 1, 2, 3]
//	Output: [3, 2, 1, 0]
//
// Example 2:
//
//	Input:  head = []
//	Output: []
//
// How to run (from the "local" folder):
//
//	go run ./007-reverse-a-linked-list
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
	check("example 1", `listToSlice(reverseList(makeList(0, 1, 2, 3)))`,
		listToSlice(reverseList(makeList(0, 1, 2, 3))), []int{3, 2, 1, 0})

	// Example 2 from the problem: an empty list stays empty.
	check("example 2: empty list", `listToSlice(reverseList(makeList()))`,
		listToSlice(reverseList(makeList())), []int{})

	// One node: nothing changes.
	check("one node", `listToSlice(reverseList(makeList(1)))`,
		listToSlice(reverseList(makeList(1))), []int{1})

	// Two nodes swap places.
	check("two nodes", `listToSlice(reverseList(makeList(1, 2)))`,
		listToSlice(reverseList(makeList(1, 2))), []int{2, 1})

	// Five nodes.
	check("five nodes", `listToSlice(reverseList(makeList(1, 2, 3, 4, 5)))`,
		listToSlice(reverseList(makeList(1, 2, 3, 4, 5))), []int{5, 4, 3, 2, 1})

	// All values the same: looks the same after reversing.
	check("all the same", `listToSlice(reverseList(makeList(7, 7, 7)))`,
		listToSlice(reverseList(makeList(7, 7, 7))), []int{7, 7, 7})

	// Negative values.
	check("negative values", `listToSlice(reverseList(makeList(-1, 0, 1)))`,
		listToSlice(reverseList(makeList(-1, 0, 1))), []int{1, 0, -1})

	// Reversing twice gives the original list back.
	check("reverse twice", `listToSlice(reverseList(reverseList(makeList(1, 2, 3))))`,
		listToSlice(reverseList(reverseList(makeList(1, 2, 3)))), []int{1, 2, 3})

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
