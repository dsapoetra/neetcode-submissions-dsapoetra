// Problem:    Meeting Rooms
// Difficulty: Easy
// Link:       https://neetcode.io/problems/meeting-schedule
// Source:     Data Structures & Algorithms/meeting-schedule/submission-1.go (the solution below is a copy of it)
//
// THE PROBLEM
// Given a list of meetings (start, end), return true if one person can
// attend all of them, that is: no two meetings overlap. A meeting that ends
// at time t and one that starts at time t do not overlap.
//
// Example 1:
//
//	Input:  intervals = [(0, 30), (5, 10), (15, 20)]
//	Output: false
//
// Example 2:
//
//	Input:  intervals = [(5, 8), (9, 15)]
//	Output: true
//
// How to run (from the "local" folder):
//
//	go run ./021-meeting-schedule
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
	"sort"
)

// Interval is one meeting with a start and an end time.
// NeetCode gives us this type; locally we have to define it ourselves.
type Interval struct {
	start int
	end   int
}

// ===================== Solution =====================

// Idea: sort the meetings by start time. If any meeting starts before
// the previous one ends, the two clash and we cannot attend all.
// Time: O(n log n).
func canAttendMeetings(intervals []Interval) bool {
	// sort by start time
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a].start < intervals[b].start
	})

	// compare each meeting with the one right before it
	for i := 1; i < len(intervals); i++ {
		if intervals[i].start < intervals[i-1].end {
			return false // clash
		}
	}
	return true
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: (0,30) clashes with the other two.
	check("example 1: clash", `canAttendMeetings([]Interval{{0, 30}, {5, 10}, {15, 20}})`,
		canAttendMeetings([]Interval{{0, 30}, {5, 10}, {15, 20}}), false)

	// Example 2 from the problem: one after the other.
	check("example 2: no clash", `canAttendMeetings([]Interval{{5, 8}, {9, 15}})`,
		canAttendMeetings([]Interval{{5, 8}, {9, 15}}), true)

	// No meetings: nothing can clash.
	check("no meetings", `canAttendMeetings([]Interval{})`,
		canAttendMeetings([]Interval{}), true)

	// One meeting.
	check("one meeting", `canAttendMeetings([]Interval{{1, 2}})`,
		canAttendMeetings([]Interval{{1, 2}}), true)

	// One ends exactly when the next starts: that is allowed.
	check("back to back", `canAttendMeetings([]Interval{{0, 8}, {8, 10}})`,
		canAttendMeetings([]Interval{{0, 8}, {8, 10}}), true)

	// Meetings given in the wrong order.
	check("unsorted, no clash", `canAttendMeetings([]Interval{{9, 15}, {5, 8}})`,
		canAttendMeetings([]Interval{{9, 15}, {5, 8}}), true)

	// Two meetings at exactly the same time.
	check("same time", `canAttendMeetings([]Interval{{1, 5}, {1, 5}})`,
		canAttendMeetings([]Interval{{1, 5}, {1, 5}}), false)

	// A short meeting inside a long one.
	check("one inside another", `canAttendMeetings([]Interval{{1, 10}, {2, 3}})`,
		canAttendMeetings([]Interval{{1, 10}, {2, 3}}), false)

	// Only the first two clash.
	check("first two clash", `canAttendMeetings([]Interval{{1, 3}, {2, 4}, {5, 6}})`,
		canAttendMeetings([]Interval{{1, 3}, {2, 4}, {5, 6}}), false)

	// Unsorted and clashing.
	check("unsorted, clash", `canAttendMeetings([]Interval{{7, 10}, {2, 8}})`,
		canAttendMeetings([]Interval{{7, 10}, {2, 8}}), false)

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
