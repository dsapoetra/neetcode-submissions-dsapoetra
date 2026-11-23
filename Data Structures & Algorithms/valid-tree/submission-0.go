// Idea: a graph with n nodes is a tree when:
//   1) it has exactly n-1 edges, and
//   2) every node can be reached from node 0 (it is connected).
// With exactly n-1 edges, a connected graph cannot have a cycle.
// Time: O(nodes + edges).
func validTree(n int, edges [][]int) bool {
	// rule 1
	if len(edges) != n-1 {
		return false
	}

	// build the neighbor list of every node
	neighbors := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		neighbors[a] = append(neighbors[a], b)
		neighbors[b] = append(neighbors[b], a)
	}

	// rule 2: visit everything reachable from node 0
	seen := make([]bool, n)
	visit(0, neighbors, seen)

	for i := 0; i < n; i++ {
		if !seen[i] {
			return false // a node was not reached
		}
	}
	return true
}

// visit marks node and everything connected to it as seen (DFS).
func visit(node int, neighbors [][]int, seen []bool) {
	seen[node] = true
	for _, next := range neighbors[node] {
		if !seen[next] {
			visit(next, neighbors, seen)
		}
	}
}
