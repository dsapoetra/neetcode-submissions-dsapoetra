// Idea: a course can be taken once all its prerequisites are done.
// For every course count how many prerequisites are still missing.
//   1) courses with 0 missing prerequisites can be taken now -> queue
//   2) take a course from the queue; every course that needed it now has
//      one missing prerequisite less. If that count reaches 0, it can be
//      taken too -> queue
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
