// Every node must be inside a (low, high) range that gets tighter going down.
func isValidBST(root *TreeNode) bool {
	return checkBST(root, nil, nil)
}

func checkBST(node *TreeNode, low *TreeNode, high *TreeNode) bool {
	if node == nil {
		return true
	}
	if low != nil && node.Val <= low.Val {
		return false
	}
	if high != nil && node.Val >= high.Val {
		return false
	}
	return checkBST(node.Left, low, node) && checkBST(node.Right, node, high)
}
