// Idea: go through every node. If we have not seen it yet, it belongs
// to a new component: count it, then use DFS to mark every node in
// that component as seen.
// Time: O(nodes + edges).
func countComponents(n int, edges [][]int) int {
	// build the neighbor list of every node
	neighbors := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		neighbors[a] = append(neighbors[a], b)
		neighbors[b] = append(neighbors[b], a)
	}

	seen := make([]bool, n)
	count := 0
	for node := 0; node < n; node++ {
		if !seen[node] {
			count++ // found a new component
			visit(node, neighbors, seen)
		}
	}
	return count
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
