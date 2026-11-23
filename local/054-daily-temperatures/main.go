// Problem:    Daily Temperatures
// Difficulty: Medium
// Link:       https://neetcode.io/problems/daily-temperatures
// Source:     Data Structures & Algorithms/daily-temperatures/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// temperatures[i] is the temperature on day i. For every day, return how
// many days you have to wait for a warmer day (0 if there is none).
//
// Example 1:
//
//	Input:  temperatures = [30, 38, 30, 36, 35, 40, 28]
//	Output: [1, 4, 1, 2, 1, 0, 0]
//
// Example 2:
//
//	Input:  temperatures = [22, 21, 20]
//	Output: [0, 0, 0]
//
// How to run (from the "local" folder):
//
//	go run ./054-daily-temperatures
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: keep a stack of days that are still WAITING for a warmer day.
// The temperatures of the waiting days get lower towards the top.
// For each new day: every waiting day that is colder than today has found
// its answer (today), so pop it and write down the distance.
// Then today starts waiting too.
// Time: O(n), every day is pushed once and popped at most once.
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures)) // 0 = no warmer day
	waiting := []int{}                    // indexes of days waiting for a warmer day

	for today, temp := range temperatures {
		// today is warmer than the day on top of the stack
		for len(waiting) > 0 && temperatures[waiting[len(waiting)-1]] < temp {
			day := waiting[len(waiting)-1]
			waiting = waiting[:len(waiting)-1]
			res[day] = today - day
		}
		waiting = append(waiting, today)
	}
	return res
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem.
	check("example 1", `dailyTemperatures([]int{30, 38, 30, 36, 35, 40, 28})`,
		dailyTemperatures([]int{30, 38, 30, 36, 35, 40, 28}), []int{1, 4, 1, 2, 1, 0, 0})

	// Example 2 from the problem: it only gets colder.
	check("example 2: getting colder", `dailyTemperatures([]int{22, 21, 20})`,
		dailyTemperatures([]int{22, 21, 20}), []int{0, 0, 0})

	// The classic example.
	check("classic", `dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})`,
		dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}), []int{1, 1, 4, 2, 1, 1, 0, 0})

	// Getting warmer every day.
	check("getting warmer", `dailyTemperatures([]int{30, 40, 50, 60})`,
		dailyTemperatures([]int{30, 40, 50, 60}), []int{1, 1, 1, 0})

	// One day.
	check("one day", `dailyTemperatures([]int{50})`,
		dailyTemperatures([]int{50}), []int{0})

	// The same temperature is NOT warmer.
	check("all the same", `dailyTemperatures([]int{70, 70, 70})`,
		dailyTemperatures([]int{70, 70, 70}), []int{0, 0, 0})

	// A long wait for the first day.
	check("long wait", `dailyTemperatures([]int{90, 30, 30, 91})`,
		dailyTemperatures([]int{90, 30, 30, 91}), []int{3, 2, 1, 0})

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
