// Idea: for every node, imagine it is the TOP of the path. Then the
// path is: best path going down-left + the node + best path going down-right.
// We try every node as the top and keep the biggest total.
//
// The helper returns the best path that starts at a node and goes DOWN
// on only one side, because the parent can only connect to one side.
// A side with a negative sum is ignored (we use 0 instead).
// Time: O(n).
func maxPathSum(root *TreeNode) int {
	best := root.Val

	var bestDown func(node *TreeNode) int
	bestDown = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// best path going down each side (0 if it would make things worse)
		left := bestDown(node.Left)
		if left < 0 {
			left = 0
		}
		right := bestDown(node.Right)
		if right < 0 {
			right = 0
		}

		// path with this node as the top, using both sides
		if node.Val+left+right > best {
			best = node.Val + left + right
		}

		// for the parent: this node plus only the better side
		if left > right {
			return node.Val + left
		}
		return node.Val + right
	}

	bestDown(root)
	return best
}
