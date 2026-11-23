// Sort by start time. If a meeting starts before the previous one ends, it clashes.
func canAttendMeetings(intervals []Interval) bool {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a].start < intervals[b].start
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i].start < intervals[i-1].end {
			return false
		}
	}
	return true
}
