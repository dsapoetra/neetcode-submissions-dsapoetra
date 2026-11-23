// Problem: Reverse Bits
// Link:    https://neetcode.io/problems/reverse-bits
// Source:  Data Structures & Algorithms/reverse-bits/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./reverse-bits
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: do this 32 times: take the last bit of n and push it onto the
// end of the result. The first bit we take ends up at the far left,
// so the bits come out reversed.
func reverseBits(n int) int {
	result := 0
	for i := 0; i < 32; i++ {
		lastBit := n & 1
		result = (result << 1) | lastBit // make room, then add the bit
		n = n >> 1                       // drop the bit we just used
	}
	return result
}

// ===================== Test cases =====================

func main() {
	// The number is always treated as 32 bits.

	// Example 1 from the problem: ...00010101 becomes 10101000...
	check("example 1: 21", `reverseBits(21)`,
		reverseBits(21), 2818572288)

	// Zero stays zero.
	check("0", `reverseBits(0)`,
		reverseBits(0), 0)

	// The lowest bit becomes the highest bit.
	check("1 -> highest bit", `reverseBits(1)`,
		reverseBits(1), 2147483648)

	// And the other way around.
	check("highest bit -> 1", `reverseBits(2147483648)`,
		reverseBits(2147483648), 1)

	// All 32 bits set: stays the same.
	check("all bits set", `reverseBits(4294967295)`,
		reverseBits(4294967295), 4294967295)

	// The classic examples.
	check("classic 43261596", `reverseBits(43261596)`,
		reverseBits(43261596), 964176192)
	check("classic 4294967293", `reverseBits(4294967293)`,
		reverseBits(4294967293), 3221225471)

	// The second-lowest bit becomes the second-highest bit.
	check("2", `reverseBits(2)`,
		reverseBits(2), 1073741824)

	// First and last bit set: a "bit palindrome" stays the same.
	check("bit palindrome", `reverseBits(2147483649)`,
		reverseBits(2147483649), 2147483649)

	// Reversing twice gives the original number back.
	check("reverse twice", `reverseBits(reverseBits(12345))`,
		reverseBits(reverseBits(12345)), 12345)

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
