// Problem: Number of 1 Bits
// Link:    https://neetcode.io/problems/number-of-one-bits
// Source:  Data Structures & Algorithms/number-of-one-bits/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./number-of-one-bits
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: look at the last bit (n & 1), add it to the count,
// then shift n right by one to look at the next bit. Stop when n is 0.
func hammingWeight(n int) int {
	count := 0
	for n != 0 {
		count += n & 1 // 1 if the last bit is set, otherwise 0
		n = n >> 1     // drop the last bit
	}
	return count
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: 23 is 10111 in binary.
	check("example 1: 23 = 10111", hammingWeight(23), 4)

	// Example 2 from the problem: 31 bits, one of them is 0.
	check("example 2: 2147483645", hammingWeight(2147483645), 30)

	// Zero has no 1 bits.
	check("0", hammingWeight(0), 0)

	// 1 is just one bit.
	check("1", hammingWeight(1), 1)

	// Powers of two have exactly one 1 bit.
	check("8 = 1000", hammingWeight(8), 1)
	check("1024", hammingWeight(1024), 1)

	// 255 is 11111111.
	check("255 = 11111111", hammingWeight(255), 8)

	// 11 is 1011.
	check("11 = 1011", hammingWeight(11), 3)

	// All 32 bits set.
	check("all 32 bits", hammingWeight(4294967295), 32)

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
