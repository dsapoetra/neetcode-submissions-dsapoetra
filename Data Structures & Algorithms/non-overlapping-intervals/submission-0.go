// Idea (greedy): sort by END time. Always keep the interval that ends
// earliest, because it leaves the most space for the others.
// Any interval that starts before the last kept one ends must be removed.
// Time: O(n log n).
func eraseOverlapIntervals(intervals [][]int) int {
	// sort by end time
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][1] < intervals[b][1]
	})

	removed := 0
	lastEnd := intervals[0][1] // end of the last interval we kept

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < lastEnd {
			// overlaps the one we kept -> remove this one
			removed++
		} else {
			// no overlap -> keep it
			lastEnd = intervals[i][1]
		}
	}
	return removed
}
