// Problem:    Add Two Numbers
// Difficulty: Medium
// Link:       https://neetcode.io/problems/add-two-numbers
// Source:     Data Structures & Algorithms/add-two-numbers/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Two numbers are given as linked lists with one digit per node, stored
// BACKWARDS (ones first). Add the two numbers and return the sum in the
// same format.
//
// Example 1:
//
//	Input:  l1 = [1, 2, 3], l2 = [4, 5, 6]
//	Output: [5, 7, 9]     (321 + 654 = 975)
//
// Example 2:
//
//	Input:  l1 = [9], l2 = [9]
//	Output: [8, 1]        (9 + 9 = 18)
//
// How to run (from the "local" folder):
//
//	go run ./049-add-two-numbers
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

// Idea: add like on paper. The lists store the digits backwards (ones
// first), so we can go from left to right: add the two digits plus the
// carry, write down the last digit of the sum, and carry the rest.
// Keep going while there are digits left OR there is still a carry.
// Time: O(longer list).
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // fake first node, makes the code simpler
	tail := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		tail.Next = &ListNode{Val: sum % 10} // write the last digit
		tail = tail.Next
		carry = sum / 10 // carry the rest (0 or 1)
	}
	return dummy.Next
}

// ===================== Test cases =====================

func main() {
	// The digits are stored backwards: makeList(1, 2, 3) is the number 321.

	// Example 1 from the problem: 321 + 654 = 975.
	check("example 1", `listToSlice(addTwoNumbers(makeList(1, 2, 3), makeList(4, 5, 6)))`,
		listToSlice(addTwoNumbers(makeList(1, 2, 3), makeList(4, 5, 6))), []int{5, 7, 9})

	// Example 2 from the problem: 9 + 9 = 18.
	check("example 2: carry makes a new digit", `listToSlice(addTwoNumbers(makeList(9), makeList(9)))`,
		listToSlice(addTwoNumbers(makeList(9), makeList(9))), []int{8, 1})

	// The classic example: 342 + 465 = 807.
	check("classic", `listToSlice(addTwoNumbers(makeList(2, 4, 3), makeList(5, 6, 4)))`,
		listToSlice(addTwoNumbers(makeList(2, 4, 3), makeList(5, 6, 4))), []int{7, 0, 8})

	// 0 + 0 = 0.
	check("zero plus zero", `listToSlice(addTwoNumbers(makeList(0), makeList(0)))`,
		listToSlice(addTwoNumbers(makeList(0), makeList(0))), []int{0})

	// 9999999 + 9999 = 10009998: the carry runs through the whole number.
	check("long carry chain", `listToSlice(addTwoNumbers(makeList(9, 9, 9, 9, 9, 9, 9), makeList(9, 9, 9, 9)))`,
		listToSlice(addTwoNumbers(makeList(9, 9, 9, 9, 9, 9, 9), makeList(9, 9, 9, 9))), []int{8, 9, 9, 9, 0, 0, 0, 1})

	// Different lengths: 1 + 99 = 100.
	check("different lengths", `listToSlice(addTwoNumbers(makeList(1), makeList(9, 9)))`,
		listToSlice(addTwoNumbers(makeList(1), makeList(9, 9))), []int{0, 0, 1})

	// 5 + 5 = 10.
	check("5 + 5", `listToSlice(addTwoNumbers(makeList(5), makeList(5)))`,
		listToSlice(addTwoNumbers(makeList(5), makeList(5))), []int{0, 1})

	// Adding zero changes nothing: 0 + 37 = 37.
	check("adding zero", `listToSlice(addTwoNumbers(makeList(0), makeList(7, 3)))`,
		listToSlice(addTwoNumbers(makeList(0), makeList(7, 3))), []int{7, 3})

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
