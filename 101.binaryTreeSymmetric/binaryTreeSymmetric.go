package binaryTreeSymmetric

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return recursion(root.Left, root.Right)
}

func recursion(left, right *TreeNode) bool {
	if left == right {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	// if !left.isSymmetric(right.Val) {
	if left.Val != right.Val {
		return false
	}

	if !recursion(left.Left, right.Right) {
		return false
	}

	return recursion(left.Right, right.Left)
}

func (t *TreeNode) isSymmetric(value int) bool {
	return t.Val == value
}
