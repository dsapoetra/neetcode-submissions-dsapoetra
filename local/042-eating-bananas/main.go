// Problem:    Koko Eating Bananas
// Difficulty: Medium
// Link:       https://neetcode.io/problems/eating-bananas
// Source:     Data Structures & Algorithms/eating-bananas/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// There are piles of bananas and h hours. With eating speed k, each hour
// you pick one pile and eat k bananas from it (if the pile has fewer, you
// finish it and wait for the next hour). Return the smallest k that lets
// you finish all piles within h hours.
//
// Example 1:
//
//	Input:  piles = [1, 4, 3, 2], h = 9
//	Output: 2
//
// Example 2:
//
//	Input:  piles = [25, 10, 23, 4], h = 4
//	Output: 25
//
// How to run (from the "local" folder):
//
//	go run ./042-eating-bananas
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: we are looking for the smallest speed k that still finishes in h
// hours. A higher speed never needs more hours, so the speeds look like:
//
//	too slow, too slow, ..., OK, OK, OK
//
// That is a sorted yes/no list, so we can binary search the first "OK".
// The slowest useful speed is 1, the fastest is the biggest pile.
// Time: O(n * log(biggest pile)).
func minEatingSpeed(piles []int, h int) int {
	// the range of speeds to search: 1 .. biggest pile
	left := 1
	right := 0
	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}

	for left < right {
		// try the speed in the middle
		mid := (left + right) / 2
		if hoursNeeded(piles, mid) <= h {
			right = mid // fast enough: maybe an even slower speed works
		} else {
			left = mid + 1 // too slow
		}
	}
	// left == right: the slowest speed that still works
	return left
}

// hoursNeeded returns how many hours it takes to eat all piles at this speed.
func hoursNeeded(piles []int, speed int) int {
	hours := 0
	for _, pile := range piles {
		// pile / speed, rounded UP (a started hour counts as a full hour)
		hours += (pile + speed - 1) / speed
	}
	return hours
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: at speed 2 the piles take 1+2+2+1 = 6 hours.
	check("example 1", `minEatingSpeed([]int{1, 4, 3, 2}, 9)`,
		minEatingSpeed([]int{1, 4, 3, 2}, 9), 2)

	// Example 2 from the problem: 4 piles in 4 hours, so one pile per hour.
	check("example 2", `minEatingSpeed([]int{25, 10, 23, 4}, 4)`,
		minEatingSpeed([]int{25, 10, 23, 4}, 4), 25)

	// The classic examples.
	check("classic 1", `minEatingSpeed([]int{3, 6, 7, 11}, 8)`,
		minEatingSpeed([]int{3, 6, 7, 11}, 8), 4)
	check("classic 2", `minEatingSpeed([]int{30, 11, 23, 4, 20}, 5)`,
		minEatingSpeed([]int{30, 11, 23, 4, 20}, 5), 30)
	check("classic 3", `minEatingSpeed([]int{30, 11, 23, 4, 20}, 6)`,
		minEatingSpeed([]int{30, 11, 23, 4, 20}, 6), 23)

	// Lots of time: speed 1 is enough.
	check("plenty of time", `minEatingSpeed([]int{5}, 5)`,
		minEatingSpeed([]int{5}, 5), 1)

	// One hour for one pile: eat it all at once.
	check("one hour", `minEatingSpeed([]int{5}, 1)`,
		minEatingSpeed([]int{5}, 1), 5)

	// 10 bananas in 3 hours: speed 3 needs 4 hours, speed 4 needs 3.
	check("rounding up", `minEatingSpeed([]int{10}, 3)`,
		minEatingSpeed([]int{10}, 3), 4)

	// A huge pile.
	check("huge pile", `minEatingSpeed([]int{1000000000}, 2)`,
		minEatingSpeed([]int{1000000000}, 2), 500000000)

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
