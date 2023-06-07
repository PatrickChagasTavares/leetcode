package binarytreemaxdepth

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depthLeft := recursion(root.Left, 1)
	depthRight := recursion(root.Right, 1)

	if depthLeft >= depthRight {
		return depthLeft
	}

	return depthRight
}

func recursion(tree *TreeNode, depth int) int {
	if tree == nil {
		return depth
	}
	depth++
	if tree.Left == nil && tree.Right == nil {
		return depth
	}

	var auxLeft int
	if tree.Left != nil {
		auxLeft = recursion(tree.Left, depth)
	}

	var auxRight int
	if tree.Right != nil {
		auxRight = recursion(tree.Right, depth)
	}
	if auxLeft >= auxRight {
		return auxLeft
	}

	return auxRight
}
