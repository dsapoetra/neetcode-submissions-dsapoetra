// Idea (greedy): the SMALLEST card left must be the start of a group,
// because nothing smaller can come before it.
// So sort the cards and count how many of each value we have. Go through
// the cards from small to big: if a card is still available, it starts a
// group, and we must also take value+1, value+2, ... If one of them is
// missing, it is impossible.
// Time: O(n log n + n * groupSize).
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false // the cards cannot be split evenly
	}

	sort.Ints(hand)
	count := map[int]int{}
	for _, card := range hand {
		count[card]++
	}

	for _, card := range hand {
		if count[card] == 0 {
			continue // already used in an earlier group
		}

		// card is the smallest card left: build card, card+1, ...
		for value := card; value < card+groupSize; value++ {
			if count[value] == 0 {
				return false // a card is missing
			}
			count[value]--
		}
	}
	return true
}
