// Idea: sort the intervals by start time. Then walk through them:
// if the current interval overlaps the last one we kept, merge them,
// otherwise keep it as a new interval.
// Time: O(n log n) because of the sort.
func merge(intervals [][]int) [][]int {
	// sort by start time
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	// start the answer with the first interval
	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		last := res[len(res)-1] // the last interval in the answer

		if cur[0] <= last[1] {
			// cur starts before last ends -> overlap.
			// extend the end of last if cur ends later.
			if cur[1] > last[1] {
				last[1] = cur[1]
			}
		} else {
			// no overlap -> cur is a new separate interval
			res = append(res, cur)
		}
	}
	return res
}
