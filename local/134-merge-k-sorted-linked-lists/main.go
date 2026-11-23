// Problem:    Merge K Sorted Lists
// Difficulty: Hard
// Link:       https://neetcode.io/problems/merge-k-sorted-linked-lists
// Source:     Data Structures & Algorithms/merge-k-sorted-linked-lists/submission-1.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given k sorted linked lists, merge them all into one sorted linked list
// and return its first node.
//
// Example 1:
//
//	Input:  lists = [[1, 2, 4], [1, 3, 5], [3, 6]]
//	Output: [1, 1, 2, 3, 3, 4, 5, 6]
//
// Example 2:
//
//	Input:  lists = []
//	Output: []
//
// How to run (from the "local" folder):
//
//	go run ./134-merge-k-sorted-linked-lists
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

// Idea: we already know how to merge TWO sorted lists.
// So merge the lists in pairs: list 0 with list 1, list 2 with list 3, ...
// That cuts the number of lists in half. Repeat until one list is left.
// (Merging everything into one big list one by one also works, but it
// walks over the big list again and again, which is too slow.)
// Time: O(total nodes * log k).
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		merged := []*ListNode{}

		// merge neighbors: (0,1), (2,3), (4,5), ...
		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				merged = append(merged, mergeTwo(lists[i], lists[i+1]))
			} else {
				// odd one out: no partner, keep it for the next round
				merged = append(merged, lists[i])
			}
		}

		lists = merged // now we have half as many lists
	}
	return lists[0]
}

// mergeTwo merges two sorted lists into one sorted list.
func mergeTwo(a *ListNode, b *ListNode) *ListNode {
	dummy := &ListNode{} // fake first node, makes the code simpler
	tail := dummy

	// always attach the smaller node
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
		}
		tail = tail.Next
	}

	// one list is finished: attach what is left of the other
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return dummy.Next
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", `listToSlice(mergeKLists([]*ListNode{makeList(1, 2, 4), makeList(1, 3, 5), makeList(3, 6)}))`,
		listToSlice(mergeKLists([]*ListNode{makeList(1, 2, 4), makeList(1, 3, 5), makeList(3, 6)})), []int{1, 1, 2, 3, 3, 4, 5, 6})

	// Example 2 from the problem: no lists at all.
	check("example 2: no lists", `listToSlice(mergeKLists([]*ListNode{}))`,
		listToSlice(mergeKLists([]*ListNode{})), []int{})

	// Example 3 from the problem: one empty list.
	check("example 3: one empty list", `listToSlice(mergeKLists([]*ListNode{makeList()}))`,
		listToSlice(mergeKLists([]*ListNode{makeList()})), []int{})

	// One list: it comes back unchanged.
	check("one list", `listToSlice(mergeKLists([]*ListNode{makeList(1, 2, 3)}))`,
		listToSlice(mergeKLists([]*ListNode{makeList(1, 2, 3)})), []int{1, 2, 3})

	// An empty list next to a normal one.
	check("empty + normal", `listToSlice(mergeKLists([]*ListNode{makeList(), makeList(1)}))`,
		listToSlice(mergeKLists([]*ListNode{makeList(), makeList(1)})), []int{1})

	// The classic example.
	check("classic", `listToSlice(mergeKLists([]*ListNode{makeList(1, 4, 5), makeList(1, 3, 4), makeList(2, 6)}))`,
		listToSlice(mergeKLists([]*ListNode{makeList(1, 4, 5), makeList(1, 3, 4), makeList(2, 6)})), []int{1, 1, 2, 3, 4, 4, 5, 6})

	// Negative values.
	check("negative values", `listToSlice(mergeKLists([]*ListNode{makeList(-3, -1), makeList(-2, 0)}))`,
		listToSlice(mergeKLists([]*ListNode{makeList(-3, -1), makeList(-2, 0)})), []int{-3, -2, -1, 0})

	// Lists with one node each, not in order.
	check("single nodes", `listToSlice(mergeKLists([]*ListNode{makeList(3), makeList(1), makeList(2)}))`,
		listToSlice(mergeKLists([]*ListNode{makeList(3), makeList(1), makeList(2)})), []int{1, 2, 3})

	// An odd number of lists: one list has no partner in the first round.
	check("five lists", `listToSlice(mergeKLists([]*ListNode{makeList(1), makeList(2), makeList(3), makeList(4), makeList(5)}))`,
		listToSlice(mergeKLists([]*ListNode{makeList(1), makeList(2), makeList(3), makeList(4), makeList(5)})), []int{1, 2, 3, 4, 5})

	// Equal values everywhere.
	check("equal values", `listToSlice(mergeKLists([]*ListNode{makeList(2, 2), makeList(2, 2)}))`,
		listToSlice(mergeKLists([]*ListNode{makeList(2, 2), makeList(2, 2)})), []int{2, 2, 2, 2})

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
