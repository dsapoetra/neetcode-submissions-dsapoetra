// Problem:    Counting Bits
// Difficulty: Easy
// Link:       https://neetcode.io/problems/counting-bits
// Source:     Data Structures & Algorithms/counting-bits/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a number n, return a list of length n + 1 where entry i is the
// number of 1 bits in the binary form of i.
//
// Example 1:
//
//	Input:  n = 4
//	Output: [0, 1, 1, 2, 1]      (0, 1, 10, 11, 100 in binary)
//
// How to run (from the "local" folder):
//
//	go run ./024-counting-bits
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: i/2 is i with its last bit removed. So:
//
//	ones in i = ones in i/2 + the last bit of i (i % 2)
//
// We already know the answer for i/2 because it is smaller than i.
// Example: 6 = 110 -> ones(3 = 11) + 0 = 2
// Time: O(n).
func countBits(n int) []int {
	bits := make([]int, n+1) // bits[0] is 0
	for i := 1; i <= n; i++ {
		bits[i] = bits[i/2] + i%2
	}
	return bits
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 0, 1, 10, 11, 100 in binary.
	check("example 1: n = 4", `countBits(4)`,
		countBits(4), []int{0, 1, 1, 2, 1})

	// Only the number 0.
	check("n = 0", `countBits(0)`,
		countBits(0), []int{0})

	// 0 and 1.
	check("n = 1", `countBits(1)`,
		countBits(1), []int{0, 1})

	// 0, 1, 10.
	check("n = 2", `countBits(2)`,
		countBits(2), []int{0, 1, 1})

	// Up to 101.
	check("n = 5", `countBits(5)`,
		countBits(5), []int{0, 1, 1, 2, 1, 2})

	// 8 is 1000: back to a single 1 bit.
	check("n = 8", `countBits(8)`,
		countBits(8), []int{0, 1, 1, 2, 1, 2, 2, 3, 1})

	// 15 is 1111: four 1 bits.
	check("n = 15", `countBits(15)`,
		countBits(15), []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4})

	// The answer always has n + 1 numbers.
	check("length is n + 1", `len(countBits(100))`,
		len(countBits(100)), 101)

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
