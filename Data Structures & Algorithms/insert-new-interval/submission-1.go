// Idea: the intervals are already sorted, so do it in 3 steps:
// 1) copy every interval that ends before the new one starts
// 2) merge every interval that overlaps the new one into it
// 3) copy every interval that is left
// Time: O(n).
func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	i := 0

	// step 1: intervals completely before the new one
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	// step 2: intervals that overlap the new one.
	// grow newInterval so it covers them.
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	res = append(res, newInterval)

	// step 3: intervals completely after the new one
	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}
	return res
}
