// Idea: look at the cars from the one closest to the target to the one
// farthest away. For each car compute the time it needs to reach the target
// if nothing is in its way.
//   - if it needs MORE time than the fleet in front of it, it can never
//     catch up, so it starts a new fleet
//   - otherwise it catches the fleet in front and joins it (it then
//     arrives at the same time as that fleet)
// Time: O(n log n) because of the sort.
func carFleet(target int, position []int, speed []int) int {
	// order[i] = index of a car; sorted so the car closest to the target is first
	order := make([]int, len(position))
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(a, b int) bool {
		return position[order[a]] > position[order[b]]
	})

	fleets := 0
	slowestTime := 0.0 // arrival time of the fleet right in front

	for _, car := range order {
		time := float64(target-position[car]) / float64(speed[car])
		if time > slowestTime {
			// cannot catch the fleet in front: this car leads a new fleet
			fleets++
			slowestTime = time
		}
		// otherwise it joins the fleet in front, nothing changes
	}
	return fleets
}
