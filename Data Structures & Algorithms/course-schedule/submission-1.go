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
