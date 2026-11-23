// Topological sort (BFS): keep taking courses that have no prerequisites left.
func canFinish(numCourses int, prerequisites [][]int) bool {
	next := make([][]int, numCourses) // next[a] = courses unlocked by a
	need := make([]int, numCourses)   // need[b] = prerequisites b still needs
	for _, p := range prerequisites {
		next[p[1]] = append(next[p[1]], p[0])
		need[p[0]]++
	}
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if need[i] == 0 {
			queue = append(queue, i)
		}
	}
	taken := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		taken++
		for _, n := range next[course] {
			need[n]--
			if need[n] == 0 {
				queue = append(queue, n)
			}
		}
	}
	return taken == numCourses
}
