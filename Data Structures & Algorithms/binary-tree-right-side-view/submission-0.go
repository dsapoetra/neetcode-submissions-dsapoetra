// Idea: standing on the right side we see the LAST node of every level.
// So walk through the tree level by level (BFS with a queue) and keep the
// last value of each level.
// Time: O(n).
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue) // number of nodes in this level

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			// the last node of the level is the one we can see
			if i == size-1 {
				res = append(res, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}
