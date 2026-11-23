// Problem: Meeting Rooms II
// Link:    https://neetcode.io/problems/meeting-schedule-ii
// Source:  Data Structures & Algorithms/meeting-schedule-ii/submission-0.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./meeting-schedule-ii
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

// Idea: put all start times in one sorted list and all end times in another.
// Walk through the start times in order:
//   - if the next meeting starts before the earliest meeting ends,
//     we need one more room
//   - otherwise a meeting has ended, so a room becomes free
//
// The answer is the biggest number of rooms used at the same time.
// Time: O(n log n).
func minMeetingRooms(intervals []Interval) int {
	starts := []int{}
	ends := []int{}
	for _, meeting := range intervals {
		starts = append(starts, meeting.start)
		ends = append(ends, meeting.end)
	}
	sort.Ints(starts)
	sort.Ints(ends)

	rooms := 0 // rooms in use right now
	best := 0  // most rooms ever in use
	s, e := 0, 0

	for s < len(starts) {
		if starts[s] < ends[e] {
			// a meeting starts while others are still running
			rooms++
			s++
			if rooms > best {
				best = rooms
			}
		} else {
			// the earliest meeting ended, free its room
			rooms--
			e++
		}
	}
	return best
}

// ===================== Test cases =====================

func main() {
	// Example 1 from the problem: (0,40) overlaps the other two, which do not overlap each other.
	check("example 1", minMeetingRooms([]Interval{{0, 40}, {5, 10}, {15, 20}}), 2)

	// Example 2 from the problem: one meeting.
	check("example 2: one meeting", minMeetingRooms([]Interval{{4, 9}}), 1)

	// No meetings: no rooms.
	check("no meetings", minMeetingRooms([]Interval{}), 0)

	// One ends exactly when the next starts: the room can be reused.
	check("back to back", minMeetingRooms([]Interval{{0, 10}, {10, 20}}), 1)

	// Three meetings all overlap at the same moment.
	check("three overlap", minMeetingRooms([]Interval{{1, 5}, {2, 6}, {3, 7}}), 3)

	// No overlap at all.
	check("no overlap", minMeetingRooms([]Interval{{1, 2}, {3, 4}, {5, 6}}), 1)

	// One long meeting and three short ones inside it.
	check("short ones inside a long one", minMeetingRooms([]Interval{{1, 10}, {2, 3}, {4, 5}, {6, 7}}), 2)

	// Two meetings at exactly the same time.
	check("same time", minMeetingRooms([]Interval{{1, 5}, {1, 5}}), 2)

	// The meetings are not sorted.
	check("unsorted", minMeetingRooms([]Interval{{9, 10}, {4, 9}, {4, 17}}), 2)

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
