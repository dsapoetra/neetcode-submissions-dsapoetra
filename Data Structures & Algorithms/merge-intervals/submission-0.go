// Sort by start time, then merge each interval into the last one if they overlap.
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		last := res[len(res)-1]
		if cur[0] <= last[1] {
			// overlap: extend the last interval
			if cur[1] > last[1] {
				last[1] = cur[1]
			}
		} else {
			res = append(res, cur)
		}
	}
	return res
}
