// Problem:    Pow(x, n)
// Difficulty: Medium
// Link:       https://neetcode.io/problems/pow-x-n
// Source:     Data Structures & Algorithms/pow-x-n/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a decimal number x and a whole number n (which may be negative),
// return x to the power of n.
//
// Example 1:
//
//	Input:  x = 2.00000, n = 5
//	Output: 32.00000
//
// Example 2:
//
//	Input:  x = 1.10000, n = 10
//	Output: 2.59374
//
// Example 3:
//
//	Input:  x = 2.00000, n = -3
//	Output: 0.12500
//
// How to run (from the "local" folder):
//
//	go run ./127-pow-x-n
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (fast power): x^n = (x^(n/2))^2. So we only need to compute the
// half power once and square it. If n is odd, multiply by one extra x.
// That needs about log2(n) multiplications instead of n.
// A negative n means 1 / x^(-n).
// Time: O(log n).
func myPow(x float64, n int) float64 {
	// negative power: compute the positive power and take 1 / result
	if n < 0 {
		// (-n is now positive)
		return 1 / power(x, -n)
	}
	return power(x, n)
}

// power computes x^n for n >= 0.
func power(x float64, n int) float64 {
	// anything to the power 0 is 1; this also stops the recursion
	if n == 0 {
		return 1
	}

	// compute the half power only ONCE and reuse it
	// (n/2 is rounded down, so for n = 5 this is x^2)
	half := power(x, n/2)
	// even n: x^n = half * half
	if n%2 == 0 {
		return half * half
	}
	// odd n: one x is missing because n/2 was rounded down
	return half * half * x
}

// ===================== Test cases =====================

func main() {
	// rounded calls myPow and rounds the answer to 5 decimals, because
	// decimal numbers are never perfectly exact in a computer.
	rounded := func(x float64, n int) string {
		return fmt.Sprintf("%.5f", myPow(x, n))
	}

	// Example 1 from the problem.
	check("example 1", `rounded(2.0, 5)`,
		rounded(2.0, 5), "32.00000")

	// Example 2 from the problem.
	check("example 2", `rounded(1.1, 10)`,
		rounded(1.1, 10), "2.59374")

	// Example 3 from the problem: a negative power is 1 / (2*2*2).
	check("example 3: negative power", `rounded(2.0, -3)`,
		rounded(2.0, -3), "0.12500")

	// The classic examples.
	check("2^10", `rounded(2.0, 10)`,
		rounded(2.0, 10), "1024.00000")
	check("2.1^3", `rounded(2.1, 3)`,
		rounded(2.1, 3), "9.26100")
	check("2^-2", `rounded(2.0, -2)`,
		rounded(2.0, -2), "0.25000")

	// Anything to the power 0 is 1.
	check("power 0", `rounded(5.0, 0)`,
		rounded(5.0, 0), "1.00000")

	// 0 to a positive power is 0.
	check("base 0", `rounded(0.0, 5)`,
		rounded(0.0, 5), "0.00000")

	// 1 to a huge power is still 1 (and must not take forever).
	check("huge power", `rounded(1.0, 1000000000)`,
		rounded(1.0, 1000000000), "1.00000")

	// A negative base: odd power stays negative, even power becomes positive.
	check("negative base, odd power", `rounded(-2.0, 3)`,
		rounded(-2.0, 3), "-8.00000")
	check("negative base, even power", `rounded(-2.0, 2)`,
		rounded(-2.0, 2), "4.00000")

	// A base below 1 gets smaller.
	check("base below 1", `rounded(0.5, 2)`,
		rounded(0.5, 2), "0.25000")

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
