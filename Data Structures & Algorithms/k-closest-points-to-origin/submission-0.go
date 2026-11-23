// Idea: sort the points by their distance to (0, 0) and take the first k.
// We compare x*x + y*y instead of the real distance: the square root does
// not change which one is smaller, so we can skip it.
// Time: O(n log n).
func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(a, b int) bool {
		distA := points[a][0]*points[a][0] + points[a][1]*points[a][1]
		distB := points[b][0]*points[b][0] + points[b][1]*points[b][1]
		return distA < distB
	})
	return points[:k]
}
