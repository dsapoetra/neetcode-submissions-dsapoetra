// Idea: sort the meetings by start time. If any meeting starts before
// the previous one ends, the two clash and we cannot attend all.
// Time: O(n log n).
func canAttendMeetings(intervals []Interval) bool {
	// sort by start time
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a].start < intervals[b].start
	})

	// compare each meeting with the one right before it
	for i := 1; i < len(intervals); i++ {
		if intervals[i].start < intervals[i-1].end {
			return false // clash
		}
	}
	return true
}
