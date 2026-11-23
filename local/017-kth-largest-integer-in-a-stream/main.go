// Problem:    Kth Largest Element in a Stream
// Difficulty: Easy
// Link:       https://neetcode.io/problems/kth-largest-integer-in-a-stream
// Source:     Data Structures & Algorithms/kth-largest-integer-in-a-stream/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Build a structure that receives numbers one by one and always knows the
// k-th largest number so far (in sorted order, duplicates count):
//
//	Constructor(k, nums)  starts with the numbers in nums
//	Add(val)              adds a number and returns the k-th largest
//
// Example:
//
//	Input:  KthLargest(3, [1, 2, 3, 3]), Add(3), Add(5), Add(6), Add(7), Add(8)
//	Output: -, 3, 3, 3, 5, 6
//
// How to run (from the "local" folder):
//
//	go run ./017-kth-largest-integer-in-a-stream
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: we only ever need the k biggest numbers. Keep them in a sorted
// list (smallest first). Then the k-th largest number of everything is
// simply the FIRST number of that list.
// Add: put the new number at its sorted place; if the list now has more
// than k numbers, throw away the smallest one.
// Add is O(k). (A min-heap would make it O(log k), but needs more code.)

type KthLargest struct {
	k   int
	top []int // the k biggest numbers so far, sorted, smallest first
}

// Constructor adds the starting numbers one by one, using the same Add.
func Constructor(k int, nums []int) KthLargest {
	this := KthLargest{k: k}
	for _, num := range nums {
		this.Add(num)
	}
	return this
}

func (this *KthLargest) Add(val int) int {
	// add at the end, then move it left until the list is sorted again
	this.top = append(this.top, val)
	for i := len(this.top) - 1; i > 0 && this.top[i] < this.top[i-1]; i-- {
		this.top[i], this.top[i-1] = this.top[i-1], this.top[i]
	}

	// keep only the k biggest
	if len(this.top) > this.k {
		// more than k numbers: drop the smallest (it is at the front)
		this.top = this.top[1:]
	}
	// the smallest of the k biggest = the k-th largest
	return this.top[0]
}

// ===================== Test cases =====================

func main() {
	// The object remembers all numbers, so the lines starting with "--"
	// show how it was created.

	// Example from the problem: always the 3rd largest.
	kth := Constructor(3, []int{1, 2, 3, 3})
	fmt.Println(`-- kth := Constructor(3, []int{1, 2, 3, 3})`)
	check("numbers 1,2,3,3,3: third largest is 3", `kth.Add(3)`,
		kth.Add(3), 3)
	check("add 5: top three are 5,3,3", `kth.Add(5)`,
		kth.Add(5), 3)
	check("add 6: top three are 6,5,3", `kth.Add(6)`,
		kth.Add(6), 3)
	check("add 7: top three are 7,6,5", `kth.Add(7)`,
		kth.Add(7), 5)
	check("add 8: top three are 8,7,6", `kth.Add(8)`,
		kth.Add(8), 6)

	// The classic example.
	classic := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(`-- classic := Constructor(3, []int{4, 5, 8, 2})`)
	check("classic: add 3", `classic.Add(3)`,
		classic.Add(3), 4)
	check("classic: add 5", `classic.Add(5)`,
		classic.Add(5), 5)
	check("classic: add 10", `classic.Add(10)`,
		classic.Add(10), 5)
	check("classic: add 9", `classic.Add(9)`,
		classic.Add(9), 8)
	check("classic: add 4", `classic.Add(4)`,
		classic.Add(4), 8)

	// Start with no numbers, k = 1: always the biggest number so far.
	biggest := Constructor(1, []int{})
	fmt.Println(`-- biggest := Constructor(1, []int{})`)
	check("first number", `biggest.Add(-3)`,
		biggest.Add(-3), -3)
	check("a bigger number", `biggest.Add(-2)`,
		biggest.Add(-2), -2)
	check("a smaller number changes nothing", `biggest.Add(-4)`,
		biggest.Add(-4), -2)
	check("a new biggest number", `biggest.Add(0)`,
		biggest.Add(0), 0)

	// Start with fewer than k numbers.
	second := Constructor(2, []int{0})
	fmt.Println(`-- second := Constructor(2, []int{0})`)
	check("second largest of 0,-1", `second.Add(-1)`,
		second.Add(-1), -1)
	check("second largest of 1,0,-1", `second.Add(1)`,
		second.Add(1), 0)
	check("a small number changes nothing", `second.Add(-2)`,
		second.Add(-2), 0)
	check("second largest of 3,1,0,...", `second.Add(3)`,
		second.Add(3), 1)

	report()
}

// ===================== Helpers for the tests =====================

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
