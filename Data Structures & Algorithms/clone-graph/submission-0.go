// DFS. A map from old node to its copy stops us from copying a node twice.
func cloneGraph(node *Node) *Node {
	copies := map[*Node]*Node{}
	return cloneNode(node, copies)
}

func cloneNode(node *Node, copies map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if c, ok := copies[node]; ok {
		return c
	}
	c := &Node{Val: node.Val}
	copies[node] = c
	for _, nb := range node.Neighbors {
		c.Neighbors = append(c.Neighbors, cloneNode(nb, copies))
	}
	return c
}
