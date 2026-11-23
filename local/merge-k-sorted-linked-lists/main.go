// Problem: Merge K Sorted Lists
// Link:    https://neetcode.io/problems/merge-k-sorted-linked-lists
// Source:  Data Structures & Algorithms/merge-k-sorted-linked-lists/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./merge-k-sorted-linked-lists
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
	check("example 1",
		listToSlice(mergeKLists([]*ListNode{makeList(1, 2, 4), makeList(1, 3, 5), makeList(3, 6)})),
		[]int{1, 1, 2, 3, 3, 4, 5, 6})

	// Example 2 from the problem: no lists at all.
	check("example 2: no lists", listToSlice(mergeKLists([]*ListNode{})), []int{})

	// Example 3 from the problem: one empty list.
	check("example 3: one empty list", listToSlice(mergeKLists([]*ListNode{makeList()})), []int{})

	// One list: it comes back unchanged.
	check("one list", listToSlice(mergeKLists([]*ListNode{makeList(1, 2, 3)})), []int{1, 2, 3})

	// An empty list next to a normal one.
	check("empty + normal", listToSlice(mergeKLists([]*ListNode{makeList(), makeList(1)})), []int{1})

	// The classic example.
	check("classic",
		listToSlice(mergeKLists([]*ListNode{makeList(1, 4, 5), makeList(1, 3, 4), makeList(2, 6)})),
		[]int{1, 1, 2, 3, 4, 4, 5, 6})

	// Negative values.
	check("negative values", listToSlice(mergeKLists([]*ListNode{makeList(-3, -1), makeList(-2, 0)})), []int{-3, -2, -1, 0})

	// Lists with one node each, not in order.
	check("single nodes", listToSlice(mergeKLists([]*ListNode{makeList(3), makeList(1), makeList(2)})), []int{1, 2, 3})

	// An odd number of lists: one list has no partner in the first round.
	check("five lists",
		listToSlice(mergeKLists([]*ListNode{makeList(1), makeList(2), makeList(3), makeList(4), makeList(5)})),
		[]int{1, 2, 3, 4, 5})

	// Equal values everywhere.
	check("equal values", listToSlice(mergeKLists([]*ListNode{makeList(2, 2), makeList(2, 2)})), []int{2, 2, 2, 2})

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
