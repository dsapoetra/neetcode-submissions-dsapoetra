// Problem: Linked List Cycle Detection
// Link:    https://neetcode.io/problems/linked-list-cycle-detection
// Source:  Data Structures & Algorithms/linked-list-cycle-detection/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./linked-list-cycle-detection
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

// Idea (slow and fast pointer): slow moves 1 step, fast moves 2 steps.
//   - no cycle: fast reaches the end of the list
//   - cycle:    fast keeps going around and catches up with slow,
//     like a faster runner on a round track
//
// Time: O(n), no extra memory.
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 1 step
		fast = fast.Next.Next // 2 steps

		if slow == fast {
			return true // they met: there is a cycle
		}
	}

	// fast reached the end: no cycle
	return false
}

// ===================== Test cases =====================

func main() {
	// makeCycleList(values, pos): the last node points back to the node at
	// index pos. pos = -1 means there is no cycle.

	// Example 1 from the problem: the tail points back to index 1.
	check("example 1: tail -> index 1", `hasCycle(makeCycleList([]int{1, 2, 3, 4}, 1))`,
		hasCycle(makeCycleList([]int{1, 2, 3, 4}, 1)), true)

	// Example 2 from the problem: a normal list.
	check("example 2: no cycle", `hasCycle(makeCycleList([]int{1, 2}, -1))`,
		hasCycle(makeCycleList([]int{1, 2}, -1)), false)

	// An empty list has no cycle.
	check("empty list", `hasCycle(nil)`,
		hasCycle(nil), false)

	// One node, no cycle.
	check("one node, no cycle", `hasCycle(makeCycleList([]int{1}, -1))`,
		hasCycle(makeCycleList([]int{1}, -1)), false)

	// One node that points to itself.
	check("one node pointing to itself", `hasCycle(makeCycleList([]int{1}, 0))`,
		hasCycle(makeCycleList([]int{1}, 0)), true)

	// Two nodes pointing at each other.
	check("two nodes in a circle", `hasCycle(makeCycleList([]int{1, 2}, 0))`,
		hasCycle(makeCycleList([]int{1, 2}, 0)), true)

	// The whole list is one big circle (tail -> head).
	check("tail -> head", `hasCycle(makeCycleList([]int{3, 2, 0, -4}, 0))`,
		hasCycle(makeCycleList([]int{3, 2, 0, -4}, 0)), true)

	// The last node points to itself.
	check("tail -> tail", `hasCycle(makeCycleList([]int{1, 2, 3}, 2))`,
		hasCycle(makeCycleList([]int{1, 2, 3}, 2)), true)

	// A long list without a cycle.
	check("long list, no cycle", `hasCycle(makeCycleList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -1))`,
		hasCycle(makeCycleList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -1)), false)

	// Equal values do NOT mean a cycle (we compare nodes, not values).
	check("equal values, no cycle", `hasCycle(makeCycleList([]int{1, 1, 1, 1}, -1))`,
		hasCycle(makeCycleList([]int{1, 1, 1, 1}, -1)), false)

	report()
}

// ===================== Helpers for the tests =====================

// makeCycleList builds a linked list from vals. If pos is 0 or more, the last
// node points back to the node at index pos (that makes a cycle).
// pos = -1 means no cycle.
func makeCycleList(vals []int, pos int) *ListNode {
	nodes := []*ListNode{}
	for _, v := range vals {
		nodes = append(nodes, &ListNode{Val: v})
	}
	if len(nodes) == 0 {
		return nil
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	if pos >= 0 {
		nodes[len(nodes)-1].Next = nodes[pos]
	}
	return nodes[0]
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
