// Idea: the SHORTEST interval should answer as many queries as it can,
// then the second shortest answers the queries that are still open, ...
//   1) sort the intervals by length, shortest first
//   2) sort the queries by value (we remember where each one came from)
//   3) for each interval, find the first sorted query that is >= left
//      (binary search) and walk right while the query is <= right.
//      Every open query we meet gets this interval's length.
// To not walk over answered queries again and again, skip[i] points to
// the next query at or after i that might still be open. That way every
// query is answered once and then jumped over.
// Time: O(n log n + q log q).
func minInterval(intervals [][]int, queries []int) []int {
	// 1) shortest interval first
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][1]-intervals[a][0] < intervals[b][1]-intervals[b][0]
	})

	// 2) order[i] = index (in queries) of the i-th smallest query
	q := len(queries)
	order := make([]int, q)
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(a, b int) bool {
		return queries[order[a]] < queries[order[b]]
	})

	res := make([]int, q)
	for i := range res {
		res[i] = -1 // -1 = no interval contains this query
	}

	// skip[i] = i means sorted query i is still open.
	// Otherwise it points to a later position to continue from.
	skip := make([]int, q+1)
	for i := range skip {
		skip[i] = i
	}
	// nextOpen returns the first open sorted query at or after i (or q if none).
	var nextOpen func(i int) int
	nextOpen = func(i int) int {
		if skip[i] != i {
			skip[i] = nextOpen(skip[i]) // shortcut for the next time
		}
		return skip[i]
	}

	// 3) hand out the answers, shortest interval first
	for _, interval := range intervals {
		left, right := interval[0], interval[1]

		// binary search: first sorted query with value >= left
		low, high := 0, q
		for low < high {
			mid := (low + high) / 2
			if queries[order[mid]] < left {
				low = mid + 1
			} else {
				high = mid
			}
		}

		for i := nextOpen(low); i < q && queries[order[i]] <= right; i = nextOpen(i) {
			res[order[i]] = right - left + 1
			skip[i] = i + 1 // answered: jump over it from now on
		}
	}
	return res
}
