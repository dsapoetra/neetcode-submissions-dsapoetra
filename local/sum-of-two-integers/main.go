// Problem: Sum of Two Integers
// Link:    https://neetcode.io/problems/sum-of-two-integers
// Source:  Data Structures & Algorithms/sum-of-two-integers/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./sum-of-two-integers
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: add like on paper, but in binary and without using +.
//
//	a ^ b        adds each bit WITHOUT the carry (1+1 gives 0)
//	(a & b) << 1 is the carry (1+1 carries a 1 to the next position)
//
// Then we still have to add the carry, so repeat until no carry is left.
func getSum(a int, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b // sum without carry
		b = carry // the carry still has to be added
	}
	return a
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1: 1 + 1", getSum(1, 1), 2)

	// Example 2 from the problem.
	check("example 2: 4 + 7", getSum(4, 7), 11)

	// Zeros.
	check("0 + 0", getSum(0, 0), 0)
	check("0 + 5", getSum(0, 5), 5)
	check("5 + 0", getSum(5, 0), 5)

	// A number and its negative.
	check("-1 + 1", getSum(-1, 1), 0)
	check("1000 + -1000", getSum(1000, -1000), 0)

	// Two negatives.
	check("-5 + -7", getSum(-5, -7), -12)

	// Positive + negative.
	check("10 + -3", getSum(10, -3), 7)
	check("-10 + 3", getSum(-10, 3), -7)

	// Several carries in a row: 7 + 1 = 0111 + 0001 = 1000.
	check("carry chain: 7 + 1", getSum(7, 1), 8)

	// Bigger numbers.
	check("123 + 456", getSum(123, 456), 579)

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
