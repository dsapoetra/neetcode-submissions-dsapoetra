// Problem: Meeting Rooms
// Link:    https://neetcode.io/problems/meeting-schedule
// Source:  Data Structures & Algorithms/meeting-schedule/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./meeting-schedule
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
	check("example 1: clash", canAttendMeetings([]Interval{{0, 30}, {5, 10}, {15, 20}}), false)

	// Example 2 from the problem: one after the other.
	check("example 2: no clash", canAttendMeetings([]Interval{{5, 8}, {9, 15}}), true)

	// No meetings: nothing can clash.
	check("no meetings", canAttendMeetings([]Interval{}), true)

	// One meeting.
	check("one meeting", canAttendMeetings([]Interval{{1, 2}}), true)

	// One ends exactly when the next starts: that is allowed.
	check("back to back", canAttendMeetings([]Interval{{0, 8}, {8, 10}}), true)

	// Meetings given in the wrong order.
	check("unsorted, no clash", canAttendMeetings([]Interval{{9, 15}, {5, 8}}), true)

	// Two meetings at exactly the same time.
	check("same time", canAttendMeetings([]Interval{{1, 5}, {1, 5}}), false)

	// A short meeting inside a long one.
	check("one inside another", canAttendMeetings([]Interval{{1, 10}, {2, 3}}), false)

	// Only the first two clash.
	check("first two clash", canAttendMeetings([]Interval{{1, 3}, {2, 4}, {5, 6}}), false)

	// Unsorted and clashing.
	check("unsorted, clash", canAttendMeetings([]Interval{{7, 10}, {2, 8}}), false)

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
