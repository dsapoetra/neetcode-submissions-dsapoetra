// Idea (union-find): start with every node in its own group. Add the
// edges one by one; an edge joins the groups of its two nodes.
// If the two nodes are ALREADY in the same group, they are connected
// without this edge, so this edge closes a cycle: it is the answer.
// parent[x] points to another node of the same group; following the
// parents leads to the "boss" of the group.
func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1) // nodes are numbered from 1
	for i := range parent {
		parent[i] = i // every node is its own boss
	}

	// find returns the boss of x's group.
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) // shortcut: point straight to the boss
		}
		return parent[x]
	}

	for _, edge := range edges {
		bossA := find(edge[0])
		bossB := find(edge[1])

		if bossA == bossB {
			return edge // already connected: this edge is the extra one
		}
		parent[bossA] = bossB // join the two groups
	}
	return []int{}
}
