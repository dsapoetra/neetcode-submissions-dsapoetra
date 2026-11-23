// Problem:    Min Stack
// Difficulty: Medium
// Link:       https://neetcode.io/problems/minimum-stack
// Source:     Data Structures & Algorithms/minimum-stack/submission-1.go (the solution below is a copy of it)
//
// THE PROBLEM
// Build a stack that can also report its smallest value. All four
// operations must run in O(1):
//
//	Push(val)   puts a value on top
//	Pop()       removes the top value
//	Top()       returns the top value
//	GetMin()    returns the smallest value in the stack
//
// Example:
//
//	Input:  Push(1), Push(2), Push(0), GetMin(), Pop(), Top(), GetMin()
//	Output: -, -, -, 0, -, 2, 1
//
// How to run (from the "local" folder):
//
//	go run ./052-minimum-stack
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: use two stacks of the same height.
//   values: the normal stack
//   mins:   mins[i] = the smallest value among values[0..i]
// When we push, the new minimum is the smaller of (new value, old minimum).
// When we pop, we pop both stacks, so mins always matches values.
// Every operation is O(1).

type MinStack struct {
	values []int
	mins   []int
}

func Constructor() MinStack {
	return MinStack{}
}

// Push adds val on top, and stores the new minimum next to it.
func (this *MinStack) Push(val int) {
	// the new minimum is val, unless the old minimum is smaller
	min := val
	// if the stack is not empty, compare with the minimum so far
	if len(this.mins) > 0 && this.mins[len(this.mins)-1] < min {
		min = this.mins[len(this.mins)-1]
	}
	// both stacks grow together, so they always have the same height
	this.values = append(this.values, val)
	this.mins = append(this.mins, min)
}

// Pop removes the top value, and the minimum that belonged to it.
func (this *MinStack) Pop() {
	this.values = this.values[:len(this.values)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

// Top returns the last pushed value without removing it.
func (this *MinStack) Top() int {
	return this.values[len(this.values)-1]
}

// GetMin returns the smallest value in the stack: the top of mins.
func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

// ===================== Test cases =====================

func main() {
	// The stack remembers what happened before, so the lines starting with
	// "--" show every Push and Pop in the output.

	// Example from the problem.
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(0)
	fmt.Println(`-- stack.Push(1); stack.Push(2); stack.Push(0)`)
	check("minimum of 1, 2, 0", `stack.GetMin()`,
		stack.GetMin(), 0)
	stack.Pop()
	fmt.Println(`-- stack.Pop()   (removes the 0)`)
	check("top after pop", `stack.Top()`,
		stack.Top(), 2)
	check("minimum after pop", `stack.GetMin()`,
		stack.GetMin(), 1)

	// The classic example with negative numbers.
	second := Constructor()
	second.Push(-2)
	second.Push(0)
	second.Push(-3)
	fmt.Println(`-- second.Push(-2); second.Push(0); second.Push(-3)`)
	check("minimum is -3", `second.GetMin()`,
		second.GetMin(), -3)
	second.Pop()
	fmt.Println(`-- second.Pop()   (removes the -3)`)
	check("top is 0", `second.Top()`,
		second.Top(), 0)
	check("minimum is back to -2", `second.GetMin()`,
		second.GetMin(), -2)

	// The minimum appears twice: popping one copy must keep the other.
	third := Constructor()
	third.Push(2)
	third.Push(2)
	third.Pop()
	fmt.Println(`-- third.Push(2); third.Push(2); third.Pop()`)
	check("duplicate minimum survives a pop", `third.GetMin()`,
		third.GetMin(), 2)

	// Numbers going up: the minimum stays the first number.
	third.Push(5)
	third.Push(9)
	fmt.Println(`-- third.Push(5); third.Push(9)`)
	check("top is the newest", `third.Top()`,
		third.Top(), 9)
	check("minimum is still the oldest", `third.GetMin()`,
		third.GetMin(), 2)

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
