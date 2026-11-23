// Problem: Merge Two Sorted Lists
// Link:    https://neetcode.io/problems/merge-two-sorted-linked-lists
// Source:  Data Structures & Algorithms/merge-two-sorted-linked-lists/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./merge-two-sorted-linked-lists
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

// Idea: compare the first nodes of both lists and always attach the
// smaller one to the result. A dummy first node saves us from special
// code for the very first node.
// Time: O(n + m).
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // fake first node
	tail := dummy        // last node of the result so far

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	// one list is finished: attach what is left of the other
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next // skip the fake node
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", `listToSlice(mergeTwoLists(makeList(1, 2, 4), makeList(1, 3, 5)))`,
		listToSlice(mergeTwoLists(makeList(1, 2, 4), makeList(1, 3, 5))), []int{1, 1, 2, 3, 4, 5})

	// Example 2 from the problem: the first list is empty.
	check("example 2: first empty", `listToSlice(mergeTwoLists(makeList(), makeList(1, 2)))`,
		listToSlice(mergeTwoLists(makeList(), makeList(1, 2))), []int{1, 2})

	// Example 3 from the problem: both lists are empty.
	check("example 3: both empty", `listToSlice(mergeTwoLists(makeList(), makeList()))`,
		listToSlice(mergeTwoLists(makeList(), makeList())), []int{})

	// The second list is empty.
	check("second empty", `listToSlice(mergeTwoLists(makeList(1), makeList()))`,
		listToSlice(mergeTwoLists(makeList(1), makeList())), []int{1})

	// Everything in list 1 is smaller than everything in list 2.
	check("list 1 completely first", `listToSlice(mergeTwoLists(makeList(1, 2, 3), makeList(4, 5, 6)))`,
		listToSlice(mergeTwoLists(makeList(1, 2, 3), makeList(4, 5, 6))), []int{1, 2, 3, 4, 5, 6})

	// Everything in list 2 is smaller than everything in list 1.
	check("list 2 completely first", `listToSlice(mergeTwoLists(makeList(4, 5, 6), makeList(1, 2, 3)))`,
		listToSlice(mergeTwoLists(makeList(4, 5, 6), makeList(1, 2, 3))), []int{1, 2, 3, 4, 5, 6})

	// The lists take turns (like a zipper).
	check("taking turns", `listToSlice(mergeTwoLists(makeList(1, 3, 5), makeList(2, 4, 6)))`,
		listToSlice(mergeTwoLists(makeList(1, 3, 5), makeList(2, 4, 6))), []int{1, 2, 3, 4, 5, 6})

	// Equal values in both lists: all of them must be kept.
	check("equal values", `listToSlice(mergeTwoLists(makeList(1, 1), makeList(1, 1)))`,
		listToSlice(mergeTwoLists(makeList(1, 1), makeList(1, 1))), []int{1, 1, 1, 1})

	// Negative values.
	check("negative values", `listToSlice(mergeTwoLists(makeList(-5, 0), makeList(-3, 7)))`,
		listToSlice(mergeTwoLists(makeList(-5, 0), makeList(-3, 7))), []int{-5, -3, 0, 7})

	// Very different lengths.
	check("different lengths", `listToSlice(mergeTwoLists(makeList(1), makeList(2, 3, 4, 5)))`,
		listToSlice(mergeTwoLists(makeList(1), makeList(2, 3, 4, 5))), []int{1, 2, 3, 4, 5})

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
