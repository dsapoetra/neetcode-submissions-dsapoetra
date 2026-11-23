// Idea: we are looking for the smallest speed k that still finishes in h
// hours. A higher speed never needs more hours, so the speeds look like:
//   too slow, too slow, ..., OK, OK, OK
// That is a sorted yes/no list, so we can binary search the first "OK".
// The slowest useful speed is 1, the fastest is the biggest pile.
// Time: O(n * log(biggest pile)).
func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := 0
	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}

	for left < right {
		mid := (left + right) / 2
		if hoursNeeded(piles, mid) <= h {
			right = mid // fast enough: maybe an even slower speed works
		} else {
			left = mid + 1 // too slow
		}
	}
	return left
}

// hoursNeeded returns how many hours it takes to eat all piles at this speed.
func hoursNeeded(piles []int, speed int) int {
	hours := 0
	for _, pile := range piles {
		// pile / speed, rounded UP (a started hour counts as a full hour)
		hours += (pile + speed - 1) / speed
	}
	return hours
}
