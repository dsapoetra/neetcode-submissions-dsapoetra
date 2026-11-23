// Idea (Prim's algorithm): grow one connected group, starting with point 0.
// In every step add the point that is CHEAPEST to connect to the group.
// cost[p] = the cheapest known distance from point p to any point that is
// already in the group. After adding a point, update the others.
// Time: O(n^2).
func minCostConnectPoints(points [][]int) int {
	const inf = 1000000000
	n := len(points)

	inGroup := make([]bool, n)
	cost := make([]int, n)
	for i := range cost {
		cost[i] = inf
	}
	cost[0] = 0 // start with point 0, it costs nothing

	total := 0
	for step := 0; step < n; step++ {
		// find the cheapest point that is not in the group yet
		best := -1
		for p := 0; p < n; p++ {
			if !inGroup[p] && (best == -1 || cost[p] < cost[best]) {
				best = p
			}
		}

		// add it
		inGroup[best] = true
		total += cost[best]

		// maybe the new point is a cheaper way to reach the others
		for p := 0; p < n; p++ {
			if !inGroup[p] {
				distance := absDiff(points[best][0], points[p][0]) + absDiff(points[best][1], points[p][1])
				if distance < cost[p] {
					cost[p] = distance
				}
			}
		}
	}
	return total
}

// absDiff returns the distance between two numbers (never negative).
func absDiff(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
