// Idea (Bellman-Ford): dist[x] = the shortest time we know so far to
// reach node x. At the start only k is reachable (time 0).
// Then repeat: look at every edge u -> v. If going through u is faster
// than what we know for v, update v.
// A shortest path uses at most n-1 edges, so after n-1 rounds all the
// times are final. The signal is everywhere when the SLOWEST node got it.
// Time: O(n * edges). (Dijkstra with a heap is faster, but needs more code.)
func networkDelayTime(times [][]int, n int, k int) int {
	const inf = 1000000000

	dist := make([]int, n+1) // nodes are numbered from 1
	for i := range dist {
		dist[i] = inf
	}
	dist[k] = 0

	for round := 0; round < n-1; round++ {
		for _, edge := range times {
			from, to, time := edge[0], edge[1], edge[2]
			if dist[from] != inf && dist[from]+time < dist[to] {
				dist[to] = dist[from] + time
			}
		}
	}

	// the answer is the slowest node
	slowest := 0
	for node := 1; node <= n; node++ {
		if dist[node] == inf {
			return -1 // this node never gets the signal
		}
		if dist[node] > slowest {
			slowest = dist[node]
		}
	}
	return slowest
}
