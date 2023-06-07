package binarytreeinorder

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	value := []int{}
	value = append(value, inorderTraversal(root.Left)...)
	value = append(value, root.Val)
	value = append(value, inorderTraversal(root.Right)...)
	return value
}
