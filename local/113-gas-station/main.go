// Problem:    Gas Station
// Difficulty: Medium
// Link:       https://neetcode.io/problems/gas-station
// Source:     Data Structures & Algorithms/gas-station/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Gas stations stand on a circular road. At station i you get gas[i], and
// driving from station i to the next one costs cost[i]. You start with an
// empty tank. Return the index of the station from which you can drive
// once around the whole circle, or -1 if there is none. If an answer
// exists, there is only one.
//
// Example 1:
//
//	Input:  gas = [1, 2, 3, 4], cost = [2, 2, 4, 1]
//	Output: 3
//
// Example 2:
//
//	Input:  gas = [1, 2, 3], cost = [2, 3, 2]
//	Output: -1
//
// How to run (from the "local" folder):
//
//	go run ./113-gas-station
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (greedy):
//  1. if the total gas is less than the total cost, no start can work
//  2. otherwise an answer exists. Drive from station 0 and track the
//     tank. If the tank goes below 0 at station i, then NO station from
//     our start up to i can be the start (they would all arrive here
//     with even less gas). So try again with start = i + 1.
//
// Time: O(n).
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	totalCost := 0
	for i := range gas {
		totalGas += gas[i]
		totalCost += cost[i]
	}
	if totalGas < totalCost {
		return -1
	}

	start := 0
	tank := 0
	for i := range gas {
		tank += gas[i] - cost[i]
		if tank < 0 {
			// cannot get past station i: start fresh after it
			start = i + 1
			tank = 0
		}
	}
	return start
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: only station 3 works as the start.
	check("example 1", `canCompleteCircuit([]int{1, 2, 3, 4}, []int{2, 2, 4, 1})`,
		canCompleteCircuit([]int{1, 2, 3, 4}, []int{2, 2, 4, 1}), 3)

	// Example 2 from the problem: 6 gas in total, but the trip costs 7.
	check("example 2: not enough gas", `canCompleteCircuit([]int{1, 2, 3}, []int{2, 3, 2})`,
		canCompleteCircuit([]int{1, 2, 3}, []int{2, 3, 2}), -1)

	// The classic examples.
	check("classic", `canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})`,
		canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}), 3)
	check("classic impossible", `canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})`,
		canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}), -1)

	// One station.
	check("one station, enough", `canCompleteCircuit([]int{5}, []int{4})`,
		canCompleteCircuit([]int{5}, []int{4}), 0)
	check("one station, not enough", `canCompleteCircuit([]int{4}, []int{5})`,
		canCompleteCircuit([]int{4}, []int{5}), -1)
	check("one station, exactly enough", `canCompleteCircuit([]int{2}, []int{2})`,
		canCompleteCircuit([]int{2}, []int{2}), 0)

	// Start at the first station.
	check("start at 0", `canCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2})`,
		canCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2}), 0)

	// Start at the last station.
	check("start at the last", `canCompleteCircuit([]int{1, 1, 3}, []int{2, 2, 1})`,
		canCompleteCircuit([]int{1, 1, 3}, []int{2, 2, 1}), 2)

	// The tank goes negative twice before the right start is found.
	check("two restarts", `canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1})`,
		canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}), 4)

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
