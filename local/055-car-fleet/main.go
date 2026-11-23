// Problem:    Car Fleet
// Difficulty: Medium
// Link:       https://neetcode.io/problems/car-fleet
// Source:     Data Structures & Algorithms/car-fleet/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Cars drive on a one-lane road towards the same target. Car i starts at
// position[i] and drives at speed[i]. A car cannot overtake: when it catches
// up with a slower car, it slows down and they continue together as one
// "fleet". Return how many fleets arrive at the target. (A single car is a
// fleet too; catching up exactly at the target still counts as joining.)
//
// Example 1:
//
//	Input:  target = 10, position = [1, 4], speed = [3, 2]
//	Output: 1
//
// Example 2:
//
//	Input:  target = 10, position = [4, 1, 0, 7], speed = [2, 2, 1, 1]
//	Output: 3
//
// How to run (from the "local" folder):
//
//	go run ./055-car-fleet
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"sort"
)

// ===================== Solution =====================

// Idea: look at the cars from the one closest to the target to the one
// farthest away. For each car compute the time it needs to reach the target
// if nothing is in its way.
//   - if it needs MORE time than the fleet in front of it, it can never
//     catch up, so it starts a new fleet
//   - otherwise it catches the fleet in front and joins it (it then
//     arrives at the same time as that fleet)
//
// Time: O(n log n) because of the sort.
func carFleet(target int, position []int, speed []int) int {
	// order[i] = index of a car; sorted so the car closest to the target is first
	order := make([]int, len(position))
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(a, b int) bool {
		return position[order[a]] > position[order[b]]
	})

	fleets := 0
	slowestTime := 0.0 // arrival time of the fleet right in front

	for _, car := range order {
		time := float64(target-position[car]) / float64(speed[car])
		if time > slowestTime {
			// cannot catch the fleet in front: this car leads a new fleet
			fleets++
			slowestTime = time
		}
		// otherwise it joins the fleet in front, nothing changes
	}
	return fleets
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: the car at 1 catches the car at 4 exactly at the target.
	check("example 1", `carFleet(10, []int{1, 4}, []int{3, 2})`,
		carFleet(10, []int{1, 4}, []int{3, 2}), 1)

	// Example 2 from the problem.
	check("example 2", `carFleet(10, []int{4, 1, 0, 7}, []int{2, 2, 1, 1})`,
		carFleet(10, []int{4, 1, 0, 7}, []int{2, 2, 1, 1}), 3)

	// The classic example.
	check("classic", `carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3})`,
		carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}), 3)

	// One car is one fleet.
	check("one car", `carFleet(10, []int{3}, []int{3})`,
		carFleet(10, []int{3}, []int{3}), 1)

	// Faster cars behind slower cars: they all bunch up.
	check("all bunch up", `carFleet(100, []int{0, 2, 4}, []int{4, 2, 1})`,
		carFleet(100, []int{0, 2, 4}, []int{4, 2, 1}), 1)

	// Same speed: the car behind never catches up.
	check("same speed", `carFleet(10, []int{0, 5}, []int{1, 1})`,
		carFleet(10, []int{0, 5}, []int{1, 1}), 2)

	// The car behind is faster, but not fast enough before the target.
	check("too slow to catch up", `carFleet(10, []int{6, 8}, []int{3, 2})`,
		carFleet(10, []int{6, 8}, []int{3, 2}), 2)

	// The slower car is BEHIND: nobody catches anybody.
	check("slow car behind", `carFleet(10, []int{0, 5}, []int{1, 3})`,
		carFleet(10, []int{0, 5}, []int{1, 3}), 2)

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
