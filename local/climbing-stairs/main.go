// Problem: Climbing Stairs
// Link:    https://neetcode.io/problems/climbing-stairs
// Source:  Data Structures & Algorithms/climbing-stairs/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./climbing-stairs
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
	check("example 1: n = 2", climbStairs(2), 2)

	// Example 2 from the problem: 1+1+1, 1+2, 2+1.
	check("example 2: n = 3", climbStairs(3), 3)

	// One step: only one way.
	check("n = 1", climbStairs(1), 1)

	// From here on it is the Fibonacci sequence: 5, 8, 13, ...
	check("n = 4", climbStairs(4), 5)
	check("n = 5", climbStairs(5), 8)
	check("n = 6", climbStairs(6), 13)

	// Bigger numbers.
	check("n = 10", climbStairs(10), 89)
	check("n = 30", climbStairs(30), 1346269)

	// The biggest input the problem allows.
	check("n = 45", climbStairs(45), 1836311903)

	report()
}

// ===================== Helpers for the tests =====================

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
