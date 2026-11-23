// Idea: in a BST every node must be bigger than everything on its left
// and smaller than everything on its right. So each node has an allowed
// range (low, high). Going left lowers the upper limit, going right
// raises the lower limit.
// Time: O(n).
func isValidBST(root *TreeNode) bool {
	// the root has no limits (nil means no limit)
	return checkBST(root, nil, nil)
}

// checkBST returns true if every node in this subtree has a value
// strictly between low.Val and high.Val.
func checkBST(node *TreeNode, low *TreeNode, high *TreeNode) bool {
	if node == nil {
		return true
	}
	// must be bigger than the lower limit
	if low != nil && node.Val <= low.Val {
		return false
	}
	// must be smaller than the upper limit
	if high != nil && node.Val >= high.Val {
		return false
	}
	// left subtree: values must be smaller than this node
	// right subtree: values must be bigger than this node
	return checkBST(node.Left, low, node) && checkBST(node.Right, node, high)
}
