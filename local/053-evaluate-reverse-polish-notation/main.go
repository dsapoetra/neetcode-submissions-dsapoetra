// Problem:    Evaluate Reverse Polish Notation
// Difficulty: Medium
// Link:       https://neetcode.io/problems/evaluate-reverse-polish-notation
// Source:     Data Structures & Algorithms/evaluate-reverse-polish-notation/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Evaluate an expression written in Reverse Polish Notation, where the
// operator comes AFTER its two numbers ("1 2 +" means 1 + 2). The operators
// are + - * /, and division cuts off the decimals.
//
// Example 1:
//
//	Input:  tokens = ["1", "2", "+", "3", "*", "4", "-"]
//	Output: 5             (((1 + 2) * 3) - 4)
//
// How to run (from the "local" folder):
//
//	go run ./053-evaluate-reverse-polish-notation
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"strconv"
)

// ===================== Solution =====================

// Idea: use a stack of numbers.
//   - a number   -> push it
//   - an operator -> pop the last two numbers, calculate, push the result
//
// Careful with the order: the number popped FIRST is the RIGHT side
// (for "a b -" the stack gives b first, and the answer is a - b).
// Time: O(n).
func evalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		// is this token an operator?
		if token == "+" || token == "-" || token == "*" || token == "/" {
			// pop the right side, then the left side
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			result := 0
			// calculate "left operator right"
			if token == "+" {
				result = left + right
			} else if token == "-" {
				result = left - right
			} else if token == "*" {
				result = left * right
			} else {
				result = left / right // Go cuts off the decimals (rounds toward zero)
			}
			// the result goes back on the stack; it can be used by a later operator
			stack = append(stack, result)
		} else {
			// a number: turn the text into an int and push it
			number, _ := strconv.Atoi(token)
			stack = append(stack, number)
		}
	}
	// at the end only one number is left: the answer
	return stack[0]
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: ((1 + 2) * 3) - 4 = 5.
	check("example 1", `evalRPN([]string{"1", "2", "+", "3", "*", "4", "-"})`,
		evalRPN([]string{"1", "2", "+", "3", "*", "4", "-"}), 5)

	// The classic examples.
	check("classic 1: (2 + 1) * 3", `evalRPN([]string{"2", "1", "+", "3", "*"})`,
		evalRPN([]string{"2", "1", "+", "3", "*"}), 9)
	check("classic 2: 4 + (13 / 5)", `evalRPN([]string{"4", "13", "5", "/", "+"})`,
		evalRPN([]string{"4", "13", "5", "/", "+"}), 6)
	check("classic 3: long", `evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})`,
		evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}), 22)

	// Only one number.
	check("one number", `evalRPN([]string{"42"})`,
		evalRPN([]string{"42"}), 42)

	// The order matters for minus: 7 - 2, not 2 - 7.
	check("7 - 2", `evalRPN([]string{"7", "2", "-"})`,
		evalRPN([]string{"7", "2", "-"}), 5)
	check("2 - 7", `evalRPN([]string{"2", "7", "-"})`,
		evalRPN([]string{"2", "7", "-"}), -5)

	// Division cuts off the decimals (rounds toward zero).
	check("7 / 2", `evalRPN([]string{"7", "2", "/"})`,
		evalRPN([]string{"7", "2", "/"}), 3)
	check("-7 / 2", `evalRPN([]string{"-7", "2", "/"})`,
		evalRPN([]string{"-7", "2", "/"}), -3)

	// A negative number is a number, not a minus sign.
	check("3 * -4", `evalRPN([]string{"3", "-4", "*"})`,
		evalRPN([]string{"3", "-4", "*"}), -12)

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
