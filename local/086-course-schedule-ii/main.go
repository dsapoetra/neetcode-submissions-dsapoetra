// Problem:    Course Schedule II
// Difficulty: Medium
// Link:       https://neetcode.io/problems/course-schedule-ii
// Source:     Data Structures & Algorithms/course-schedule-ii/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// There are numCourses courses, numbered from 0. A pair [a, b] means: you
// must take course b before course a. Return an order in which all courses
// can be taken (any valid order is correct), or an empty list if it is
// impossible.
//
// Example 1:
//
//	Input:  numCourses = 3, prerequisites = [[1, 0]]
//	Output: [0, 1, 2]     (other valid orders are correct too)
//
// Example 2:
//
//	Input:  numCourses = 3, prerequisites = [[0, 1], [1, 2], [2, 0]]
//	Output: []
//
// How to run (from the "local" folder):
//
//	go run ./086-course-schedule-ii
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

// Idea: a course can be taken once all its prerequisites are done.
// For every course count how many prerequisites are still missing.
//  1. courses with 0 missing prerequisites can be taken now -> queue
//  2. take a course from the queue; every course that needed it now has
//     one missing prerequisite less. If that count reaches 0, it can be
//     taken too -> queue
//
// If we manage to take all courses, the order we took them is the answer.
// If not, there is a cycle and no order exists.
// Time: O(courses + prerequisites).
func findOrder(numCourses int, prerequisites [][]int) []int {
	unlocks := make([][]int, numCourses) // unlocks[b] = courses that need b first
	missing := make([]int, numCourses)   // missing[a] = prerequisites of a not taken yet

	for _, pair := range prerequisites {
		course, needed := pair[0], pair[1]
		unlocks[needed] = append(unlocks[needed], course)
		missing[course]++
	}

	// 1) courses we can take right away
	queue := []int{}
	for course := 0; course < numCourses; course++ {
		if missing[course] == 0 {
			queue = append(queue, course)
		}
	}

	// 2) take courses one by one
	order := []int{}
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		order = append(order, course)

		for _, next := range unlocks[course] {
			missing[next]--
			if missing[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	if len(order) != numCourses {
		return []int{} // a cycle: some courses could never be taken
	}
	return order
}

// ===================== Test cases =====================

func main() {
	// A pair [a, b] means: you must take course b before course a.

	// validOrder is for cases with SEVERAL correct answers. It checks that
	// the answer contains every course once and respects every pair.
	validOrder := func(numCourses int, prerequisites [][]int) bool {
		order := findOrder(numCourses, prerequisites)
		position := map[int]int{}
		for i, course := range order {
			position[course] = i
		}
		if len(order) != numCourses || len(position) != numCourses {
			return false
		}
		for _, pair := range prerequisites {
			if position[pair[1]] > position[pair[0]] {
				return false
			}
		}
		return true
	}

	// Example 1 from the problem: 0 before 1; course 2 can go anywhere.
	check("example 1: valid order", `validOrder(3, [][]int{{1, 0}})`,
		validOrder(3, [][]int{{1, 0}}), true)

	// Example 2 from the problem: 0 needs 1, 1 needs 2, 2 needs 0. Impossible.
	check("example 2: cycle", `findOrder(3, [][]int{{0, 1}, {1, 2}, {2, 0}})`,
		findOrder(3, [][]int{{0, 1}, {1, 2}, {2, 0}}), []int{})

	// A chain has exactly one answer.
	check("chain", `findOrder(4, [][]int{{1, 0}, {2, 1}, {3, 2}})`,
		findOrder(4, [][]int{{1, 0}, {2, 1}, {3, 2}}), []int{0, 1, 2, 3})

	// Two courses.
	check("0 before 1", `findOrder(2, [][]int{{1, 0}})`,
		findOrder(2, [][]int{{1, 0}}), []int{0, 1})
	check("1 before 0", `findOrder(2, [][]int{{0, 1}})`,
		findOrder(2, [][]int{{0, 1}}), []int{1, 0})

	// Two courses that need each other.
	check("two courses need each other", `findOrder(2, [][]int{{0, 1}, {1, 0}})`,
		findOrder(2, [][]int{{0, 1}, {1, 0}}), []int{})

	// One course.
	check("one course", `findOrder(1, [][]int{})`,
		findOrder(1, [][]int{}), []int{0})

	// The classic diamond: 0 first, then 1 and 2 in any order, then 3.
	check("diamond: valid order", `validOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})`,
		validOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}), true)

	// No prerequisites: any order is fine.
	check("no prerequisites: valid order", `validOrder(3, [][]int{})`,
		validOrder(3, [][]int{}), true)

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
