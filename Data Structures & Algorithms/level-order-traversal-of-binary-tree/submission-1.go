// Idea: BFS with a queue. The queue holds the nodes of the current level.
// For each level, take out exactly that many nodes and put their children in.
// Time: O(n).
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := []int{}
		size := len(queue) // number of nodes in this level

		for i := 0; i < size; i++ {
			// take the first node out of the queue
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			// its children belong to the next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
