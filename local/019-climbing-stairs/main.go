// Problem:    Climbing Stairs
// Difficulty: Easy
// Link:       https://neetcode.io/problems/climbing-stairs
// Source:     Data Structures & Algorithms/climbing-stairs/submission-1.go (the solution below is a copy of it)
//
// THE PROBLEM
// You climb a staircase with n steps. Each time you can go up 1 step or 2
// steps. Return the number of different ways to reach the top.
//
// Example 1:
//
//	Input:  n = 2
//	Output: 2             (1+1, 2)
//
// Example 2:
//
//	Input:  n = 3
//	Output: 3             (1+1+1, 1+2, 2+1)
//
// How to run (from the "local" folder):
//
//	go run ./019-climbing-stairs
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: to reach step i we come from step i-1 or step i-2, so
// ways(i) = ways(i-1) + ways(i-2). That is the Fibonacci sequence.
// We only need the last two values, not a whole array.
// Time: O(n).
func climbStairs(n int) int {
	prev := 1 // ways to reach step i-2
	cur := 1  // ways to reach step i-1

	for i := 2; i <= n; i++ {
		next := prev + cur
		prev = cur
		cur = next
	}
	return cur
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 1+1 or 2.
	check("example 1: n = 2", `climbStairs(2)`,
		climbStairs(2), 2)

	// Example 2 from the problem: 1+1+1, 1+2, 2+1.
	check("example 2: n = 3", `climbStairs(3)`,
		climbStairs(3), 3)

	// One step: only one way.
	check("n = 1", `climbStairs(1)`,
		climbStairs(1), 1)

	// From here on it is the Fibonacci sequence: 5, 8, 13, ...
	check("n = 4", `climbStairs(4)`,
		climbStairs(4), 5)
	check("n = 5", `climbStairs(5)`,
		climbStairs(5), 8)
	check("n = 6", `climbStairs(6)`,
		climbStairs(6), 13)

	// Bigger numbers.
	check("n = 10", `climbStairs(10)`,
		climbStairs(10), 89)
	check("n = 30", `climbStairs(30)`,
		climbStairs(30), 1346269)

	// The biggest input the problem allows.
	check("n = 45", `climbStairs(45)`,
		climbStairs(45), 1836311903)

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
