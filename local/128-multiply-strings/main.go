// Problem:    Multiply Strings
// Difficulty: Medium
// Link:       https://neetcode.io/problems/multiply-strings
// Source:     Data Structures & Algorithms/multiply-strings/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Two non-negative numbers are given as strings. Return their product, also
// as a string. The numbers can be far too long for a normal integer, and
// converting the whole string to a number is not allowed.
//
// Example 1:
//
//	Input:  num1 = "3", num2 = "4"
//	Output: "12"
//
// Example 2:
//
//	Input:  num1 = "111", num2 = "222"
//	Output: "24642"
//
// How to run (from the "local" folder):
//
//	go run ./128-multiply-strings
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: multiply like on paper. The digit at position i of num1 times the
// digit at position j of num2 belongs at position i + j + 1 of the
// result (and its carry at position i + j).
// A number with a digits times a number with b digits has at most a + b
// digits, so the result fits in a list of that length.
// Time: O(len(num1) * len(num2)).
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	result := make([]int, len(num1)+len(num2))

	// go from the last digits to the first, like on paper
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			product := int(num1[i]-'0') * int(num2[j]-'0')

			sum := product + result[i+j+1] // add what is already there
			result[i+j+1] = sum % 10       // keep one digit
			result[i+j] += sum / 10        // carry the rest to the left
		}
	}

	// turn the digits into text, without the zeros at the front
	text := ""
	for i, digit := range result {
		if text == "" && digit == 0 && i < len(result)-1 {
			continue
		}
		text += string(rune('0' + digit))
	}
	return text
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", `multiply("3", "4")`,
		multiply("3", "4"), "12")

	// Example 2 from the problem.
	check("example 2", `multiply("111", "222")`,
		multiply("111", "222"), "24642")

	// The classic examples.
	check("classic 1", `multiply("2", "3")`,
		multiply("2", "3"), "6")
	check("classic 2", `multiply("123", "456")`,
		multiply("123", "456"), "56088")

	// Anything times zero is "0" (not "000").
	check("zero first", `multiply("0", "999")`,
		multiply("0", "999"), "0")
	check("zero second", `multiply("999", "0")`,
		multiply("999", "0"), "0")

	// A carry makes a second digit.
	check("9 * 9", `multiply("9", "9")`,
		multiply("9", "9"), "81")
	check("99 * 99", `multiply("99", "99")`,
		multiply("99", "99"), "9801")

	// Times one.
	check("times one", `multiply("1", "12345")`,
		multiply("1", "12345"), "12345")

	// Zeros at the end must stay.
	check("100 * 100", `multiply("100", "100")`,
		multiply("100", "100"), "10000")

	// Too big for a normal 64-bit number... almost: this is why we use strings.
	check("big numbers", `multiply("123456789", "987654321")`,
		multiply("123456789", "987654321"), "121932631112635269")

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
