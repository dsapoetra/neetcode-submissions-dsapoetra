// Idea: DFS through the graph and create a copy of each node.
// A map (old node -> its copy) remembers what we already copied,
// so a node is never copied twice and cycles do not loop forever.
// Time: O(nodes + edges).
func cloneGraph(node *Node) *Node {
	copies := map[*Node]*Node{}
	return cloneNode(node, copies)
}

// cloneNode returns the copy of node, creating it if needed.
func cloneNode(node *Node, copies map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	// already copied -> reuse it
	existing, found := copies[node]
	if found {
		return existing
	}

	// make the copy and save it in the map BEFORE visiting neighbors
	newNode := &Node{Val: node.Val}
	copies[node] = newNode

	// copy every neighbor and connect it
	for _, neighbor := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, cloneNode(neighbor, copies))
	}
	return newNode
}
