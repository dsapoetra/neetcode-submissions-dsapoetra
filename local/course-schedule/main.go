// Problem: Course Schedule
// Link:    https://neetcode.io/problems/course-schedule
// Source:  Data Structures & Algorithms/course-schedule/submission-1.go (the solution below is a copy of it)
//
// How to run (from the "local" folder):
//
//	go run ./course-schedule
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea (topological sort with BFS): a course can be taken when it has
// no prerequisites left. Take those courses first, which may unlock
// other courses. If we manage to take every course there is no cycle.
// Time: O(courses + prerequisites).
func canFinish(numCourses int, prerequisites [][]int) bool {
	unlocks := make([][]int, numCourses) // unlocks[a] = courses that need a first
	need := make([]int, numCourses)      // need[b] = how many prerequisites b still needs

	// pair [a, b] means: take b before a
	for _, p := range prerequisites {
		course, pre := p[0], p[1]
		unlocks[pre] = append(unlocks[pre], course)
		need[course]++
	}

	// start with every course that needs nothing
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if need[i] == 0 {
			queue = append(queue, i)
		}
	}

	taken := 0
	for len(queue) > 0 {
		// take one course
		course := queue[0]
		queue = queue[1:]
		taken++

		// every course that needed it now needs one less
		for _, next := range unlocks[course] {
			need[next]--
			if need[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// if a cycle exists, some courses could never be taken
	return taken == numCourses
}

// ===================== Test cases =====================

func main() {
	// A pair [a, b] means: you must take course b before course a.

	// Example 1 from the problem: take 1, then 0.
	check("example 1: one prerequisite", `canFinish(2, [][]int{{0, 1}})`,
		canFinish(2, [][]int{{0, 1}}), true)

	// Example 2 from the problem: 0 needs 1 and 1 needs 0.
	check("example 2: two courses need each other", `canFinish(2, [][]int{{0, 1}, {1, 0}})`,
		canFinish(2, [][]int{{0, 1}, {1, 0}}), false)

	// One course without prerequisites.
	check("one course", `canFinish(1, [][]int{})`,
		canFinish(1, [][]int{}), true)

	// Many courses, no prerequisites.
	check("no prerequisites", `canFinish(5, [][]int{})`,
		canFinish(5, [][]int{}), true)

	// A chain: 0 -> 1 -> 2 -> 3.
	check("chain", `canFinish(4, [][]int{{1, 0}, {2, 1}, {3, 2}})`,
		canFinish(4, [][]int{{1, 0}, {2, 1}, {3, 2}}), true)

	// A cycle of three courses.
	check("cycle of three", `canFinish(3, [][]int{{0, 1}, {1, 2}, {2, 0}})`,
		canFinish(3, [][]int{{0, 1}, {1, 2}, {2, 0}}), false)

	// A diamond: 3 needs 1 and 2, both need 0. No cycle.
	check("diamond", `canFinish(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})`,
		canFinish(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}), true)

	// A course that needs itself.
	check("course needs itself", `canFinish(1, [][]int{{0, 0}})`,
		canFinish(1, [][]int{{0, 0}}), false)

	// The cycle (1 <-> 2) is only in one part; courses 3 and 4 are fine.
	check("cycle in one part only", `canFinish(5, [][]int{{1, 0}, {2, 1}, {1, 2}, {4, 3}})`,
		canFinish(5, [][]int{{1, 0}, {2, 1}, {1, 2}, {4, 3}}), false)

	// Two separate chains.
	check("two separate chains", `canFinish(4, [][]int{{1, 0}, {3, 2}})`,
		canFinish(4, [][]int{{1, 0}, {3, 2}}), true)

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
