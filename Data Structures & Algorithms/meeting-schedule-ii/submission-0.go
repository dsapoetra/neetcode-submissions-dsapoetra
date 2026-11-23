// Idea: put all start times in one sorted list and all end times in another.
// Walk through the start times in order:
//   - if the next meeting starts before the earliest meeting ends,
//     we need one more room
//   - otherwise a meeting has ended, so a room becomes free
// The answer is the biggest number of rooms used at the same time.
// Time: O(n log n).
func minMeetingRooms(intervals []Interval) int {
	starts := []int{}
	ends := []int{}
	for _, meeting := range intervals {
		starts = append(starts, meeting.start)
		ends = append(ends, meeting.end)
	}
	sort.Ints(starts)
	sort.Ints(ends)

	rooms := 0 // rooms in use right now
	best := 0  // most rooms ever in use
	s, e := 0, 0

	for s < len(starts) {
		if starts[s] < ends[e] {
			// a meeting starts while others are still running
			rooms++
			s++
			if rooms > best {
				best = rooms
			}
		} else {
			// the earliest meeting ended, free its room
			rooms--
			e++
		}
	}
	return best
}
