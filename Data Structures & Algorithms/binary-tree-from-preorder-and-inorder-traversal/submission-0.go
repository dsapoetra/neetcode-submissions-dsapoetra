// Idea:
//   - preorder is: root, left subtree, right subtree
//     -> the FIRST value of preorder is the root
//   - inorder is:  left subtree, root, right subtree
//     -> everything before the root in inorder is the left subtree,
//        everything after it is the right subtree
// So: take the root, find it in inorder, split both lists, and build
// the left and right subtrees the same way (recursion).
// Time: O(n^2) because we search for the root in inorder each time.
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// find the root in inorder
	mid := 0
	for inorder[mid] != rootVal {
		mid++
	}
	// mid is also the number of nodes in the left subtree

	root.Left = buildTree(preorder[1:1+mid], inorder[:mid])
	root.Right = buildTree(preorder[1+mid:], inorder[mid+1:])
	return root
}
