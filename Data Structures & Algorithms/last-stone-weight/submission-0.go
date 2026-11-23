// Idea: do exactly what the problem says. Sort the stones, take the two
// heaviest and smash them. If they are different, the rest goes back.
// Repeat until at most one stone is left.
// Time: O(n^2 log n) because we sort every round. That is fine here, there
// are only a few stones. (A max-heap would be faster, but needs more code.)
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)

		heaviest := stones[len(stones)-1]
		second := stones[len(stones)-2]
		stones = stones[:len(stones)-2] // remove both

		if heaviest != second {
			stones = append(stones, heaviest-second) // what is left of the heavier one
		}
	}

	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}
