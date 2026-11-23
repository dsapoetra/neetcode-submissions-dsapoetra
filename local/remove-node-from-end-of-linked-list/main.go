// Problem: Remove Nth Node From End of List
// Link:    https://neetcode.io/problems/remove-node-from-end-of-linked-list
// Source:  Data Structures & Algorithms/remove-node-from-end-of-linked-list/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./remove-node-from-end-of-linked-list
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

// Idea (two pointers): move "fast" n steps ahead first. Then move "slow"
// and "fast" together. When fast reaches the last node, slow is right
// before the node we must remove.
// A dummy node before head makes removing the head itself easy.
// Time: O(n).
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := dummy

	// move fast n steps ahead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// move both until fast is at the last node
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// slow.Next is the node to remove: skip over it
	slow.Next = slow.Next.Next
	return dummy.Next
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: the 2nd from the end is 3.
	check("example 1", `listToSlice(removeNthFromEnd(makeList(1, 2, 3, 4), 2))`,
		listToSlice(removeNthFromEnd(makeList(1, 2, 3, 4), 2)), []int{1, 2, 4})

	// Example 2 from the problem: the only node is removed.
	check("example 2: only node", `listToSlice(removeNthFromEnd(makeList(5), 1))`,
		listToSlice(removeNthFromEnd(makeList(5), 1)), []int{})

	// Example 3 from the problem: remove the head of a two-node list.
	check("example 3: remove the head", `listToSlice(removeNthFromEnd(makeList(1, 2), 2))`,
		listToSlice(removeNthFromEnd(makeList(1, 2), 2)), []int{2})

	// Remove the tail of a two-node list.
	check("remove the tail", `listToSlice(removeNthFromEnd(makeList(1, 2), 1))`,
		listToSlice(removeNthFromEnd(makeList(1, 2), 1)), []int{1})

	// The classic example.
	check("classic", `listToSlice(removeNthFromEnd(makeList(1, 2, 3, 4, 5), 2))`,
		listToSlice(removeNthFromEnd(makeList(1, 2, 3, 4, 5), 2)), []int{1, 2, 3, 5})

	// n = length of the list: remove the first node.
	check("n = length (first node)", `listToSlice(removeNthFromEnd(makeList(1, 2, 3, 4, 5), 5))`,
		listToSlice(removeNthFromEnd(makeList(1, 2, 3, 4, 5), 5)), []int{2, 3, 4, 5})

	// n = 1: remove the last node.
	check("n = 1 (last node)", `listToSlice(removeNthFromEnd(makeList(1, 2, 3, 4, 5), 1))`,
		listToSlice(removeNthFromEnd(makeList(1, 2, 3, 4, 5), 1)), []int{1, 2, 3, 4})

	// Remove the middle node.
	check("middle node", `listToSlice(removeNthFromEnd(makeList(1, 2, 3), 2))`,
		listToSlice(removeNthFromEnd(makeList(1, 2, 3), 2)), []int{1, 3})

	// Equal values: exactly one node must go.
	check("equal values", `listToSlice(removeNthFromEnd(makeList(7, 7, 7), 3))`,
		listToSlice(removeNthFromEnd(makeList(7, 7, 7), 3)), []int{7, 7})

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
